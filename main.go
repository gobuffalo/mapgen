package main

import (
	"context"
	"flag"
	"log"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/gogen"
	"github.com/gobuffalo/mapgen/genny/maps"
)

func main() {
	m := maps.Map{}

	flag.StringVar(&m.Name, "name", "", "set the name of the map")
	flag.StringVar(&m.Import, "import", "", "import a needed package")
	flag.StringVar(&m.Type, "type", "", "set the type of the map (string, int, y.Z, etc...)")
	flag.StringVar(&m.Zero, "zero", "", "sets the zero value of the type")

	flag.Parse()

	opts := &maps.Options{
		Maps: []maps.Map{m},
	}
	run := genny.WetRunner(context.Background())

	gg, err := maps.New(opts)
	if err != nil {
		log.Fatal(err)
	}
	run.WithGroup(gg)

	err = run.WithNew(gogen.Fmt(run.Root))
	if err != nil {
		log.Fatal(err)
	}

	err = run.Run()
	if err != nil {
		log.Fatal(err)
	}
}
