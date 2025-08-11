package validator

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9_'^&/+-]+(?:\\.[a-zA-Z0-9_'^&/+-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}$\n")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// PermittedValue returns true if a specific value is in a list
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all the values in an array are unique
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)
	for _, v := range values {
		uniqueValues[v] = true
	}

	return len(uniqueValues) == len(values)
}
