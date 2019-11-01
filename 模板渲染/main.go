package main

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
)

func main() {
	tpl := `Hello {{cat "2321" "sdfsd"}}`

	// Get the Sprig function map.
	fmap := sprig.FuncMap()
	tt := template.Must(template.New("test").Funcs(fmap).Parse(tpl))

	err := tt.Execute(os.Stdout, "")
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}

	//tpl := template.Must(
	//	template.New("base").Funcs(sprig.FuncMap()).ParseGlob("*.html"),
	//)
	//fmt.Println(tpl.Name())
	//tpl.ex
}
