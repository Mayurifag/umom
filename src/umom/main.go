package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "github.com/Mayurifag/umom/converters"
	"github.com/Mayurifag/umom/processors"
)

func main() {
	rootDir := "path/to/your/music/files"

	// // List of supported formats
	// supportedFormats := []string{".flac"}

	// // Find and convert files
	// converters.ConvertFilesToMP3UsingFFMPEG(rootDir, supportedFormats)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if filepath.Ext(path) == ".mp3" {
				fmt.Printf("Converting %s\n", path)
				err := processors.ProcessMP3FileTags(path)
				if err != nil {
					log.Printf("Error converting %s: %v", path, err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", rootDir, err)
	}

}
