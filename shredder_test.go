package shredder

import (
	"os"
	"testing"
)

const testingDir = "./test_data"

func TestShredFile(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	// Shred the file
	err := Shred(filePath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestFileWithConfig(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	config := Config{
		Iterations: 3,
		Remove:    false,
	}

	err := config.File(filePath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check that the file still exists after shredding with Config
	_, err = os.Stat(filePath)
	if err != nil {
		t.Errorf("Expected file to exist, but it did not.")
	}
}

func TestFileWithConfigAndRemoval(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	config := Config{
		Iterations: 3,
		Remove:    true,
	}

	err := config.File(filePath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check that the file no longer exists after shredding with Config and removal
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("Expected file to be removed, but it still exists.")
	}
}

func TestShredNonExistentFile(t *testing.T) {
	filePath := "non_existent_file.txt" // This file does not exist
	err := Shred(filePath)
	if err == nil {
		t.Error("Expected an error for a non-existent file, but got none.")
	}
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
	if err == nil {
		t.Error("Expected an error when shredding a file with invalid permissions, but got none.")
	}

	// Restore file permissions
	err = os.Chmod(filePath, 0644)
	if err != nil {
		t.Fatalf("Error restoring file permissions: %v", err)
	}
}

func TestFileWithConfig_Remove(t *testing.T) {
	// Create a test file with some data
	filePath := createTestFile([]byte("test data"), t)
	defer os.Remove(filePath)

	config := Config{
		Iterations: 1,
		Remove:    true,
	}

	err := config.File(filePath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check that the file no longer exists after shredding with Config and removal
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("Expected file to be removed, but it still exists.")
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
	if n != len(data) {
		t.Fatalf("wrote %d bytes only, %d bytes expected", n, len(data))
	}

	err = f.Close()
	if err != nil {
		t.Fatalf("can't close file %s: %s", filePath, err)
	}

	return filePath
}

