package processors_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Mayurifag/umom/processors"
	id3v2 "github.com/bogem/id3v2/v2"
)

func TestProcessMP3FileName(t *testing.T) {
	// Create a temporary dir for testing
	tempDir := t.TempDir()
	defer os.Remove(tempDir)

	t.Run("ExistingFile with correct tags", func(t *testing.T) {
		tempFile := filepath.Join(tempDir, "test.mp3")
		if _, err := os.Create(tempFile); err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tempFile)

		tag, err := id3v2.Open(tempFile, id3v2.Options{Parse: true})
		if err != nil {
			log.Fatal("Error while opening mp3 file: ", err)
		}

		tag.SetArtist("Aphex Twin")
		tag.SetTitle("Xtal")

		if err = tag.Save(); err != nil {
			log.Fatal("Error while saving a tag: ", err)
		}
		tag.Close()
		newFilePath, err := processors.ProcessMP3FileName(tempFile)

		if err != nil {
			t.Errorf("ProcessMP3FileName returned error: %v", err)
		}
		if newFilePath == "" {
			t.Error("Expected non-empty new file path, got empty")
		}
		// Verify if the file exists after renaming
		if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
			t.Errorf("Renamed file does not exist: %v", err)
		}
		os.Remove(newFilePath)
	})

	t.Run("NonExistingFile", func(t *testing.T) {
		_, err := processors.ProcessMP3FileName(filepath.Join(tempDir, "f.mp3"))
		if err == nil {
			t.Error("Expected error for non-existing file, got nil")
		}
	})
}
