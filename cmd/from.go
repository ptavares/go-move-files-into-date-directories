package cmd

import (
	"errors"
	"fmt"
	"move-files-into-date-directories/config"
	"move-files-into-date-directories/exception"
	"move-files-into-date-directories/file"
	"move-files-into-date-directories/helper"
	"move-files-into-date-directories/internal/name"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var absDirPath string

// fromCmd represents the config command
var fromCmd = &cobra.Command{
	Use:   "from [from_dir_path]",
	Short: fmt.Sprintf("Perform %s into a specified directory", name.ApplicationName),
	Long: fmt.Sprintf(`
=======================================================================
=                 %s from              =
=======================================================================

Command to move files from a specified directory into a new directory whose
name is based on the file's date.

If no destination directory is specified, will use current directory too.
You cans specify the scope at which directories should be created, accepted
values are %s.

Exemple : If you specify "day" 'default value), files will be moved
from current directory to 'destination\yyyyMMdd'

`, name.ApplicationName, config.DataScopes),
	PreRun: checkHereArgument,
	Args:   directoryArg(),
	Run:    executeFrom,
}

// Initialize subcommand
func init() {
	rootCmd.AddCommand(fromCmd)

	// -> Init subcommand flags
	// "Hour", "Day", "Month", or "Year". e.g. If you specify "Day" files will be moved from the `SourceDirectoryPath` to `TargetDirectoryPath\yyyy-MM-dd`.')]
	dateScope = config.Day
	fromCmd.PersistentFlags().VarP(&dateScope, "date-scope", "", fmt.Sprintf("the scope at which directories should be created. Accepted values %s", config.DataScopes))
	fromCmd.PersistentFlags().BoolVarP(&recursive, "recursive", "r", true, "will move all files and sub-directories files")
	fromCmd.PersistentFlags().StringVarP(&destinationDir, "destination", "", "", "destination directory, where files will be copied (if none, will use current directory)")
	fromCmd.PersistentFlags().StringVarP(&separator, "separator", "s", "", "separator to use when generating date file's date directory (default none)")
}

// directoryArg returns an error if arg is not an existing directory.
func directoryArg() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		if len(args) != 1 {
			return errors.New("requires one argument")
		}
		dirPath := args[0]
		absDirPath, err = filepath.Abs(dirPath)
		// if err, will use relative path
		if err != nil {
			absDirPath = dirPath
		}
		if _, err := os.Stat(absDirPath); os.IsNotExist(err) {
			helper.HandleErrorExit(exception.UnexistingDirError(absDirPath, err))
		}
		return nil
	}
}

func executeFrom(cmd *cobra.Command, args []string) {
	logger.Debug("start - executeFrom")
	// Init file client
	client, err := file.Init(absDirPath, destinationDir, separator, dateScope, recursive, dryRun)
	helper.HandleErrorExit(err)
	// Perform move files
	helper.HandleErrorExit(client.MoveFiles())
	logger.Debug("end   - executeFrom")

}
