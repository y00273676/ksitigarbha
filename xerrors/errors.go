package xerrors

import "strings"

// Errors allows us to combine multiple errors in to one.
type Errors struct {
	errs []error
}

// Errors returns the stored errors.
func (e *Errors) Errors() []error {
	return e.errs
}

func (e *Errors) Error() string {
	var sb strings.Builder
	sb.WriteString("errors: ")
	for i, err := range e.errs {
		sb.WriteString(err.Error())
		if i < e.Len()-1 {
			sb.WriteString("; ")
		}
	}

	return sb.String()
}

// Append another error into the errors list.
// Should notice that this is not thread safe.
func (e *Errors) Append(errs ...error) {
	e.errs = append(e.errs, errs...)
}

// Len shows how many errors appended.
func (e *Errors) Len() int {
	return len(e.errs)
}

// Valid shall be called since there might be no error
// inside Errors, in such case return this Errors struct as an error
// seems not appropriate. So preferably you should call Valid method
// before return Errors as a real error.
func (e *Errors) Valid() bool {
	return e.Len() > 0
}

// NewErrors ...
func NewErrors(errs ...error) *Errors {
	return &Errors{
		errs: errs,
	}
}
