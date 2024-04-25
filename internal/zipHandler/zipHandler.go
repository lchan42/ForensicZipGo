package zipHandler

import (
	"os"
	"archive/zip"
	"hash"

)

// recursivly look for fine name that are in artefactMap and copy them into destination file
// func recursive(zipFile *zip.ReadCloser, artefactMap map[string]bool, hasherFunc func() hash.Hash) {

// }

func CopySpecificFiles(zipFile *zip.ReadCloser, artefactMap map[string]bool, hasherFunc func() hash.Hash) error {

	// create destination file
	outFile, err := os.Create("filtered_output.zip")
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := zip.NewWriter(outFile)
	defer writer.Close()

	return nil
}
