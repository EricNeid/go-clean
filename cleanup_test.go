package goclean

import (
	"os"
	"testing"
	"time"

	"github.com/EricNeid/go-clean/internal/verify"
)

const TEST_DIR = "tmp"
const TEST_FILE_1 = "tmp/file-1"
const TEST_FILE_2 = "tmp/file-2"
const TEST_FILE_3 = "tmp/file-3"

func TestGetExpiredFiles(t *testing.T) {
	// arrange
	// cleanup directory for testdata
	err := os.RemoveAll(TEST_DIR)
	verify.Ok(t, err)
	err = os.MkdirAll(TEST_DIR, os.ModePerm)
	verify.Ok(t, err)
	// create testfiles
	os.Create(TEST_FILE_1)
	os.Create(TEST_FILE_2)
	os.Create(TEST_FILE_3)
	os.Chtimes(TEST_FILE_1, time.Now(), time.Now())
	os.Chtimes(TEST_FILE_2, time.Now().Add(-1*24*time.Hour), time.Now().Add(-1*24*time.Hour))
	os.Chtimes(TEST_FILE_3, time.Now().Add(-2*24*time.Hour), time.Now().Add(-2*24*time.Hour))

	// action
	// get old files that were not touched in the last 24h
	result := GetExpiredFiles(TEST_DIR, 1*23)

	// verify
	verify.Equals(t, 2, len(result))
}
