package file

import "move-files-into-date-directories/config"

// Client : represent a client file
type Client struct {
	FromPath    string
	Destination string
	Separator   string
	DateScope   config.DateScope
	Recursive   bool
	DryRun      bool
}

// FileInfo: represents needed file information
type FileInfo struct {
	Path        string
	Name        string
	IsDir       bool
	DestDirName string
}
