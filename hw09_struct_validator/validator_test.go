package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr ValidationErrors
	}{
		{
			in: User{
				ID:    "33",
				Name:  "Pupa",
				Age:   10,
				Email: "dfsf",
				Role:  "guest",
				Phones: []string{
					"8800353535",
					"911",
				},
				meta: json.RawMessage{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   ErrLen,
				},
				ValidationError{
					Field: "Age",
					Err:   ErrMin,
				},
				ValidationError{
					Field: "Email",
					Err:   ErrRegExp,
				},
				ValidationError{
					Field: "Role",
					Err:   ErrIn,
				},
				ValidationError{
					Field: "Phones",
					Err:   ErrLen,
				},
				ValidationError{
					Field: "Phones",
					Err:   ErrLen,
				},
			},
		},

		{
			in: App{
				Version: "ver",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   ErrLen,
				},
			},
		},
		{
			in: Response{
				Code: 606,
				Body: "",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   ErrIn,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			err := Validate(tt.in)
			assert.Equal(t, &tt.expectedErr, err)

			_ = tt
		})
	}
}
