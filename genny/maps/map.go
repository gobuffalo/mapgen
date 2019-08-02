package maps

import (
	"strings"

	"github.com/gobuffalo/flect"
)

// Map represents an implementation of the map
type Map struct {
	Name   string // foo (becomes fooMap)
	Import string // github.com/x/y
	Type   string // string | int | y.Z | etc...
	Zero   string // "" | 0 | y.Z{} | etc...

}

func (m *Map) Validate() error {
	if len(m.Type) == 0 {
		m.Type = "string"
	}

	if len(m.Name) == 0 {
		m.Name = m.Type
	}
	n := flect.New(m.Name)
	m.Name = n.Camelize().String()

	if !strings.HasSuffix(m.Name, "Map") {
		m.Name = m.Name + "Map"
	}

	if len(m.Zero) == 0 {
		m.Zero = `""`
	}

	return nil
}
