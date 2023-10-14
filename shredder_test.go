package shredder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testingDir = "./test_data"

func TestShredFile(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	// Shred the file
	err := Shred(filePath)
	assert.NoError(t, err, "Expected no error when shredding the file")
}

func TestFileWithConfig(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	config := Config{
		Iterations: 3,
		Remove:     false,
	}

	err := config.File(filePath)
	assert.NoError(t, err, "Expected no error when shredding with Config")

	// Check that the file still exists after shredding with Config
	_, err = os.Stat(filePath)
	assert.NoError(t, err, "Expected file to exist after shredding with Config")
}

func TestFileWithConfigAndRemoval(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	config := Config{
		Iterations: 3,
		Remove:     true,
	}

	err := config.File(filePath)
	assert.NoError(t, err, "Expected no error when shredding with Config and removal")

	// Check that the file no longer exists after shredding with Config and removal
	_, err = os.Stat(filePath)
	assert.Error(t, err, "Expected file to be removed after shredding with Config and removal")
}

func TestShredNonExistentFile(t *testing.T) {
	filePath := "non_existent_file.txt" // This file does not exist
	err := Shred(filePath)
	assert.Error(t, err, "Expected an error for shredding a non-existent file")
}

func TestShredFileWithInvalidPermissions(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	// Change the file permissions to read-only
	err := os.Chmod(filePath, 0400)
	if err != nil {
		t.Fatalf("Error changing file permissions: %v", err)
	}

	err = Shred(filePath)
	assert.Error(t, err, "Expected an error when shredding a file with invalid permissions")

	// Restore file permissions
	err = os.Chmod(filePath, 0644)
	if err != nil {
		t.Fatalf("Error restoring file permissions: %v", err)
	}
}

func createTestFile(data []byte, t *testing.T) string {
	// Create test directory if it doesn't exist.
	err := os.MkdirAll(testingDir, 0755)
	if err != nil {
		t.Fatalf("can't create test directory %s: %s", testingDir, err)
	}

	// Create test file
	filePath := testingDir + "/testfile"
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("can't create test file %s: %s", filePath, err)
	}

	// Write given data to test file
	n, err := f.Write(data)
	if err != nil {
		t.Fatalf("can't write to test file %s: %s", filePath, err)
	}
	assert.Equal(t, len(data), n, "Expected to write the correct number of bytes")

	err = f.Close()
	if err != nil {
		t.Fatalf("can't close file %s: %s", filePath, err)
	}

	return filePath
}
