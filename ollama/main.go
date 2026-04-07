package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

var modelSizes = []string{"0.8b", "2b", "4b", "9b", "35b"}

func main() {
	var currentSize string

	tmpl, err := template.New("Modelfile.gotmpl").Funcs(template.FuncMap{
		"size": func() string {
			return currentSize
		},
	}).ParseFiles("Modelfile.gotmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing template: %v\n", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	for _, size := range modelSizes {
		currentSize = size
		buf.Reset()

		if err := tmpl.Execute(&buf, nil); err != nil {
			fmt.Fprintf(os.Stderr, "error executing template for size %s: %v\n", size, err)
			os.Exit(1)
		}

		if err := os.WriteFile("Modelfile-"+size, buf.Bytes(), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "error writing Modelfile-%s: %v\n", size, err)
			os.Exit(1)
		}
	}
}
