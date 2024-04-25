package main

import (
	// "forensic-zip-tool/internal"
	"log"
	"fmt"
	"flag"
)

func handleError(err error, customMessage string) {
	if err != nil {
		log.Fatalf("Error %s: %s\n", customMessage, err)
	}
}

func main () {
	//parse commande line using flags. Flag -adp, -ha and -z stand respectively for artefact definition path, hash algorithm and zip 
	definitionPath := flag.String("adp", "../../configs/windowsArtifact.yaml", "[Optional] path to the artefact_definition.yaml")
	hashAlg := flag.String("ha", "sha256", "[Optional] Digest algorithm to use. default sha256")
    // archivePath := flag.String("z", "", "[Mandatory] Path to store the final ZIP archive")

	flag.Parse()

	// result, errParser := internal.ArtifactParser("/home/lchan/Documents/formations/projet_go/ForensicZipGo/test_directory/windows.yaml", "FILE")
	// handleError(errParser, "Parsing file\n")

	// fmt.Print(result)

	fmt.Printf("%v\n", *definitionPath)
	fmt.Printf("%v\n", *hashAlg)
}

