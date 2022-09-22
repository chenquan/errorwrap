package errorwrap

type ErrorFn struct {
	fns []func() error
}

func Fn(fn func() error) *ErrorFn {
	return &ErrorFn{fns: []func() error{fn}}
}

func (e *ErrorFn) Fn(fn func() error) *ErrorFn {
	e.fns = append(e.fns, fn)
	return e
}

func (e *ErrorFn) Fns(fns ...func() error) *ErrorFn {
	e.fns = append(e.fns, func() error {
		var err Errs
		for _, fn := range fns {
			err.Add(fn())
		}
		return &err
	})

	return e
}

func (e *ErrorFn) Finish() error {
	for _, fn := range e.fns {
		err := fn()

		i, ok := err.(interface {
			NotNil() bool
		})
		if ok {
			if i.NotNil() {
				return err
			}
		} else {
			if err != nil {
				return err
			}
		}

	}

	return nil
}
