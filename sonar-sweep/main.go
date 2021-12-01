package main

import (
	"advent-of-code-21/sonar-sweep/pkg/filereader"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {

	fileReader, err := filereader.NewFileReader("./input.txt", make(chan string))
	if err != nil {
		log.Fatal(err)
	}
	increases := CountIncreases(fileReader)
	log.Infof("Number of Increases: %d", increases)
}

func CountIncreases(fileReader filereader.FileReader) int {
	go fileReader.ReadAllToCh()
	increases := 0
	previousValue := -1
	for {
		select {
		case line := <- fileReader.OutCh:
			depthValue, err := strconv.Atoi(line)
			log.Infof("Depth: %dm", depthValue)
			if err != nil {
			}
			if previousValue == -1 {
				previousValue = depthValue
				break
			}
			if previousValue < depthValue {
				increases++
			}
			previousValue = depthValue
		case <- fileReader.DoneCh:
			log.Info("Done reading inputs")
			return increases
		}
	}
}