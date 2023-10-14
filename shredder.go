package shredder

import (
	"crypto/rand"
	"os"
)

// Config represents the configuration for file shredding.
type Config struct {
	Iterations int // Number of shredding iterations
	Remove     bool // Whether to remove the file after shredding
}

// File shreds a file based on the given configuration.
func (config Config) File(Path string) error {
	for i := 0; i < config.Iterations; i++ {
		if err := Shred(Path); err != nil {
			return err
		}
	}

	if config.Remove {
		if err := os.Remove(Path); err != nil {
			return err
		}
	}
	return nil
}

// Shred securely overwrites the contents of a file.
func Shred(Path string) error {
	// Open the file for writing
	file, err := os.OpenFile(Path, os.O_WRONLY, 0)
	if err != nil {
		return err
	}

	// Retrieve the file's size
	fileSize, err := file.Stat()
	if err != nil {
		return err
	}

	// Create a buffer with random data to overwrite the file
	buff := make([]byte, fileSize.Size())
	rand.Read(buff)

	// Write the random data to the file
	file.Write(buff)

	// Reset the file pointer to the beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	// Close the file
	file.Close()

	return nil
}
