package exception

import "fmt"

// GetHomeDirError : custom exception
func GetPWDDirError(err error) error {
	return fmt.Errorf("unable to get current directory : %w", err)
}

// UnexistingDirError : custom exception
func UnexistingDirError(dir string, err error) error {
	return fmt.Errorf("directory <%s> don't exist : %w", dir, err)
}
