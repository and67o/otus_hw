package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type returnCode int

type testCase struct {
	name       string
	command    string
	args       []string
	env        Environment
	returnCode returnCode
}

func TestRunCmd(t *testing.T) {
	for _, tst := range [...]testCase{
		{
			name:       "ok command",
			command:    "/bin/bash",
			args:       []string{"-c", "ls"},
			returnCode: 0,
		},
		{
			name:       "ok command(1)",
			command:    "pwd",
			returnCode: 0,
		},
		{
			name:       "fail unknown command",
			command:    "pwd1",
			returnCode: -1,
		},
		{
			name:       "fail permission denied",
			command:    "/bin/bash",
			args:       []string{"-c", "/dev/null"},
			returnCode: 126,
		},
		{
			name:       "fail exist file",
			command:    "/bin/bash",
			args:       []string{"-c", "touch", "executor.go"},
			returnCode: 1,
		},
	} {
		t.Run(tst.name, func(t *testing.T) {
			cmd := []string{tst.command}
			cmd = append(cmd, tst.args...)

			require.Equal(t, int(tst.returnCode), RunCmd(cmd, tst.env))
		})
	}

}
