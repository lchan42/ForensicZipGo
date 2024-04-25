package parser

import "flag"

//Flag -adp, -ha and -z stand respectively for artefact definition path, hash algorithm and zip
var (
	DefinitionPath	= flag.String("adp", "../../configs/windowsArtifact.yaml", "[Optional] path to the artefact_definition.yaml")
	HashAlg			= flag.String("ha", "SHA256", "[Optional] Digest algorithm to use. default sha256")
    ArchivePath 	= flag.String("z", "../../test_directory/test.zip", "[Mandatory] Path to store the final ZIP archive")
	// You can define other flags here...
)
// ParseFlags uses flag.Parse() from "flag" pkg to parse the command line input. It returns an error if anything went wrong
func ParseFlags() error {
	flag.Parse()
	//if ArchivePath == nil {
	// 	return error
	// }
	return nil
	// todo add flag parsing logic
}
