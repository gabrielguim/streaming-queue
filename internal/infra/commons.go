package infra

import (
	"log"
	"os"
)

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Printf("failed to close file: %v\n", err)
	}
}
