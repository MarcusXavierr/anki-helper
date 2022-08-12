package sentenceCheck

import (
	"fmt"
	"io"
	"strings"

	"github.com/MarcusXavierr/anki-helper/app/IO"
)

func CheckIfSentenceExists(out io.Writer, sentence string, filePathList ...string) bool {
	for _, filePath := range filePathList {
		if hasSentence(sentence, filePath) {
			message := fmt.Sprintf("sentence %q already exists on file %s", sentence, filePath)
			IO.PrintRed(out, message)
			return true
		}
	}
	return false
}

func hasSentence(sentence, filePath string) bool {
	buffer := IO.ReadFile(filePath)
	return verifyIsSentenceExists(sentence, buffer)
}

func verifyIsSentenceExists(sentence, buffer string) bool {
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
