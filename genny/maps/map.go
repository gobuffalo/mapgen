package maps

import "html/template"

// Map represents an implementation of the map
type Map struct {
	Name    template.HTML
	GoType  template.HTML
	Zero    template.HTML
	Package string
	A       template.HTML
	B       template.HTML
	BB      template.HTML
	C       template.HTML
}
