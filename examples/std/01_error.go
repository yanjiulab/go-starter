package main

import (
	"errors"
	"fmt"
)

var ErrNegative = errors.New("negative number")

// ValidationError is a custom error type carrying extra context.
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s invalid: %s", e.Field, e.Msg)
}

// classify returns different error kinds for demonstration.
func classify(n int) error {
	if n < 0 {
		return fmt.Errorf("classify: %w", ErrNegative)
	}
	if n == 0 {
		return &ValidationError{Field: "n", Msg: "cannot be zero"}
	}
	return nil
}

// safePanic demonstrates recover in deferred function.
func safePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	panic("demo panic")
}

func main() {
	// Typical pattern: check returned error and branch.
	for _, n := range []int{-1, 0, 3} {
		err := classify(n)
		if err == nil {
			fmt.Println("classify", n, "ok")
			continue
		}

		fmt.Println("classify", n, "err:", err)
		fmt.Println("  is ErrNegative?", errors.Is(err, ErrNegative))

		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("  as ValidationError:", ve.Field, ve.Msg)
		}
	}

	safePanic()
}
