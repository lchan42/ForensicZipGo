package main

import (
	"forensic-zip-tool/internal/parser"
	"log"
	"fmt"
	// "flag"
)

func handleError(customMessage string, err error ) {
	if err != nil {
		log.Fatalf("Error %s: %s\n", customMessage, err)
	}
}

//Flag -adp, -ha and -z stand respectively for artefact definition path, hash algorithm and zip
func main () {
	//parse flag from cmd line.
	if err := parser.ParseFlags(); err!=nil {
		handleError("parsing flag", err)
	}
	// parse artefact definition file to only select type = FILE
	artefactDefMap, artifactParsingErr := parser.ParseArtifact(*parser.DefinitionPath, "FILE")
	handleError("parsing artefact definition", artifactParsingErr)
	
	fmt.Println(artefactDefMap)
}

