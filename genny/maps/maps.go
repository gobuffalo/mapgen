package maps

import (
	"fmt"

	"github.com/gobuffalo/flect/name"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/here"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
	"github.com/pkg/errors"
)

// New returns a generator capable of making sync.Map
// implementations
func New(opts *Options) (*genny.Group, error) {
	gg := &genny.Group{}
	box := packr.New("mapgen:genny:maps", "../maps/templates")

	if err := opts.Validate(); err != nil {
		return gg, errors.WithStack(err)
	}

	for _, m := range opts.Maps {
		g := genny.New()

		if err := (&m).Validate(); err != nil {
			return gg, err
		}

		nm := name.New(m.Name)

		s, err := box.FindString("map.go.plush")
		if err != nil {
			return gg, err
		}

		fn := fmt.Sprintf("%s.go.plush", nm.File())
		tmpl := genny.NewFileS(fn, s)

		g.RunFn(func(r *genny.Runner) error {
			hi, err := here.Dir(r.Root)
			if err != nil {
				return err
			}

			ctx := plush.NewContext()
			ctx.Set("opts", opts)
			ctx.Set("m", m)
			ctx.Set("here", hi)

			tf := plushgen.Transformer(ctx)
			f, err := tf.Transform(tmpl)
			if err != nil {
				return err
			}
			return r.File(f)
		})

		gg.Add(g)
	}
	return gg, nil
}
