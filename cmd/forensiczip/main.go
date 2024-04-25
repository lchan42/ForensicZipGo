package main

import (
	"forensic-zip-tool/internal"
	"log"
	"fmt"
)

func handleError(err error, customMessage string) {
	if err != nil {
		log.Fatalf("Error %s: %s\n", customMessage, err)
	}
}


func main () {
	result, errParser := internal.ArtifactParser("/home/lchan/Documents/formations/projet_go/ForensicZipGo/test_directory/windows.yaml", "FILE")
	handleError(errParser, "Error while Parsing file\n")

	fmt.Print(result)
}
