/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/MarcusXavierr/anki-helper/app/IO"
	"github.com/MarcusXavierr/anki-helper/app/dictionary"
	"github.com/MarcusXavierr/anki-helper/app/sentenceCheck"
	"github.com/MarcusXavierr/anki-helper/app/utils"
	"github.com/spf13/cobra"
)

const defaultConfigFilename = "anki-config"

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
		utils.PrettyPrintDefinition(response.Normalize())
	}
}

func saveSentence(sentence string) {
	wordsTrackerFile, trash := getFiles()

	if !sentenceCheck.CheckIfSentenceExists(os.Stdout, sentence, wordsTrackerFile, trash) {
		utils.WriteSentenceOnFile(sentence, wordsTrackerFile.FilePath)
	}
}

func getFiles() (IO.File, IO.File) {
	wordsTrackerFile := IO.File{FilePath: IO.GetHomeDir() + "/english_words/words.txt"}
	trash := IO.File{FilePath: IO.GetHomeDir() + "/english_words/trash.txt"}
	return wordsTrackerFile, trash
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.PersistentFlags().BoolP("definition", "d", false, "")
}

func getSentenceFromArgs(args []string) string {

	if len(args) < 1 || len(args) > 1 {
		utils.Usage()
	}
	sentence := args[0]
	return sentence
}
