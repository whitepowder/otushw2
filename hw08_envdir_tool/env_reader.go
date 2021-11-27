package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	// ErrDirPathEmty is an error
	ErrDirPathEmty = errors.New("dir path is empty")
	// ErrWrongPath is also an error
	ErrWrongPath = errors.New("wrong dir path")
	// ErrPath guess what? Error
	ErrPath = errors.New("bad file path")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if dir == "" {
		return nil, ErrDirPathEmty
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrWrongPath
	}

	env := make(Environment, len(files))
	for _, f := range files {
		if f.IsDir() || strings.Contains(f.Name(), "=") {
			continue
		}
		fPath := filepath.Join(dir, f.Name())
		info, err := os.Stat(fPath)
		if err != nil {
			return nil, ErrPath
		}

		fName := f.Name()

		if info.Size() == 0 {
			env[fName] = EnvValue{
				Value:      "",
				NeedRemove: true,
			}
			continue
		}

		f, err := os.Open(fPath)
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(f)
		if scanner.Scan() {
			v := scanner.Text()
			v = strings.TrimRight(v, " \t")
			v = string(bytes.ReplaceAll([]byte(v), []byte{0x00}, []byte("\n")))
			env[fName] = EnvValue{
				Value:      v,
				NeedRemove: false,
			}
		}
	}

	return env, nil
}
