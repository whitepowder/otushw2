package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	// ErrLen is related to the value length error
	ErrLen = errors.New("lenth is not correct")
	// ErrRegExp is related to the symbol matching error
	ErrRegExp = errors.New("does not contain required symbols")
	// ErrMin is related to the min value error
	ErrMin = errors.New("less than min")
	// ErrMax is related to the max value error
	ErrMax = errors.New("larger than max")
	// ErrIn is related to the values comparasion error
	ErrIn = errors.New("should be specific")
)

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var sBuilder strings.Builder
	for i, err := range v {
		sBuilder.WriteString(err.Field)
		sBuilder.WriteString(" field")
		sBuilder.WriteString(" - ")
		sBuilder.WriteString(err.Err.Error())
		if i != len(v)-1 {
			sBuilder.WriteString("\n")
		}
	}
	return sBuilder.String()
}

func Validate(v interface{}) error {
	reflectvalue := reflect.ValueOf(v)

	if reflectvalue.Kind() != reflect.Struct {
		return fmt.Errorf("call police, we got %T instead of a pointer", v)
	}

	valueType := reflectvalue.Type()
	validationErr := new(ValidationErrors)

	var err error

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		reflectField := reflectvalue.Field(i)

		var tags []string
		tag := field.Tag.Get("validate")
		if strings.Contains(tag, "|") {
			tags = strings.Split(tag, "|")
		} else {
			tags = append(tags, tag)
		}
		if len(tag) != 0 {
			validationErr, err = kindValidation(field.Name, reflectField, tags, *validationErr)
			if err != nil {
				return err
			}
		}
	}
	return validationErr
}

func kindValidation(field string, value reflect.Value, tags []string, validationErr ValidationErrors) (*ValidationErrors, error) {
	var err error
	switch {
	case value.Kind() == reflect.String:
		v := value.Interface()
		switch i := v.(type) {
		case string:
			validationErr, err = stringValidation(field, i, tags, validationErr)
		default:
			validationErr, err = stringValidation(field, value.String(), tags, validationErr)
		}
	case value.Kind() == reflect.Int:
		v := value.Interface()
		switch i := v.(type) {
		case int:
			validationErr, err = intValidation(field, i, tags, validationErr)
		}
	case value.Kind() == reflect.Slice:
		v := value.Interface()
		switch i := v.(type) {
		case []string:
			for _, v := range i {
				validationErr, err = stringValidation(field, v, tags, validationErr)
			}
		case []int:
			for _, v := range i {
				validationErr, err = intValidation(field, v, tags, validationErr)
			}
		}
	}
	return &validationErr, err
}

func stringValidation(field string, value string, tags []string, validationErr ValidationErrors) (ValidationErrors, error) {
	var err error
	for _, tag := range tags {
		tagVal := strings.Split(tag, ":")[1]
		switch {
		case strings.HasPrefix(tag, "len:"):
			validationErr, err = lenValidation(field, value, tagVal, validationErr)
		case strings.HasPrefix(tag, "in:"):
			validationErr, err = inValidation(field, value, tagVal, validationErr)
		case strings.HasPrefix(tag, "regexp:"):
			validationErr, err = regexValidation(field, value, tagVal, validationErr)
		}
	}
	return validationErr, err
}

func intValidation(field string, value int, tags []string, validationErr ValidationErrors) (ValidationErrors, error) {
	var err error
	for _, tag := range tags {
		tagvalue := strings.Split(tag, ":")[1]
		switch {
		case strings.HasPrefix(tag, "in:"):
			i := strconv.Itoa(value)
			validationErr, err = inValidation(field, i, tagvalue, validationErr)
		case strings.HasPrefix(tag, "min:"):
			validationErr, err = minValidation(field, value, tagvalue, validationErr)
		case strings.HasPrefix(tag, "max:"):
			validationErr, err = maxValidation(field, value, tagvalue, validationErr)
		}
	}
	return validationErr, err
}

func lenValidation(field string, value string, tag string, validationErr ValidationErrors) (ValidationErrors, error) {
	var valErr ValidationError
	i, err := strconv.Atoi(tag)
	if err != nil {
		return validationErr, fmt.Errorf("can't decode: %w", err)
	}
	if len(value) != i {
		valErr.Field = field
		valErr.Err = ErrLen
		return append(validationErr, valErr), nil
	}
	return validationErr, nil
}

func inValidation(field string, value string, tag string, validationErr ValidationErrors) (ValidationErrors, error) {
	var valErr ValidationError
	var err error
	i := strings.Split(tag, ",")
	for _, v := range i {
		if err != nil {
			return validationErr, fmt.Errorf("can't range: %w", err) // voices in my head|linter asking me to do that
		}
		if v != value {
			valErr.Field = field
			valErr.Err = ErrIn
			return append(validationErr, valErr), nil
		}
	}
	return validationErr, nil
}

func regexValidation(field string, value string, tag string, validationErr ValidationErrors) (ValidationErrors, error) {
	var valErr ValidationError
	match, err := regexp.Match(tag, []byte(value))
	if err != nil {
		return validationErr, fmt.Errorf("can't match: %w", err)
	}
	if !match {
		valErr.Field = field
		valErr.Err = ErrRegExp
		return append(validationErr, valErr), nil
	}
	return validationErr, nil
}

func minValidation(field string, value int, tag string, validationErr ValidationErrors) (ValidationErrors, error) {
	var valErr ValidationError
	i, err := strconv.Atoi(tag)
	if err != nil {
		return validationErr, fmt.Errorf("can't decode: %w", err)
	}
	if i > value {
		valErr.Field = field
		valErr.Err = ErrMin
		return append(validationErr, valErr), nil
	}
	return validationErr, nil
}

func maxValidation(field string, value int, tag string, validationErr ValidationErrors) (ValidationErrors, error) {
	var valErr ValidationError
	i, err := strconv.Atoi(tag)
	if err != nil {
		return validationErr, fmt.Errorf("can't decode: %w", err)
	}
	if i < value {
		valErr.Field = field
		valErr.Err = ErrMax
		return append(validationErr, valErr), nil
	}
	return validationErr, nil
}
