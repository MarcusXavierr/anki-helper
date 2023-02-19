package sentenceCheck

import (
	"fmt"
	"io"
	"strings"

	"github.com/MarcusXavierr/anki-helper/pkg/IO"
)

func CheckIfSentenceExists(out io.Writer, sentence string, files ...IO.IFile) bool {
	for _, file := range files {
		if hasSentence(sentence, file) {
			message := fmt.Sprintf("sentence %q already exists on file %s\n", sentence, file)
			IO.PrintRed(out, message)
			return true
		}
	}
	return false
}

func hasSentence(sentence string, file IO.IFile) bool {
	buffer := file.ReadFile()
	return verifyIfSentenceExists(sentence, buffer)
}

func verifyIfSentenceExists(sentence, buffer string) bool {
	var allLines []string = strings.Split(buffer, "\n")
	for line := range allLines {
		if compareStrings(sentence, allLines[line]) {
			return true
		}
	}
	return false
}

//Get one line from my file and compare to a sentence
func compareStrings(sentence, line string) bool {
	sentence = strings.ToLower(sentence)
	line = strings.ToLower(line)
	line = strings.Trim(line, " ")
	return strings.Compare(line, sentence) == 0
}
