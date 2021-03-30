// +build !test

package program

import (
	"bufio"
	"go.uber.org/zap"
	"os/exec"
	"strings"
	"time"
)

// Executes a command, logs the output to the logger, and waits for completion
func ExecuteAndWait(command string, logger zap.Logger, args []string) (err error, errorText string) {
	cmdStart := time.Now()

	logger.Debug("Executing command", zap.String("command", command), zap.String("args", strings.Join(args, " ")))

	// Prepare the command
	cmd := exec.Command(command, args...)

	// Add a reader to stderr
	cmdReader, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	errorText = ""

	// Read the error output
	for scanner.Scan() {
		line := scanner.Text()

		errorText += line + "/"
		logger.Debug(line)
	}

	// Wait for command completion
	err = cmd.Wait()
	if err != nil {
		return
	}

	cmdDuration := time.Now().Sub(cmdStart)
	logger.Debug("Completed command", zap.String("command", command), zap.Duration("duration", cmdDuration))

	errorText = ""

	return
}
