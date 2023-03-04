package cmd

import (
	"fmt"
	"os"

	"github.com/MarcusXavierr/anki-helper/pkg/IO"
	ankiconnect "github.com/MarcusXavierr/anki-helper/pkg/anki-connect"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "anki-helper",
	Long: `anki helper is a CLI focused on helping you to save the new words you learn in a language.
And then you can save those words in your anki`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initializeConfig(cmd)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	write := IO.UserFilePath{WriteFile: "", TrashFile: ""}
	ankiConfig := ankiconnect.AnkiConfig{DeckName: "", ModelName: ""}
	rootCmd.AddCommand(dictionaryCmd, NewAddCmd(write), NewAnkiInsert(write, ankiConfig))
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
