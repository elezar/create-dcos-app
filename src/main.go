package main

import (
	"fmt"
	"log"

	"github.com/elezar/create-dcos-app/src/application"
)

func main() {
	inputFilename := "../examples/minimal/application.yml"
	outputPath := "../examples/minimal/output"

	application, err := application.LoadApplicationYaml(inputFilename)
	if err != nil {
		log.Fatal("Error loading application", err)
	}

	fmt.Println("Loaded application:")
	fmt.Print(application)

	err = application.Package.WriteToJsonFile(outputPath)
	if err != nil {
		log.Fatal(err)
	}
}
