package main

import (
	"log"

	"github.com/elezar/create-dcos-app/src/application"
)

func generateConfigJson(a application.Application, outputPath string) {
	log.Println("Generating config.json")
	err := a.Config.WriteToJsonFile(outputPath)
	if err != nil {
		log.Fatal(err)
	}
}

func generateMarathonJsonMustache(a application.Application, outputPath string) {
	log.Println("TODO: Generate Marathon.json.mustache")
}

func generatePackageJson(a application.Application, outputPath string) {
	log.Println("Generating package.json")
	err := a.Package.WriteToJsonFile(outputPath)
	if err != nil {
		log.Fatal(err)
	}
}

func generateResourceJson(a application.Application, outputPath string) {
	log.Println("TODO: Generate resource.json")
}

func generateSvcYaml(a application.Application, outputPath string) {
	log.Println("TODO: Generate resource.json")
}

func main() {
	inputFilename := "../examples/minimal/application.yml"
	outputPath := "../examples/minimal/output"

	application, err := application.LoadApplicationYaml(inputFilename)
	if err != nil {
		log.Fatal("Error loading application", err)
	}

	generateConfigJson(application, outputPath)
	generateMarathonJsonMustache(application, outputPath)
	generatePackageJson(application, outputPath)
	generateResourceJson(application, outputPath)

	generateSvcYaml(application, outputPath)

}
