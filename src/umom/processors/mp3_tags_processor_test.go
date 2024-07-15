package processors_test

import (
	"log"
	"os"
	"testing"

	"github.com/Mayurifag/umom/processors"
	id3v2 "github.com/bogem/id3v2/v2"
)

func TestProcessMP3FileTags(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Close the file so it can be used by ProcessMP3FileTags
	tmpfile.Close()

	tag, err := id3v2.Open(tmpfile.Name(), id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}

	tag.SetArtist("Aphex  Twin<>:\"/\\|?*")
	tag.SetTitle(" Xtal ")

	comment := id3v2.CommentFrame{
		Encoding:    id3v2.EncodingUTF8,
		Language:    "eng",
		Description: "My opinion",
		Text:        "I like this song!",
	}
	tag.AddCommentFrame(comment)

	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}
	tag.Close()

	err = processors.ProcessMP3FileTags(tmpfile.Name())
	if err != nil {
		t.Errorf("ProcessMP3FileTags returned error: %v", err)
	}

	tag, err = id3v2.Open(tmpfile.Name(), id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	// Define expected frames
	expectedFrames := map[string]string{
		"TPE1": "Aphex Twin",
		"TIT2": "Xtal",
		"TCOP": "0.0.1test",
	}

	allFrames := tag.AllFrames()

	for frameID, frames := range allFrames {
		if expectedValue, allowed := expectedFrames[frameID]; allowed {
			// Frame ID is expected, check value and encoding
			if len(frames) > 0 {
				frameEntry := tag.GetLastFrame(frameID).(id3v2.TextFrame)
				frameText := frameEntry.Text
				frameEncoding := frameEntry.Encoding.Name

				if frameText != expectedValue {
					t.Errorf("Expected value: %s, got: %s", expectedValue, frameText)
				}

				if frameEncoding != id3v2.EncodingUTF16.Name {
					t.Errorf("Expected encoding to be UTF-16, got: %s", frameEncoding)
				}
			} else {
				t.Errorf("Frame %s not found", frameID)
			}
		} else {
			// Frame ID is unexpected, check its value (for logging purposes)
			for _, frame := range frames {
				t.Errorf("Unexpected frame %s", frame)
			}
		}
	}
}
