package utils

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	"github.com/MarcusXavierr/anki-helper/pkg/check"
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	"github.com/samber/lo"
)

type UserFilePath struct {
	WriteFile string
	TrashFile string
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
	IO.PrintGreen(os.Stdout, "Result for expression: ")
	printLnWithColor(sentence, IO.ColorGold)
	lo.ForEach(response.Usages, printResponseUsages)
}

func printResponseUsages(usage scraper.Usage, _ int) {
	printLnWithColor("\n--------------------------------------------------------", IO.ColorGreen)
	printLnWithColor("Part of speech: "+usage.PartOfSpeech, IO.ColorPink)
	lo.ForEach(usage.Definitions, printDefinition)
	fmt.Println("")

}

func printDefinition(definition scraper.Definition, _ int) {
	printLnWithColor("Definition: "+definition.WordDefinition, IO.ColorGreen)
	lo.ForEach(definition.Examples, func(example scraper.Example, _ int) {
		printLnWithColor("Example: "+string(example), IO.ColorCyan)
	})
	fmt.Println("")
}

func printLnWithColor(text, color string) {
	IO.PrintWithColor(os.Stdout, text+"\n", color)
}
