//go:build test

package file

import (
	"github.com/incapable/go-utils/logger"
	"os"
	"path/filepath"
)

// On init, try to prepare the work directory
func setupVars() {
	log = logger.GlobalLogger.Named("FileUtils")

	RemoveFiles = false
}

// PrepareWorkspace Creates a unique directory for the current job, and returns its name
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
