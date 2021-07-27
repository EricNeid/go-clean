package goclean

import (
	"os"
	"testing"
	"time"

	"github.com/EricNeid/go-clean/internal/verify"
)

const TEST_DIR = "tmp"
const TEST_SUB_DIR = "tmp/subdir"

const TEST_FILE_1 = "tmp/file-1"
const TEST_FILE_2 = "tmp/file-2"
const TEST_FILE_3 = "tmp/subdir/file-3"

func TestGetExpiredFiles(t *testing.T) {
	// arrange
	// cleanup directory for testdata
	verify.Ok(t, os.RemoveAll(TEST_DIR))
	verify.Ok(t, os.MkdirAll(TEST_DIR, os.ModePerm))
	verify.Ok(t, os.MkdirAll(TEST_SUB_DIR, os.ModePerm))
	// create testfiles
	_, err := os.Create(TEST_FILE_1)
	verify.Ok(t, err)
	_, err = os.Create(TEST_FILE_2)
	verify.Ok(t, err)
	_, err = os.Create(TEST_FILE_3)
	verify.Ok(t, err)
	verify.Ok(t, os.Chtimes(TEST_FILE_1, time.Now(), time.Now()))
	verify.Ok(t, os.Chtimes(TEST_FILE_2, time.Now().Add(-1*24*time.Hour), time.Now().Add(-1*24*time.Hour)))
	verify.Ok(t, os.Chtimes(TEST_FILE_3, time.Now().Add(-2*24*time.Hour), time.Now().Add(-2*24*time.Hour)))

	// action
	// get old files that were not touched in the last 24h
	result := GetExpiredFiles(TEST_DIR, 1*23)

	// verify
	verify.Equals(t, 2, len(result))
}

func TestDeleteFiles(t *testing.T) {
	// arrange
	// cleanup directory for testdata
	verify.Ok(t, os.RemoveAll(TEST_DIR))
	verify.Ok(t, os.MkdirAll(TEST_DIR, os.ModePerm))
	verify.Ok(t, os.MkdirAll(TEST_SUB_DIR, os.ModePerm))
	// create testfiles
	f, _ := os.Create(TEST_FILE_1)
	f.Close()
	f, _ = os.Create(TEST_FILE_2)
	f.Close()
	f, _ = os.Create(TEST_FILE_3)
	f.Close()

	// action
	result := DeleteFiles([]string{TEST_FILE_2, TEST_FILE_3})

	// verify
	verify.Equals(t, 2, len(result))
}
