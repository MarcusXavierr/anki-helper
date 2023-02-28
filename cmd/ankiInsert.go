package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MarcusXavierr/anki-helper/internal/utils"
	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	"github.com/spf13/cobra"
)

// ankiInsertCmd represents the ankiInsert command
func NewAnkiInsert(write utils.UserFilePath) *cobra.Command {
	ankiInsertCmd := &cobra.Command{
		Use:   "ankiInsert",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			dir, path := filepath.Split(write.WriteFile)
			x := os.DirFS(dir)
			result, err := IO.GetWords(x, path, 1)

			if err != nil {
				panic(err)
			}

			err = IO.MoveSentenceToTrash(write.TrashFile, write.WriteFile, result[0])

			if err != nil {
				panic(err)
			}

			fmt.Println(result)

		},
	}

	defaultNewWordsPath := fmt.Sprintf("%s/english_words/words.txt", IO.GetHomeDir())
	defaultTrashPath := fmt.Sprintf("%s/english_words/trash.txt", IO.GetHomeDir())

	ankiInsertCmd.Flags().StringVarP(&write.WriteFile, "new-words-file-path", "n", defaultNewWordsPath, "path to new words file")
	ankiInsertCmd.Flags().StringVarP(&write.TrashFile, "trash-file-path", "t", defaultTrashPath, "path to learned words file")

	ankiInsertCmd.Flags().StringVar(&config.DeckName, "deck-name", "", "Name of your anki deck")
	ankiInsertCmd.Flags().StringVar(&config.ModelName, "model-name", "", "your anki model (first you need to create one)")

	return ankiInsertCmd
}
