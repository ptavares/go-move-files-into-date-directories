package file

import (
	"fmt"
	"log"
	"move-files-into-date-directories/config"
	"move-files-into-date-directories/exception"
	"move-files-into-date-directories/helper"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
)

// Init : Init a file client
func Init(currentDir, destinationDir, separator string, dateScope config.DateScope, recursive, dryRun bool) (*Client, error) {
	logger := helper.GetLogger()
	logger.With(
		zap.String("currentDir", currentDir),
		zap.String("destinationDir", destinationDir),
		zap.String("separator", separator),
		zap.String("dateScope", dateScope.String()),
		zap.Bool("recursive", recursive),
		zap.Bool("dryRun", dryRun),
	).Debug("start - Init")

	// Test if currentDir exist
	if _, err := os.Stat(currentDir); os.IsNotExist(err) {
		logger.Errorf("%s doesn't exist", currentDir)
		return &Client{}, exception.UnexistingDirError(currentDir, err)
	}

	logger.Debug("end   - Init")
	return &Client{
		FromPath:    currentDir,
		Destination: destinationDir,
		Separator:   separator,
		DateScope:   dateScope,
		Recursive:   recursive,
		DryRun:      dryRun,
	}, nil
}

// getFormattedDate: format date to return the name for the directory
func (c *Client) getFormattedDate(date time.Time, dateScope config.DateScope, separator string) string {
	switch dateScope {
	case config.Hour:
		return fmt.Sprintf("%d%s%02d%s%02d%s%02d", date.Year(), separator, date.Month(), separator, date.Day(), separator, date.Hour())
	case config.Day:
		return fmt.Sprintf("%d%s%02d%s%02d", date.Year(), separator, date.Month(), separator, date.Day())
	case config.Month:
		return fmt.Sprintf("%d%s%02d", date.Year(), separator, date.Month())
	case config.Year:
		return fmt.Sprintf("%d", date.Year())
	default:
		return fmt.Sprintf("%d%s%02d%s%02d", date.Year(), separator, date.Month(), separator, date.Day())
	}
}

// visit: get all files for the given root directory
func (c *Client) visit(root string, files *[]FileInfo, recursive bool) filepath.WalkFunc {
	// Calculate original depth from the root file
	originalDepth := len(strings.Split(root, string(os.PathSeparator)))

	return func(path string, info os.FileInfo, err error) error {
		// if err, log and continue
		if err != nil {
			log.Fatal(err)
		}
		// if non recursive, ignore directories
		if info.IsDir() && !recursive {
			return nil
		}
		// check depth (ignore files in directories)
		if len(strings.Split(path, string(os.PathSeparator))) != originalDepth+1 {
			return nil
		}
		// add the file to the list of files
		*files = append(*files, FileInfo{
			Path:        path,
			Name:        info.Name(),
			IsDir:       info.IsDir(),
			DestDirName: c.getFormattedDate(info.ModTime(), c.DateScope, c.Separator),
		})
		return nil
	}
}

// MoveFiles: execute move files
func (c *Client) MoveFiles() error {
	logger := helper.GetLogger()
	logger.Debug("start - MoveFiles")

	var files []FileInfo
	var err error

	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgHiMagenta).SprintFunc()

	// Get all files to moves
	err = filepath.Walk(c.FromPath, c.visit(c.FromPath, &files, c.Recursive))
	if err != nil {
		panic(err)
	}

	// Sort files from path
	sort.Slice(files, func(i, j int) bool {
		return files[i].DestDirName < files[j].DestDirName
	})

	// Move files
	for _, file := range files {
		logger.Debugf("%+v\n", file)
		// Create destination dire before moving file
		dirDest := filepath.Join(c.Destination, file.DestDirName)
		absDirDest, err := filepath.Abs(dirDest)
		// if err, will use relative path
		if err != nil {
			absDirDest = dirDest
		}
		if _, err := os.Stat(absDirDest); os.IsNotExist(err) {
			logger.Infof("Creating destination directory %s", yellow(absDirDest))
			if !c.DryRun {
				err = os.MkdirAll(absDirDest, os.ModePerm)
				if err != nil {
					logger.Errorf("-> unable to create directory %s before moving file : %s", absDirDest, err)
					continue
				}
			}
		}
		// Move file
		dest := filepath.Join(c.Destination, file.DestDirName, file.Name)
		// Get absolute Path
		absDest, err := filepath.Abs(dest)
		// if err, will use relative path
		if err != nil {
			absDest = dest
		}
		logger.Infof("Moving %s to %s", magenta(file.Path), yellow(absDest))
		if !c.DryRun {
			err := os.Rename(file.Path, dest)
			if err != nil {
				logger.Errorf("-> unable to move %s to %s : %s", file.Path, dest, err)
			}
		}
	}

	logger.Debug("end   - MoveFiles")
	return nil
}
