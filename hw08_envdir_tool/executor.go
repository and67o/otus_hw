package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := cmd[0]
	args := cmd[1:]
	c := exec.Command(command, args...)

	c.Env = getEnviron(env)

	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout

	err := c.Run()
	if err != nil {
		returnCode = c.ProcessState.ExitCode()
	}

	return
}

func getEnviron(env Environment) []string {
	for key, value := range env {
		if value == "" {
			err := os.Unsetenv(key)
			if err != nil {
				log.Fatal(err)
			}
			continue
		}

		err := os.Setenv(key, value)
		if err != nil {
			log.Fatal(err)
		}
	}

	return os.Environ()
}
