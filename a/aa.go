package a

import (
	"log"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/pop/v5"
)

func init() {

	dropDatabaseYml()

	box := packr.New("buffalo:a:init", "./")
	if box.Has("inflections.json") {
		s, err := box.FindString("inflections.json")
		if err != nil {
			log.Fatal(err)
		}
		r := strings.NewReader(s)
		err = flect.LoadInflections(r)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func dropDatabaseYml() {
	box := packr.New("buffalo:a:init", "./")
	if box.Has("database.yml") {
		s, err := box.FindString("database.yml")
		if err != nil {
			log.Fatal(err)
		}
		r := strings.NewReader(s)
		err = pop.LoadFrom(r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
