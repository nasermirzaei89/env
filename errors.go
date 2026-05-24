package env

import "fmt"

// NotSetError is returned by LookupXXX when the environment variable is not set.
type NotSetError struct {
	Key string
}

func (err NotSetError) Error() string {
	return fmt.Sprintf("environment variable %q not set", err.Key)
}

// InvalidValueError is returned by LookupXXX when the environment variable has an invalid value.
type InvalidValueError struct {
	Key   string
	Value string
	Err   error
}

func (err InvalidValueError) Error() string {
	return fmt.Sprintf("environment variable %q has an invalid value: %q", err.Key, err.Value)
}

func (err InvalidValueError) Unwrap() error {
	return err.Err
}
