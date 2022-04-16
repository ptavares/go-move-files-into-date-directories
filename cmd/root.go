package cmd

import (
	"fmt"
	"move-files-into-date-directories/helper"
	"move-files-into-date-directories/internal/name"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Common variable
var (
	debug  bool
	dryRun bool
	logger *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "move-files-into-date-directories",
	Short: "Moves files from a directory into a new directory whose name is based on the file's date",
	Long: fmt.Sprintf(`
=======================================================================
=                   %s                  =
=======================================================================

Moves files from a directory into a new directory whose name is based on
the file's date.

A common use-case of this script is to move photos into date-named
directories based on when the photo was taken.

	`, name.ApplicationName),
	// PersistentPreRunE: runInit,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug message")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "", false, "run command in dry-run mode")
	// Run init
	cobra.OnInitialize(initRoot)
}

// initRoot : will load common configuration (logger)
func initRoot() {
	logger = helper.InitLogger(debug)
}
