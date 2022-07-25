package base

import (
	"fmt"
	"io/ioutil"

	"github.com/flosch/pongo2/v6"
)

func ReturnStringFromFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	return string(bytes)
}

func CompileGroovy(context pongo2.Context, template string) string {
	tpl, err := pongo2.FromString(template)
	if err != nil {
		panic(err)
	}

	out, err := tpl.Execute(context)
	if err != nil {
		panic(err)
	}

	return out
}
