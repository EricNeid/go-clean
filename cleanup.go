package goclean

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func GetExpiredFiles(dir string, maxAgeH int) []string {
	log.Println("GetExpiredFiles")
	latestModTime := time.Now().Add(-1 * time.Duration(maxAgeH) * time.Hour)
	var expiredFiles []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			f, err := os.Stat(path)
			if err != nil {
				log.Printf("could not read %s, error %v\n", path, err)
			} else {
				if f.ModTime().Before(latestModTime) {
					expiredFiles = append(expiredFiles, path)
				}
			}
		}
		return nil
	})
	log.Printf("GetExpiredFiles detected %d files\n", len(expiredFiles))
	return expiredFiles
}

func DeleteFiles(files []string) []string {
	log.Println("DeleteFiles")
	var deletedFiles []string
	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			log.Printf("could not remove %s, error %v\n", f, err)
		} else {
			deletedFiles = append(deletedFiles, f)
		}
	}
	log.Printf("DeleteFiles deleted %d files\n", len(deletedFiles))
	return deletedFiles
}
