package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var err error

	if len(cmd) < 1 {
		return 1
	}

	for key, value := range env {
		if value.NeedRemove {
			err = os.Unsetenv(key)
			if err != nil {
				log.Println(err)
				return 1
			}
		}
		err = os.Setenv(key, value.Value)
		if err != nil {
			log.Println(err)
			return 1
		}
	}

	command, args := cmd[0], cmd[1:]
	subc := exec.Command(command, args...)
	subc.Stdin = os.Stdin
	subc.Stdout = os.Stdout
	subc.Stderr = os.Stderr
	err = subc.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		return 1
	}
	return 0
}
