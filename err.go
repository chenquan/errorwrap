package errorwrap

import (
	"errors"
	"strings"
)

// Errs an error that can hold multiple errors.
type Errs struct {
	errs []error
}

// Add adds errs to be, nil errors are ignored.
func (b *Errs) Add(err error) {
	if err == nil {
		return
	}

	i, ok := err.(interface {
		NotNil() bool
	})
	if ok {
		if b == err || !i.NotNil() {
			return
		}
	}

	b.errs = append(b.errs, err)
}

// Err returns an error that represents all errors.
func (b *Errs) Error() string {
	switch len(b.errs) {
	case 0:
		return ""
	case 1:
		return b.errs[0].Error()
	default:
		var builder strings.Builder
		builder.WriteString(b.errs[0].Error())
		for _, err := range b.errs[1:] {
			builder.WriteRune('\n')
			builder.WriteString(err.Error())
		}

		return builder.String()
	}
}

// Errors returns a copy of errs.
func (b *Errs) Errors() []error {
	if len(b.errs) == 0 {
		return nil
	}

	errs := make([]error, len(b.errs))
	copy(errs, b.errs)

	return errs
}

// NotNil checks if any error inside.
func (b *Errs) NotNil() bool {
	return len(b.errs) != 0
}

// Is reports whether any errs in err's chain matches target.
func (b *Errs) Is(target error) bool {
	for _, e := range b.errs {
		if errors.Is(e, target) {
			return true
		}
	}

	return false
}

// As finds the first errs in err's chain that matches target, and if one is found, sets
// target to that error value and returns true. Otherwise, it returns false.
func (b *Errs) As(target interface{}) bool {
	for _, err := range b.errs {
		if errors.As(err, target) {
			return true
		}
	}

	return false
}
