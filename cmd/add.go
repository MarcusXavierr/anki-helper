/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/MarcusXavierr/anki-helper/internal/utils"
	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	"github.com/MarcusXavierr/anki-helper/pkg/sentenceCheck"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/api"
)

const defaultConfigFilename = ".anki-config"

func NewAddCmd() *cobra.Command {
	write := utils.UserFilePath{WriteFile: "", TrashFile: ""}

	addCmd := &cobra.Command{
		Use:   "add [sentence]",
		Short: "Adds a new word or sentence you learned",
		Run: func(cmd *cobra.Command, args []string) {
			saveSentenceAndPrintDefinition(cmd, args, write)
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}

	defaultNewWordsPath := fmt.Sprintf("%s/english_words/words.txt", IO.GetHomeDir())
	defaultTrashPath := fmt.Sprintf("%s/english_words/trash.txt", IO.GetHomeDir())

	addCmd.Flags().StringVarP(&write.WriteFile, "new-words-file-path", "n", defaultNewWordsPath, "só um teste")
	addCmd.Flags().StringVarP(&write.TrashFile, "trash-file-path", "t", defaultTrashPath, "só um teste trash")
	return addCmd
}

func saveSentenceAndPrintDefinition(cmd *cobra.Command, args []string, userFiles utils.UserFilePath) {
	definition, _ := cmd.Flags().GetBool("definition")
	sentence := getSentenceFromArgs(args)

	if definition {
		printWiktionaryDefinition(sentence)
	}

	saveSentence(sentence, userFiles)
}

func printWiktionaryDefinition(sentence string) {
	resp, err := api.GetDefinition("https://en.wiktionary.org/api/rest_v1/page/definition/", sentence)

	if err == api.NotFoundError {
		IO.PrintRed(os.Stdout, "word not found on dictionary api\n\n")
	}

	if err == nil {
		utils.PrintWiktionary(resp, sentence)
	}
}

func saveSentence(sentence string, userFiles utils.UserFilePath) {
	wordsTrackerFile, trash := getFiles(userFiles)

	if !sentenceCheck.CheckIfSentenceExists(os.Stdout, sentence, wordsTrackerFile, trash) {
		utils.WriteSentenceOnFile(sentence, wordsTrackerFile.FilePath)
	}
}

func getFiles(userFiles utils.UserFilePath) (IO.File, IO.File) {
	wordsTrackerFile := IO.File{FilePath: userFiles.WriteFile}
	trash := IO.File{FilePath: userFiles.TrashFile}
	return wordsTrackerFile, trash
}

func init() {
	rootCmd.AddCommand(NewAddCmd())
	rootCmd.PersistentFlags().BoolP("definition", "d", false, "")
}

func getSentenceFromArgs(args []string) string {

	if len(args) < 1 || len(args) > 1 {
		utils.Usage()
	}
	sentence := args[0]
	return sentence
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(defaultConfigFilename)

	v.AddConfigPath(IO.GetHomeDir())

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	bindFlags(cmd, v)
	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		configName := flag.Name

		if !flag.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
		}
	})

}
