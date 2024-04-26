package zipHandler

import (
	// "os"
	"archive/zip"
	"fmt"
	"hash"
	"io"
	"path/filepath"

	"os"
	"path"
)

func copyFileToZip(sourceFile *zip.File, destinationZip *zip.Writer) error {
	fmt.Printf("coping: %s...\n", sourceFile.Name)

	// Open the source file in the source zip
	sourceReader, err := sourceFile.Open()
	if err != nil {
		return fmt.Errorf("error opening source file %s: %w", sourceFile.Name, err)
	}
	defer sourceReader.Close()

	// Create a new file header for the destination zip
	fileHeader := &zip.FileHeader{
		Name:   sourceFile.Name,
		Method: sourceFile.Method,
		Extra: append([]byte(nil), sourceFile.FileHeader.Extra...),
		Flags: sourceFile.Flags,
		Modified: sourceFile.Modified,
	}

	if sourceFile.FileHeader.CreatorVersion >> 8 == 3 {
		// Unix-specific attributes
		fileHeader.SetMode(sourceFile.Mode())
		fileHeader.ExternalAttrs = sourceFile.ExternalAttrs
	}
		// add other os specific attribute here when error comes later ....

	//Create the new file in the destination zip with the updated header
	writer, err := destinationZip.CreateHeader(fileHeader)
	if err != nil {
		return fmt.Errorf("error creating destination file %s: %w", sourceFile.Name, err)
	}

	// Copy the content of the source file to the destination zip
	_, err = io.Copy(writer, sourceReader)
	if err != nil {
		return fmt.Errorf("error copying content to destination file %s: %w", sourceFile.Name, err)
	}

	return nil
}

//returns a fileName cleared from path and extention
//ex: /pathto/fileName.txt --> fileName
func getPureFileName(nameWithPath string) string {
	nameWithExt := path.Base(nameWithPath)

	return nameWithExt[:len(nameWithExt)-len(filepath.Ext(nameWithExt))]
}

func CopySpecificFiles(zipFile *zip.ReadCloser, artefactMap map[string]bool, hasherFunc func() hash.Hash) error {

	// create destination file
	outFile, err := os.Create("filtered_output.zip")
	if err != nil {
		return err
	}
	defer outFile.Close()

	// create a writer to add content to destination file
	writer := zip.NewWriter(outFile)
	defer writer.Close()


	for _, file := range zipFile.File {
		// if file name (cleared from path and extention) is in map
		if artefactMap[getPureFileName(file.Name)] {
			copyFileToZip(file, writer)
			// fmt.Printf("unknown = %v, file = %v\n", unknown, file.Name)
		}
	}

	return nil
}


//"path/filepath"
// func hasExtention(fileName, extention string) bool {
// 	return filepath.Ext(fileName) == extention
// }
