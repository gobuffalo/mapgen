package main

import (
	"context"
	"flag"
	"html/template"
	"log"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/gobuffalo/mapgen/genny/maps"
)

func main() {
	name := flag.String("name", "", "name of the map")
	gotype := flag.String("go-type", "", "name of the go type")
	pkg := flag.String("pkg", "", "package of the go type")
	zero := flag.String("zero", "", "zero value of go type")
	a := flag.String("a", "", "A value for testing")
	b := flag.String("b", "", "B value for testing")
	bb := flag.String("bb", "", "BB value for testing")
	c := flag.String("c", "", "C value for testing")
	flag.Parse()
	opts := &maps.Options{
		Maps: []maps.Map{
			{
				Name:    template.HTML(*name),
				GoType:  template.HTML(*gotype),
				Package: *pkg,
				Zero:    template.HTML(*zero),
				A:       template.HTML(*a),
				B:       template.HTML(*b),
				BB:      template.HTML(*bb),
				C:       template.HTML(*c),
			},
		},
	}
	// opts := &maps.Options{
	// 	Maps: []maps.Map{
	// 		{
	// 			Name:   "",
	// 			GoType: "interface{}",
	// 			Zero:   "nil",
	// 			A:      "0",
	// 			B:      "1",
	// 			BB:     "-1",
	// 			C:      "2",
	// 		},
	// 		{
	// 			Name:   "Int",
	// 			GoType: "int",
	// 			Zero:   "0",
	// 			A:      "0",
	// 			B:      "1",
	// 			BB:     "-1",
	// 			C:      "2",
	// 		},
	// 		{
	// 			Name:   "String",
	// 			GoType: "string",
	// 			Zero:   `""`,
	// 			A:      `"A"`,
	// 			B:      `"B"`,
	// 			BB:     `"BB"`,
	// 			C:      `"C"`,
	// 		},
	// 		{
	// 			Name:   "Byte",
	// 			GoType: "[]byte",
	// 			Zero:   `[]byte("")`,
	// 			A:      `[]byte("A")`,
	// 			B:      `[]byte("B")`,
	// 			BB:     `[]byte("BB")`,
	// 			C:      `[]byte("C")`,
	// 		},
	// 		{
	// 			Name:    "UUID",
	// 			GoType:  "uuid.UUID",
	// 			Package: "github.com/gofrs/uuid",
	// 			Zero:    `uuid.Nil`,
	// 			A:       `uuid.Must(uuid.FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c1"))`,
	// 			B:       `uuid.Must(uuid.FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c2"))`,
	// 			BB:      `uuid.Must(uuid.FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c3"))`,
	// 			C:       `uuid.Must(uuid.FromString("6ba7b814-9dad-11d1-80b4-00c04fd430c4"))`,
	// 		},
	// 	},
	// }

	// err := toml.NewEncoder(os.Stdout).Encode(struct {
	// 	Maps []maps.Map
	// }{
	// 	Maps: opts.Maps,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	run := genny.WetRunner(context.Background())

	gg, err := maps.New(opts)
	if err != nil {
		log.Fatal(err)
	}
	run.WithGroup(gg)

	err = run.WithNew(gotools.GoFmt(run.Root))
	if err != nil {
		log.Fatal(err)
	}

	err = run.Run()
	if err != nil {
		log.Fatal(err)
	}
}
