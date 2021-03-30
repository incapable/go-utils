package program

import (
	"errors"
	"os"
)

// Tries to get an argument value from the programs arguments
func ReadProgramArgument(key string) (string, error) {
	args := os.Args

	for index, arg := range args {
		if arg[:1] == "-" && arg[1:] == key {
			if len(args) < index+2 {
				return "", errors.New("value not found")
			} else {
				return args[index+1], nil
			}
		}
	}

	return "", errors.New("argument not found")
}

// Reads a flag from the programs arguments, false if not found
func ReadProgramFlag(key string) bool {
	for _, arg := range os.Args {
		if arg[:1] == "-" && arg[1:] == key {
			return true
		}
	}

	return false
}
