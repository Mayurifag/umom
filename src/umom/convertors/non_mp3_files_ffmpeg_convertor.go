package convertors

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var knownMusicFormats = []string{".mp3", ".flac"}

func ffmpegExists() bool {
	_, err := exec.LookPath("ffmpeg")
	return err == nil
}

func ProcessNonMP3ViaFFMPEG(path string) (newFilePath string, isMusic bool, err error) {
	fileFormat := filepath.Ext(path)
	isMusic = false

	for _, format := range knownMusicFormats {
		if fileFormat == format {
			isMusic = true
			break
		}
	}

	if !isMusic {
		return path, false, nil
	}

	if fileFormat == ".mp3" {
		return path, true, nil
	}

	if !ffmpegExists() {
		return path, true, fmt.Errorf("ffmpeg is not installed on the system or not in PATH")
	}

	// Assuming we only want to convert .flac to .mp3
	if fileFormat == ".flac" {
		newFilePath = strings.TrimSuffix(path, fileFormat) + ".mp3"
		cmd := exec.Command("ffmpeg", "-i", path, "-ab", "320k", "-map_metadata", "0", "-id3v2_version", "3", newFilePath)
		err = cmd.Run()
		if err != nil {
			return path, true, fmt.Errorf("failed to convert %s to mp3: %v", path, err)
		}

		err = os.Remove(path)
		if err != nil {
			log.Printf("Failed to remove %s: %v", path, err)
		}

		return newFilePath, true, nil
	}

	return path, true, nil
}
