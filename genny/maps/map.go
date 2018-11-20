package maps

import (
	"errors"
	"html/template"
)

// Map represents an implementation of the map
type Map struct {
	Destination string
	Package     string
	Name        template.HTML
	GoType      template.HTML
	Zero        template.HTML
	A           template.HTML
	B           template.HTML
	BB          template.HTML
	C           template.HTML
}

func (m Map) Validate() error {
	if m.Destination == "" {
		return errors.New("you must set a destination package")
	}

	if m.GoType == "" {
		return errors.New("you must set a go type")
	}

	if m.Zero == "" {
		return errors.New("you must set a zero value")
	}

	if m.A == "" {
		return errors.New("you must set a A value")
	}

	if m.B == "" {
		return errors.New("you must set a B value")
	}

	if m.BB == "" {
		return errors.New("you must set a BB value")
	}

	if m.C == "" {
		return errors.New("you must set a C value")
	}
	return nil
}
