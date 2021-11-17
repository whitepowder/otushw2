package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	name  string
	param Param
}

type Param struct {
	fromPath string
	toPath   string
	offset   int64
	limit    int64
	error    error
}

func TestCopy(t *testing.T) {
	t.Run("Triggering Errors", func(t *testing.T) {
		defer os.Remove("out.txt")
		testsErr := []testCase{
			{name: "No source file", param: Param{"", "out.txt", 0, 0, Err404}},
			{name: "Unsupported file", param: Param{"testdata", "out.txt", 0, 0, ErrUnsupportedFile}},
			{name: "ExceedsFileSize", param: Param{"testdata/input.txt", "out.txt", 1000000, 0, ErrUnsupportedFile}},
			{name: "Can't create file", param: Param{"testdata/input.txt", "", 0, 0, ErrCantCreate}},
		}

		for _, tc := range testsErr {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				err := Copy(tc.param.fromPath, tc.param.toPath, tc.param.offset, tc.param.limit)
				require.Error(t, tc.param.error, err)
			})
		}
		t.Run("SucTest", func(t *testing.T) {
			suc := Copy("testdata/input.txt", "out.txt", 0, 10000000)
			require.Equal(t, nil, suc, "Should be equal")
		})
	})
}
