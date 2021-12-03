package main

import (
	"advent-of-code-21/pkg/filereader"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var validCommandsMap = make(map[string]struct{})
var validCommands = []string{"forward","down","up"}

func main() {
	fileReader, err := filereader.NewFileReader("./input.txt", make(chan string))
	if err != nil {
		log.Fatal(err)
	}
	populateValidCommandsMap()
	depth, horizontal := CalculatePosition(fileReader)
	log.Infof("Final depth: %dm, Horizontal distance: %dm", depth, horizontal)
	log.Infof("Multiplied value: %d", depth*horizontal)
}

func populateValidCommandsMap() {
	for _, command := range validCommands {
		validCommandsMap[command] = struct{}{}
	}
}

func CalculatePosition(fileReader filereader.FileReader) (int, int) {
	go fileReader.ReadAllToCh()
	depth := 0
	distance := 0
	aim := 0
	for {
		select {
		case line :=  <- fileReader.OutCh:
			command, value, err := ProcessLine(line)
			if err != nil {
				log.Errorf("Error processing movement command line: %v", err)
				continue
			}
			switch command {
			case "forward":
				distance += value
				depth += value * aim
			case "down":
				aim += value
			case "up":
				aim -= value
			}
		case <- fileReader.DoneCh:
			return depth, distance
		}
	}
}

func ProcessLine(line string) (string, int, error) {
	commandPieces := strings.Split(line, " ")
	if len(commandPieces) != 2 {
		return "", 0, fmt.Errorf("unable to split command line")
	}
	command := strings.ToLower(commandPieces[0])
	if _, ok := validCommandsMap[command]; !ok {
		return "", 0, fmt.Errorf("invalid command: %s", command)
	}
	value, err := strconv.Atoi(commandPieces[1])
	if err != nil {
		return "", 0, fmt.Errorf("unable to convert value: %s", commandPieces[1])
	}
	return command, value, nil
}