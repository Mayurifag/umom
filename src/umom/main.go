package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Mayurifag/umom/processors"
)

func main() {
	if len(os.Args) > 1 {
		processCLIArgs()
	} else {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}
		processDirectory(currentDir)
	}
}

func processCLIArgs() {
	processingPath := os.Args[1]
	fileInfo, err := os.Stat(processingPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Error: Path %q does not exist", processingPath)
		} else {
			log.Fatalf("Error stating path %q: %v", processingPath, err)
		}
	}
	if fileInfo.Mode().IsRegular() {
		// It's a file
		fmt.Printf("Processing file: %s\n", processingPath)
		processFile(processingPath)
	} else if fileInfo.Mode().IsDir() {
		// It's a directory
		fmt.Printf("Processing directory: %s\n", processingPath)
		processDirectory(processingPath)
	} else {
		// Neither file nor directory (shouldn't happen in typical scenarios)
		log.Fatalf("Path %q is neither a file nor a directory", processingPath)
	}
}

func processDirectory(folderpath string) {
	// TODO: count files first, maybe show a progress bar or warn user about a large number of files
	err := filepath.Walk(folderpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			processFile(path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", folderpath, err)
	}
}

// TODO: process ffmpeg convertible files like flac, wav, etc.
// TODO: make yes/no prompt for each file - also possible edits, etc.
func processFile(filepath string) (newFilepath string) {
	if filepath[len(filepath)-4:] != ".mp3" {
		// log.Printf("Skipping %s: not an mp3 file\n", filepath)
		return
	}

	fmt.Printf("Processing %s\n", filepath)
	err := processors.ProcessMP3FileTags(filepath)
	if err != nil {
		log.Printf("Error converting %s: %v", filepath, err)
	}
	newFilepath, err = processors.ProcessMP3FileName(filepath)
	if err != nil {
		log.Printf("Error converting %s: %v", filepath, err)
	}

	return newFilepath
}
