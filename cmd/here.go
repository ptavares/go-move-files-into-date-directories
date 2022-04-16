package cmd

import (
	"fmt"
	"move-files-into-date-directories/config"
	"move-files-into-date-directories/exception"
	"move-files-into-date-directories/file"
	"move-files-into-date-directories/helper"
	"move-files-into-date-directories/internal/name"
	"os"

	"github.com/spf13/cobra"
)

var (
	dateScope      config.DateScope
	recursive      bool
	destinationDir string
	separator      string
)

// hereCmd represents the config command
var hereCmd = &cobra.Command{
	Use:   "here",
	Short: fmt.Sprintf("Perform %s in current directory", name.ApplicationName),
	Long: fmt.Sprintf(`
=======================================================================
=                 %s here              =
=======================================================================

Command to move files from current directory into a new directory whose
name is based on the file's date.

If no destination directory is specified, will use current directory too.
You cans specify the scope at which directories should be created, accepted
values are %s.

Exemple : If you specify "day" 'default value), files will be moved
from current directory to 'destination\yyyyMMdd'

`, name.ApplicationName, config.DataScopes),
	PreRun: checkHereArgument,
	Run:    executeHere,
}

// Initialize subcommand
func init() {
	rootCmd.AddCommand(hereCmd)

	// -> Init subcommand flags
	// "Hour", "Day", "Month", or "Year". e.g. If you specify "Day" files will be moved from the `SourceDirectoryPath` to `TargetDirectoryPath\yyyy-MM-dd`.')]
	dateScope = config.Day
	hereCmd.PersistentFlags().VarP(&dateScope, "date-scope", "", fmt.Sprintf("the scope at which directories should be created. Accepted values %s", config.DataScopes))
	hereCmd.PersistentFlags().BoolVarP(&recursive, "recursive", "r", true, "will move all files and sub-directories files")
	hereCmd.PersistentFlags().StringVarP(&destinationDir, "destination", "", "", "destination directory, where files will be copied (if none, will use current directory)")
	hereCmd.PersistentFlags().StringVarP(&separator, "separator", "s", "", "separator to use when generating date file's date directory (default none)")
}

func checkHereArgument(cmd *cobra.Command, args []string) {
	logger.Debug("start - checkHereArgument")
	logger.Debug("end   - checkHereArgument")
}

func executeHere(cmd *cobra.Command, args []string) {
	logger.Debug("start - executeHere")

	logger.Debug("Get current directory")
	currentDir, err := os.Getwd()
	if err != nil {
		helper.HandleErrorExit(exception.GetPWDDirError(err))
	}
	// Init file client
	client, err := file.Init(currentDir, destinationDir, separator, dateScope, recursive, dryRun)
	helper.HandleErrorExit(err)
	// Perform move files
	helper.HandleErrorExit(client.MoveFiles())
	logger.Debug("end   - executeHere")
}
