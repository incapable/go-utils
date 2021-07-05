// Package file Contains utility functions
package file

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const WorkspaceBase = "work"

var RemoveFiles bool

var log *zap.Logger

// On init, try to prepare the work directory
func init() {
	setupVars()
}

// Generates a unique identifier
func generateWorkspaceIdentifier() (string, error) {
	generatedUuid, err := uuid.New().MarshalText()
	if err != nil {
		return "", err
	}

	return string(generatedUuid), nil
}
