package parse

import (
	"fmt"
	"strings"
	"time"
)

// Function defines a generic function for parsing and assigning a value. The provided input, i,
// is the value supplied from the command line. The output value, o, should be assigned in this function.
type Function[T any] func(i string, o *T) error

func Bool(i string, o *bool) error {
	switch strings.TrimSpace(strings.ToLower(i)) {
	case "true", "1", "yes", "y":
		*o = true
	case "false", "0", "no", "n":
		*o = false
	default:
		return fmt.Errorf("unknown input: %s", i)
	}
	return nil
}

func Duration(i string, o *time.Duration) error {
	if i == "" {
		return nil
	}
	value, err := time.ParseDuration(i)
	if err != nil {
		return err
	}
	*o = value
	return nil
}

func From[T any](funcs ...Function[T]) Function[T] {
	return func(i string, o *T) error {
		for _, f := range funcs {
			if err := f(i, o); err != nil {
				return err
			}
		}
		return nil
	}
}
