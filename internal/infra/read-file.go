package infra

import (
	"bufio"
	"log"
	"os"
)

func ReadFileStream(fileName string, rowProcessor func(row string)) bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(file)

	scanner := bufio.NewScanner(file)

	// skip first line
	scanner.Scan()
	for scanner.Scan() {
		rowProcessor(scanner.Text())
	}

	return true
}
