package goclean

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func GetExpiredFiles(dir string, maxAgeH int) []string {
	latestModTime := time.Now().Add(-1 * time.Duration(maxAgeH) * time.Hour)
	var oldFiles []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			f, err := os.Stat(path)
			if err != nil {
				log.Printf("could not read %s\n", path)
			} else {
				if f.ModTime().Before(latestModTime) {
					oldFiles = append(oldFiles, path)
				}
			}
		}
		return nil
	})

	return oldFiles
}
