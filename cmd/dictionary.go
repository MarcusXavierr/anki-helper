/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	"github.com/MarcusXavierr/anki-helper/pkg/dictionary"
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/api"
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	dict "github.com/MarcusXavierr/anki-helper/internal/tui/commands/dictionary"
)

// dictionaryCmd represents the dictionary command
var dictionaryCmd = &cobra.Command{
	Use:   "dictionary",
	Short: "gets the definition of a word",
	Run: func(cmd *cobra.Command, args []string) {
		doStuff(args[0])
	},
}

func doStuff(sentence string) {
	res, err := api.GetDefinition("https://en.wiktionary.org/api/rest_v1/page/definition/", sentence)
	if err != nil {
		IO.PrintRed(os.Stdout, "was not possible to get the definition of this word\n")
		return
	}

	runTUI(res, sentence)
}

func runTUI(res scraper.Response, sentence string) {
	p := tea.NewProgram(
		dict.MakeViewPortModel(res, sentence),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program: ", err)
		os.Exit(1)
	}
}
func PrintResults(word string) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	result, err := dictionary.GetDefinition(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("results for word %q\n", result.Word)
	for _, meaning := range result.Meanings {
		fmt.Println(meaning.PartOfSpeech)
		fmt.Println("-------------------------------------------------")
		for _, definition := range meaning.Definitions {
			IO.PrintGreen(os.Stdout, fmt.Sprintf("Definition: %q\n", definition.Def))
			IO.PrintRed(os.Stdout, fmt.Sprintf("Example: %q\n\n", definition.Example))
		}
		fmt.Println("-------------------------------------------------")
	}
}
