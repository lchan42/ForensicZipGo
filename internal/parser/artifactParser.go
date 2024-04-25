package parser

import(
	"os"
	"gopkg.in/yaml.v2"
	"io"
)

// represent the sources type of each artifactDefinition
type Source struct {
	Type		string
	Attributes	map[string]interface{}
}

// represent an artifact struct following the convention stated in https://github.com/ForensicArtifacts/artifacts
type ArtifactDefinition struct {
	Name		string
	Doc			string
	Sources		[]Source
	SupportedOs	[]string
	URLS		[]string

}

// processSources checks if sources contains a type that matched typeToRetrive and returns a boolean accordinly
func processSources(typeToRetrive string, sources []Source) bool {
	for _, source := range sources {
		if source.Type == typeToRetrive {
			return true
		}
	}
	return false
}

// artifactParser parses artifact definitions from a YAML file provided by path and
// returns a map of artifact names that match the specified type.
func ParseArtifact(path string, typeToRetrieve string ) (map[string]bool, error){

	// open file
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Init map to store artifact names
	namesMap := make(map[string]bool)

	// Create YAML decoder for the file
	decoder := yaml.NewDecoder(file)

	// Loop Through artifact decoded
	for {
		var artifact ArtifactDefinition

		if err := decoder.Decode(&artifact); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if processSources(typeToRetrieve, artifact.Sources) {
			namesMap[artifact.Name] = true
		}
	}

	// Return map of artifact
	return namesMap, nil
}


