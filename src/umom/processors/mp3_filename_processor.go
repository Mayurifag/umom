package processors

import (
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"

	id3v2 "github.com/bogem/id3v2/v2"
)

func ProcessMP3FileName(path string) (newFilePath string, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", err
	}

	mp3File, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		return "", err
	}

	artist := mp3File.Artist()
	title := mp3File.Title()
	mp3File.Close()

	if artist == "" || title == "" {
		log.Printf("Skipping renaming since %s does not have artist or title tags.\n", path)
		return path, nil
	}

	newFileName := artist + " - " + title + ".mp3"
	dir := filepath.Dir(path)
	originalFileName := filepath.Base(path)
	newFilePath = filepath.Join(dir, newFileName)

	// Skip remaming if the new file name is the same as the existing one
	// In Go, string comparison is based on byte comparison by default, not on
	// Unicode code points. This can lead to unexpected behavior when dealing
	// with non-ASCII characters.

	// Convert strings to runes
	r1, _ := utf8.DecodeRuneInString(originalFileName)
	r2, _ := utf8.DecodeRuneInString(newFileName)

	if r1 == r2 {
		// log.Printf("Skipping renaming since %s already has the correct name.\n", path)
		return path, nil
	}

	log.Printf("Renaming %q to %q\n", originalFileName, newFileName)

	// Rename the file
	if err := os.Rename(path, newFilePath); err != nil {
		return "", err
	}

	return newFilePath, nil
}
