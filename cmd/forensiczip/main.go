package main

import (
	"archive/zip"
	"forensic-zip-tool/internal/parser"
	"forensic-zip-tool/internal/zipHandler"

	"log"
	// "fmt"
	// "os"

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
	artefactMap, artifactParsingErr := parser.ParseArtifact(*parser.DefinitionPath, "FILE")
	handleError("parsing artefact definition", artifactParsingErr)

	hasherFunc, hashAlgErr := parser.ParseHashAlg(*parser.HashAlg)
	handleError("parsing hash algorithm", hashAlgErr)

	// fmt.Println(artefactMap)
	// fmt.Println(*parser.ArchivePath)

	// open file
	archive, achiveErr := zip.OpenReader(*parser.ArchivePath)
	handleError("failed to open zip file", achiveErr)
	defer archive.Close()

	//
	zipHandler.CopySpecificFiles(archive, artefactMap, hasherFunc)
}

