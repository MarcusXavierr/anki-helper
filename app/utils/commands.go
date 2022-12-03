package utils

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/MarcusXavierr/anki-helper/app/IO"
	"github.com/MarcusXavierr/anki-helper/app/check"
	"github.com/MarcusXavierr/anki-helper/app/dictionary"
)

type UserFilePath struct {
	WriteFile string
	TrashFile string
}

func PrettyPrintDefinition(response dictionary.DictionaryApiResponse) {
	IO.PrintGreen(os.Stdout, fmt.Sprintf("result for word %s\n\n", response.Word))
	for _, meaning := range response.Meanings {
		if len(meaning.Definitions) > 0 {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(meaning.Definitions))
			def := meaning.Definitions[randomIndex]
			IO.PrintGreen(
				os.Stdout,
				fmt.Sprintf("%s\nDefinition: %s\nExample: %s\n\n", meaning.PartOfSpeech, def.Def, def.Example),
			)
		}
	}

}

func Usage() {
	var message string = fmt.Sprintf("usage: %s \"sentence to add\"\n", os.Args[0])
	IO.PrintRed(os.Stdout, message)
	flag.PrintDefaults()
	os.Exit(2)
}

func WriteSentenceOnFile(sentence, filePath string) {
	err := IO.WriteFile(sentence, filePath)
	check.Check(err)
	message := fmt.Sprintf("Sentence %q added successfully\n", sentence)
	IO.PrintGreen(os.Stdout, message)
}
