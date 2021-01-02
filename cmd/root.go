package cmd

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFile string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "goremind",
	Short: "Modern implementation of remind",
	Run: func(cmd *cobra.Command, args []string) {

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

	zerolog.SetGlobalLevel(zerolog.FatalLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if configFile != "" {
		log.Debug().
			Str("configFile", configFile).
			Msg("Loading Configuration File")
		fmt.Println("reading config file", configFile)
	}

}
