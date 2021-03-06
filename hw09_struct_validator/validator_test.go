package hw09_struct_validator //nolint:golint,stylecheck

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
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

	WrongUser struct {
		Age int `validate:"len:1:2:3"`
	}

	WrongApp struct {
		Version string `validate:"test:5"`
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

func TestValidateErrors(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          "test",
			expectedErr: ErrorNotStruct,
		},
		{
			in:          WrongUser{Age: 23},
			expectedErr: ErrorNotValidDefinition,
		},
		{
			in:          WrongApp{Version: "123"},
			expectedErr: ErrorUnknownRuleKey,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestValidatePass(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:User{
				ID:     "123456789123456789123456789123456789",
				Name:   "Oleg",
				Age:    24,
				Email:  "test@mail.ru",
				Role:   "admin",
				Phones: []string{"89204653384"},
				meta:   nil,
			},
		},
		{
			in: Response{Code: 200, Body: "test"},
		},
		{
			in: App{Version: "1.2.3"},
		},
		{
			in: Token{},
		},
		{
			in: Token{
				Header: []byte("any"),
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			require.Nil(t, err)
		})
	}
}

func TestValidateStructErrors(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: Response{Code: 201, Body: "test"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Code", Err: &InValidator{201}},
			},
		},
		{
			in: App{Version: "1234"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Version", Err: &LenValidator{5, "1234"}},
			},
		},
		{
			in: User{ID: "1"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: &LenValidator{36, "1"}},
				ValidationError{Field: "Age", Err: &MinValidator{18, 0}},
				ValidationError{Field: "Email", Err: &RegexValidator{""}},
				ValidationError{Field: "Role", Err: &InValidator{""}},
			},
		},
		{
			in: User{ID: "1", Name: "Oleg", Age: 9},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: &LenValidator{36, "1"}},
				ValidationError{Field: "Age", Err: &MinValidator{18, 9}},
				ValidationError{Field: "Email", Err: &RegexValidator{""}},
				ValidationError{Field: "Role", Err: &InValidator{""}},
			},
		},
		{
			in: User{ID: "123456789123456789123456789123456789", Name: "Oleg", Age: 56},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Age", Err: &MaxValidator{50, 56}},
				ValidationError{Field: "Email", Err: &RegexValidator{""}},
				ValidationError{Field: "Role", Err: &InValidator{""}},
			},
		},
		{
			in: User{ID: "123456789123456789123456789123456789", Name: "Oleg", Age: 45, Email: "oleil.ru"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Email", Err: &RegexValidator{"oleil.ru"}},
				ValidationError{Field: "Role", Err: &InValidator{""}},
			},
		},
		{
			in: User{ID: "123456789123456789123456789123456789", Name: "Oleg", Age: 45, Email: "oleg@mail.ru", Role: "oleg"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Role", Err: &InValidator{"oleg"}},
			},
		},
		{
			in: User{ID: "1234566789123456789", Name: "Oleg", Age: 85, Email: "oleil.ru", Role: "admin"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: &LenValidator{36, "1234566789123456789"}},
				ValidationError{Field: "Age", Err: &MaxValidator{50, 85}},
				ValidationError{Field: "Email", Err: &RegexValidator{"oleil.ru"}},
			},
		},
		{
			in: User{ID: "123456789123456789123456789123456789", Name: "Oleg", Age: 34, Email: "olei@ffl.ru", Role: "admin", Phones: []string{"123"}},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Phones", Err: &LenValidator{11, "123"}},
			},
		},
		{
			in: User{ID: "123456789123456789123456789123456789", Name: "Oleg", Age: 34, Email: "olei@ffl.ru", Role: "admin", Phones: []string{"12345678911", "23"}},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Phones", Err: &LenValidator{11, "23"}},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			err := Validate(tt.in)
			errors := err.(ValidationErrors)

			for index, e := range tt.expectedErr.(ValidationErrors) {
				validationError := errors[index]
				require.Equal(t, e.Field, validationError.Field)
				require.Equal(t, e.Err, validationError.Err)
			}
		})
	}
}
