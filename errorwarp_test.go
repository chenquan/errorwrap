package errorwrap

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFn(t *testing.T) {
	t.Run("err", func(t *testing.T) {
		errFoo := errors.New("foo")
		errBar := errors.New("bar")
		err := Fn(func() error {
			return errFoo
		}).Fn(func() error {
			return errBar
		}).Finish()

		assert.ErrorIs(t, err, errFoo)

		err = Fn(func() error {
			return errFoo
		}).Fn(func() error {
			return nil
		}).Finish()

		assert.ErrorIs(t, err, errFoo)

		err = Fn(func() error {
			return nil
		}).Fn(func() error {
			return errFoo
		}).Fn(func() error {
			return nil
		}).Fns(func() error {
			return nil
		}, func() error {
			return nil
		}).Finish()

		assert.ErrorIs(t, err, errFoo)

		err = Fn(func() error {
			return nil
		}).Fn(func() error {
			return nil
		}).Fn(func() error {
			return nil
		}).Fns(func() error {
			return errBar
		}, func() error {
			return errFoo
		}).Finish()

		assert.ErrorIs(t, err, errFoo)
	})

	t.Run("no err", func(t *testing.T) {
		err := Fn(func() error {
			return nil
		}).Fn(func() error {
			return nil
		}).Finish()

		assert.NoError(t, err)
	})
}
