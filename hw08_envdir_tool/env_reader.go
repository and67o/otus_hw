package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return env, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "=") {
			continue
		}
		filepath := fmt.Sprintf("%s/%s", dir, file.Name())
		envValue, err := getEnvValue(filepath)

		if err != nil {
			return env, nil
		}
		env[file.Name()] = envValue
	}

	return env, nil
}

func getEnvValue(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	line, _, err := reader.ReadLine()
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("failed to handle line: %w", err)
	}

	return handleLine(line), nil
}

func handleLine(line []byte) string {
	line = bytes.ReplaceAll(line, []byte{0x00}, []byte{'\n'})
	return strings.TrimRight(string(line), "\t")
}
