/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarcusXavierr/anki-helper/app/IO"
	"github.com/MarcusXavierr/anki-helper/app/check"
	"github.com/MarcusXavierr/anki-helper/app/sentenceCheck"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [sentence]",
	Short: "Adds a new word or sentence you learned",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 || len(args) > 1 {
			IO.PrintRed(os.Stdout, "this function only takes on argument")
			os.Exit(2)
		}

		wordsTrackerFile, trash := getFiles()
		sentence := args[0]

		saveSentence(sentence, wordsTrackerFile, trash)
	},
}

func saveSentence(sentence string, wordsTrackerFile, trash IO.File) {
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
