package main

import (
	"flag"
	"io"
	"log"
	"os"

	cleanup "github.com/EricNeid/go-clean"
)

var (
	dir     string
	maxAgeD int
)

func init() {
	// TODO syslog

	log.SetFlags(0)
	log.SetPrefix("goclean: ")
	log.SetOutput(io.MultiWriter(os.Stdout))

	flag.StringVar(&dir, "d", dir, "set directory to cleanup")
	flag.IntVar(&maxAgeD, "t", maxAgeD, "clean all files that are older than 't' days")
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

	cleanup.GetExpiredFiles(dir, maxAgeD*24)
}
