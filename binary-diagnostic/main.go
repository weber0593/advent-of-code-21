package main

import (
	"advent-of-code-21/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func main() {
	fileReader, err := filereader.NewFileReader("./input.txt", make(chan string))
	if err != nil {
		log.Fatal(err)
	}


}

func createArrays(fileReader filereader.FileReader) [][]rune {
	var charArrays [][]rune
	for {
		select {
		case line := <-fileReader.OutCh:
			charArray := []rune(line)
			if len(charArray) != 5 {
				log.Errorf("Invalid line doesnt have 5 characters: %s", line)
			}
		case <-fileReader.DoneCh:
			return charArrays
		}
	}
}