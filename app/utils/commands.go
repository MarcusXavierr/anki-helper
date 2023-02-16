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
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	"github.com/samber/lo"
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

func printGreen(text string) {
	IO.PrintGreen(os.Stdout, text)
}

func printCyan(text string) {
	IO.PrintCyan(os.Stdout, text)
}

func printPink(text string) {
	IO.PrintPink(os.Stdout, text)
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

func PrintWiktionary(response scraper.Response, sentence string) {
	printGreen("Result for expression: ")
	IO.PrintWithColor(os.Stdout, fmt.Sprintf("%s\n", sentence), string("\033[38;2;243;134;48m"))
	lo.ForEach(response.Usages, printUsage)
}

func printUsage(usage scraper.Usage, _ int) {
	printGreen("\n--------------------------------------------------------\n")
	printPink("Part of speech: " + usage.PartOfSpeech + "\n")
	lo.ForEach(usage.Definitions, printDefinition)
	fmt.Println("")

}

func printDefinition(definition scraper.Definition, _ int) {
	printGreen("Definition: " + definition.WordDefinition + "\n")
	lo.ForEach(definition.Examples, func(example scraper.Example, _ int) {
		printCyan("Example: " + string(example) + "\n")
	})
	printGreen("\n")
}
