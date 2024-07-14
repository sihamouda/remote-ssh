package storage

import (
	"io/fs"
	"log"
	"os"
)

var (
	// RS_PEM_FOLDER => folder that contains .pem files
	pemFolder, isPemFolderExist = os.LookupEnv("RS_PEM_FOLDER")
)

// Checks if env var RS_PEM_FOLDER exists
func validateEnv() {
	if isPemFolderExist {
		log.Printf(".pem files are located inside %s directory\n", pemFolder)
	} else {
		log.Fatal("Please export RS_PEM_FOLDER\n")
	}
}

// Checks if dir exists
func validatePemDir() []fs.DirEntry {
	pemFiles, err := os.ReadDir(pemFolder)
	if err != nil {
		log.Fatalf("No such directory: %s", pemFolder)
	}

	return pemFiles
}

func Main() []fs.DirEntry {
	validateEnv()
	return validatePemDir()
}
