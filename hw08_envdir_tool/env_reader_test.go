package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	env := make(Environment)
	env["BAR"] = EnvValue{"bar", false}
	env["EMPTY"] = EnvValue{"", false}
	env["FOO"] = EnvValue{"   foo\nwith new line", false}
	env["HELLO"] = EnvValue{`"hello"`, false}
	env["UNSET"] = EnvValue{"", true}
	envs, err := ReadDir("testdata/env")
	require.NoError(t, err)
	require.Equal(t, envs, env)
}

func TestErrors(t *testing.T) {
	t.Run("EmptyDirPath", func(t *testing.T) {
		_, err := ReadDir("")
		require.ErrorIs(t, err, ErrDirPathEmty)
	})

	t.Run("WrongDirPath", func(t *testing.T) {
		_, err := ReadDir("wrongpath")
		require.ErrorIs(t, err, ErrWrongPath)
	})
}
