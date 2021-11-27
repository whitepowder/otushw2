package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Expected 0", func(t *testing.T) {
		env, _ := ReadDir("testdata/env")
		cmd := []string{"./go-envdir", "/home/wp/otushw2/hw08_envdir_tool/testdata/env", "/bin/bash", "/home/wp/otushw2/hw08_envdir_tool/testdata/echo.sh", "arg1=1", "arg2=2"}
		exitCode := RunCmd(cmd, env)
		require.Equal(t, 0, exitCode)
	})

	t.Run("Expected 1", func(t *testing.T) {
		env, _ := ReadDir("testdata/env")
		cmd := []string{"", "", "", ""}
		exitCode := RunCmd(cmd, env)
		require.Equal(t, 1, exitCode)
	})
}
