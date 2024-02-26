package infra

import (
	"bufio"
	"log"
	"os"
)

type Entry interface {
	ParseString() string
}

func WriteOnFile(fileName string, messages Entry) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer closeFile(file)

	if err != nil {
		log.Fatal(err)
	}

	fileWriter := bufio.NewWriter(file)

	fileWriter.WriteString(messages.ParseString())
	fileWriter.Flush()
}
