// +build test

package file

import (
	logger2 "github.com/incapable/videosplicer/utils/logger"
	"os"
	"path/filepath"
)

// On init, try to prepare the work directory
func setupVars() {
	logger = logger2.GlobalLogger.Named("FileUtils")

	RemoveFiles = false
}

// Creates a unique directory for the current job, and returns its name
func PrepareWorkspace() (string, error) {
	identifier, err := generateWorkspaceIdentifier()

	if err != nil {
		return "", err
	}

	workspacePath, err := filepath.Abs(WorkspaceBase + string(os.PathSeparator) + identifier)
	if err != nil {
		return "", err
	}

	return workspacePath, nil
}
