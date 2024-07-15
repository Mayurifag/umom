package processors

import (
	"log"
	"os"
	"path/filepath"

	id3v2 "github.com/bogem/id3v2/v2"
)

func ProcessMP3FileName(path string) (newFilePath string, err error) {
	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}

	// Open the MP3 file
	mp3File, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		return "", err
	}

	// Extract artist and title tags
	artist := mp3File.Artist()
	title := mp3File.Title()
	mp3File.Close()

	// Construct the new file name
	newFileName := artist + " - " + title + ".mp3"
	dir := filepath.Dir(path)
	newFilePath = filepath.Join(dir, newFileName)
	log.Printf("Renaming %s to %s\n", path, newFileName)

	// Rename the file
	if err := os.Rename(path, newFilePath); err != nil {
		return "", err
	}

	return newFilePath, nil
}
