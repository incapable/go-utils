// +build test

package program

import (
	"go.uber.org/zap"
)

// Executes a command, logs the output to the logger, and waits for completion
func ExecuteAndWait(command string, logger zap.Logger, args []string) (err error, errorText string) {
	logger.Info("ran cmd",
		zap.String("command", command),
		zap.Any("args", args))
	return nil, ""
}
