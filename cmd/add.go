/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/MarcusXavierr/anki-helper/app/IO"
	"github.com/MarcusXavierr/anki-helper/app/check"
	"github.com/MarcusXavierr/anki-helper/app/dictionary"
	"github.com/MarcusXavierr/anki-helper/app/sentenceCheck"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [sentence]",
	Short: "Adds a new word or sentence you learned",
	Run: func(cmd *cobra.Command, args []string) {
		saveSentenceAndPrintDefinition(cmd, args)
	},
}

func saveSentenceAndPrintDefinition(cmd *cobra.Command, args []string) {
	definition, _ := cmd.Flags().GetBool("definition")
	sentence := getSentenceFromArgs(args)

	if definition {
		printDefinition(sentence)
	}

	saveSentence(sentence)
}

func printDefinition(sentence string) {
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + sentence
	response, err := dictionary.GetDefinition(url)
	if err == IO.NotFoundError {
		IO.PrintRed(os.Stdout, "word not found on dictionary api\n\n")
	}
	if err == nil {
		PrettyPrintDefinition(response.Normalize())
	}
}

func saveSentence(sentence string) {
	wordsTrackerFile, trash := getFiles()

	if !sentenceCheck.CheckIfSentenceExists(os.Stdout, sentence, wordsTrackerFile, trash) {
		writeSentenceOnFile(sentence, wordsTrackerFile.FilePath)
	}
}

func getFiles() (IO.File, IO.File) {
	wordsTrackerFile := IO.File{FilePath: IO.GetHomeDir() + "/english_words/words.txt"}
	trash := IO.File{FilePath: IO.GetHomeDir() + "/english_words/trash.txt"}
	return wordsTrackerFile, trash
}

func writeSentenceOnFile(sentence, filePath string) {
	err := IO.WriteFile(sentence, filePath)
	check.Check(err)
	message := fmt.Sprintf("Sentence %q added successfully\n", sentence)
	IO.PrintGreen(os.Stdout, message)
}

func usage() {
	var message string = fmt.Sprintf("usage: %s \"sentence to add\"\n", os.Args[0])
	IO.PrintRed(os.Stdout, message)
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.PersistentFlags().BoolP("definition", "d", false, "")
}

func getSentenceFromArgs(args []string) string {

	if len(args) < 1 || len(args) > 1 {
		IO.PrintRed(os.Stdout, "this function only takes on argument")
		os.Exit(2)
	}
	sentence := args[0]
	return sentence
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
