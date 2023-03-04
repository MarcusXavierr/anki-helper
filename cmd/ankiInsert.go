package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	ankiinsert "github.com/MarcusXavierr/anki-helper/internal/tui/commands/anki_insert"
	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// ankiInsertCmd represents the ankiInsert command
func NewAnkiInsert(write IO.UserFilePath, config ankiconnect.AnkiConfig) *cobra.Command {
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
			// result, err := IO.GetWords(x, path, 1)
			//
			// if err != nil {
			// 	panic(err)
			// }
			//
			// err = IO.MoveSentenceToTrash(write.TrashFile, write.WriteFile, result[0])
			//
			// if err != nil {
			// 	panic(err)
			// }
			//
			// fmt.Println(result)
			//
			// fmt.Println(config.DeckName)
			// fmt.Println(config.ModelName)
			sentences, err := IO.GetWords(x, path, config.MaxItems)
			if err != nil {
				fmt.Printf("Error reading file: %s\n", path)
				os.Exit(1)
			}

			setLogFile()

			cards := ankiconnect.CreateCardDataMatrix(sentences, config)
			m := ankiinsert.MakeModel(config, cards, write)

			p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

			if _, err := p.Run(); err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		},
	}

	defaultNewWordsPath := fmt.Sprintf("%s/english_words/words.txt", IO.GetHomeDir())
	defaultTrashPath := fmt.Sprintf("%s/english_words/trash.txt", IO.GetHomeDir())
	defaultManualPath := filepath.Join(IO.GetHomeDir(), "english_words", "manual.txt")

	ankiInsertCmd.Flags().StringVarP(&write.WriteFile, "new-words-file-path", "n", defaultNewWordsPath, "path to new words file")
	ankiInsertCmd.Flags().StringVarP(&write.TrashFile, "trash-file-path", "t", defaultTrashPath, "path to learned words file")
	ankiInsertCmd.Flags().StringVarP(&write.ManualInsertFile, "manual-file-path", "m", defaultManualPath, "path to the file that store sentences that could not be inserted automatically")

	ankiInsertCmd.Flags().IntVarP(&config.MaxItems, "number-of-items", "i", 0, "The number of cards to be created")
	ankiInsertCmd.Flags().StringVar(&config.DeckName, "deck-name", "", "Name of your anki deck")
	ankiInsertCmd.Flags().StringVar(&config.ModelName, "model-name", "", "your anki model (first you need to create one)")

	return ankiInsertCmd
}

func setLogFile() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
}
