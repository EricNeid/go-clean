package main

import (
	"flag"
	"io"
	"log"
	"os"

	cleanup "github.com/EricNeid/go-clean"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	dir     string
	maxAgeD int
	logFile string
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("goclean: ")
	log.SetOutput(os.Stdout)

	flag.StringVar(&dir, "d", dir, "set directory to cleanup")
	flag.IntVar(&maxAgeD, "t", maxAgeD, "clean all files that are older than 't' days")
	flag.StringVar(&logFile, "l", logFile, "Optional: write log to given file")
	flag.Parse()
}

func main() {
	if len(dir) == 0 {
		log.Println("no directory to clean specified")
		flag.Usage()
		os.Exit(1)
	}

	if maxAgeD <= 0 {
		log.Println("maximum age for cleanup must be larger than 0")
		flag.Usage()
		os.Exit(1)
	}

	if len(logFile) >= 0 {
		log.SetOutput(
			io.MultiWriter(
				os.Stdout,
				&lumberjack.Logger{
					Filename:   logFile,
					MaxSize:    500, // megabytes
					MaxBackups: 3,
					MaxAge:     28, //days
				},
			),
		)
	}

	expiredFiles := cleanup.GetExpiredFiles(dir, maxAgeD*24)
	cleanup.DeleteFiles(expiredFiles)
}
