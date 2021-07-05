//go:build !test

package file

import (
	"github.com/incapable/go-utils/logger"
	"github.com/incapable/go-utils/program"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// On init, try to prepare the work directory
func setupVars() {
	log = logger.GlobalLogger.Named("FileUtils")

	RemoveFiles = !program.ReadProgramFlag("debug") // If debug is enabled, we'll keep the files

	_, err := os.Stat(WorkspaceBase)
	if os.IsNotExist(err) {
		err := os.Mkdir(WorkspaceBase, 0755)
		if err != nil {
			log.Error("ErrorResponse creating directory", zap.Error(err))
			os.Exit(ErrWorkdirCreate)
		}
	} else if err != nil {
		log.Error("Could not prepare the workdir", zap.Error(err))
		os.Exit(ErrWorkdirUnknown)
	}
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

	err = os.Mkdir(workspacePath, 0755)
	if err != nil {
		return "", err
	}

	return workspacePath, nil
}
