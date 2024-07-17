package processors

import (
	"log"
	"regexp"
	"strings"

	id3v2 "github.com/bogem/id3v2/v2"
)

func ProcessMP3FileTags(path string) error {
	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	if isAlreadyProcessedMP3File(tag) {
		return nil
	}

	tag.SetVersion(3) // ID3v2.3 - so tags will be readen everywhere fine

	if err := normalizeTags(tag); err != nil {
		return err
	}
	if err := setProcessedFlagTag(tag); err != nil {
		return err
	}

	if err := tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}

	return nil
}

const checkTag = "Copyright message"
const checkVersion = "0.0.2test"

// Checks if file has tag "Copyright message" with value "0.0.2test"
// Didnt sure what to use, https://github.com/n10v/id3v2/blob/main/common_ids.go
// There is also UserDefinedTextFrame. TODO: research
func isAlreadyProcessedMP3File(tag *id3v2.Tag) bool {
	tf := tag.GetTextFrame(tag.CommonID(checkTag))
	return tf.Text == checkVersion
}

var allowedFrames = map[string]struct{}{
	"TPE1": {}, // "Lead performer(s)/Soloist(s)"
	"TALB": {}, // "Album/Movie/Show title"
	"TIT2": {}, // "Title/song name/content description"
	"APIC": {}, // "Attached picture"
}

// Removes all tags except for artist, album and song name
func normalizeTags(tag *id3v2.Tag) error {
	allFrames := tag.AllFrames()
	for frameID, frames := range allFrames {
		if _, allowed := allowedFrames[frameID]; !allowed {
			tag.DeleteFrames(frameID)
		} else {
			for _, frame := range frames {
				if textFrame, ok := frame.(id3v2.TextFrame); ok {
					normalizedFrame := normalizeTextFrame(frameID, textFrame)
					tag.AddFrame(frameID, normalizedFrame)
				}
			}
		}
	}
	return nil
}

// SUGGEST: should I add here \t\n\r\f\v characters?
var forbiddenChars = regexp.MustCompile(`[<>:"/\\|?*]`) // Windows forbidden characters

// Converts tags to UTF-16 encoding
// Removes spaces from the beginning and end of the text
// Removes extra spaces between words
// Removes forbidden characters from tags to be compatible with Windows 10+ file system
// Normalizes a text frame by performing the specified operations.
func normalizeTextFrame(frameID string, textFrame id3v2.TextFrame) id3v2.TextFrame {
	// Normalize the text
	text := textFrame.Text
	text = strings.TrimSpace(text)
	if frameID == "TPE1" || frameID == "TIT2" {
		text = forbiddenChars.ReplaceAllString(text, "")
	}
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return id3v2.TextFrame{
		Encoding: id3v2.EncodingUTF16,
		Text:     text,
	}
}

// Sets custom tag to indicate that the file has been processed (Copyright message: 0.0.2test)
func setProcessedFlagTag(tag *id3v2.Tag) error {
	textFrame := id3v2.TextFrame{
		Encoding: id3v2.EncodingUTF16,
		Text:     checkVersion,
	}

	tag.AddFrame(tag.CommonID(checkTag), textFrame)

	return nil
}
