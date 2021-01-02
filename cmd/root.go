package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cgardner/goremind/reminder"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	configFile string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "goremind",
	Short: "Modern implementation of remind",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Str("configFile", configFile).Msg("Starting execution")
		file, err := os.Open(configFile)
		if err != nil {
			log.Fatal().Err(err).Str("configFile", configFile).Msg("Error Reading Config File")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		// lines := []string{}
		for scanner.Scan() {
			line := scanner.Text()
			if reminder.IsComment(line) {
				continue
			}
			if reminder.IsReminder(line) {
				reminder.NewReminder(line)
			}
		}
		err = scanner.Err()
		if err != nil {
			log.Fatal().Msg("Failed to read the config file")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "(required) config file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "use verbose logging")

	rootCmd.MarkPersistentFlagRequired("config")
}

func initConfig() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
