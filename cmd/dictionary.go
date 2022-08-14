/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcusXavierr/anki-helper/app/IO"
	"github.com/MarcusXavierr/anki-helper/app/dictionary"
	"github.com/spf13/cobra"
)

// dictionaryCmd represents the dictionary command
var dictionaryCmd = &cobra.Command{
	Use:   "dictionary",
	Short: "A quick test",
	Run: func(cmd *cobra.Command, args []string) {
		PrintResults(args[0])
	},
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
			if definition.Example != "" {
				IO.PrintGreen(os.Stdout, fmt.Sprintf("Definition: %q\n", definition.Def))
				IO.PrintRed(os.Stdout, fmt.Sprintf("Example: %q\n\n", definition.Example))
			}
		}
		fmt.Println("-------------------------------------------------")
	}
}
