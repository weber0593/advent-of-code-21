package main

import (
	"advent-of-code-21/pkg/filereader"
	"advent-of-code-21/pkg/ringbuffer"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {

	fileReader, err := filereader.NewFileReader("./input.txt", make(chan string))
	if err != nil {
		log.Fatal(err)
	}

	//increases := CountIncreases(fileReader)
	//log.Infof("Number of Increases: %d", increases)
	rollingIncreases := CountRollingAverageIncreases(fileReader)
	log.Infof("Number of Increases With Rolling Sum: %d", rollingIncreases)
}

func CountIncreases(fileReader filereader.FileReader) int {
	go fileReader.ReadAllToCh()
	increases := 0
	previousValue := -1
	for {
		select {
		case line := <- fileReader.OutCh:
			depthValue, err := strconv.Atoi(line)
			if err != nil {
				log.Errorf("Error converting file line to int: %v", err)
				continue
			}
			log.Infof("Depth: %dm", depthValue)
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

func CountRollingAverageIncreases(fileReader filereader.FileReader) int {
	go fileReader.ReadAllToCh()
	rb := ringbuffer.NewRingBuffer(3)
	increases := 0
	previousValue := -1
	for {
		select {
		case line := <- fileReader.OutCh:
			depthValue, err := strconv.Atoi(line)
			if err != nil {
				log.Errorf("Error converting file line to int: %v", err)
				continue
			}
			log.Infof("Depth: %dm", depthValue)
			rb.Set(depthValue)
			if rb.IsFull() {
				if previousValue == -1 {
					previousValue = rb.GetSum()
					break
				}
				if previousValue < rb.GetSum() {
					increases++
				}
				previousValue = rb.GetSum()
			}
		case <- fileReader.DoneCh:
			log.Info("Done reading inputs")
			return increases
		}
	}
}