package main

import (
	"advent-of-code-21/bingo/board"
	"advent-of-code-21/pkg/filereader"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

func main() {
	fileReader, err := filereader.NewFileReader("./input.txt", make(chan string))
	if err != nil {
		log.Fatal(err)
	}
	pickedSquares, boardInputs := processInput(fileReader)
	var boards []board.Board
	for _, boardInput := range boardInputs {
		boards = append(boards, board.NewBoard(boardInput))
	}
	log.Infof("Inputs: %v", pickedSquares)
	winningBoardIndex, finalNumber := play(boards, pickedSquares)
	log.Infof("Board %d wins! \n%s\nScore: %d", winningBoardIndex, boards[winningBoardIndex], boards[winningBoardIndex].SumUnMarked() * finalNumber)
}

func processInput(fileReader filereader.FileReader) ([]int, [][][]string) {
	go fileReader.ReadAllToCh()
	var pickedSquares []int
	var boardInputs [][][]string
	inputIndex := 0
	var partialBoardInput []string

	for {
		select {
			case line := <- fileReader.OutCh:
				if line == "" {
					if inputIndex != 0 {
						boardInput := processBoardInput(partialBoardInput)
						boardInputs = append(boardInputs, boardInput)
					}
					partialBoardInput = []string{}
					inputIndex++
					continue
				}
				if inputIndex == 0 {
					pickedSquaresStrings := strings.Split(line, ",")
					for _, s := range pickedSquaresStrings {
						num, err := strconv.Atoi(s)
						if err != nil {
							log.Errorf("Error converting input string to int %s", s)
							continue
						}
						pickedSquares = append(pickedSquares, num)
					}
					continue
				}
				partialBoardInput = append(partialBoardInput, line)
			case <- fileReader.DoneCh:
				if len(partialBoardInput) > 0 {
					boardInput := processBoardInput(partialBoardInput)
					boardInputs = append(boardInputs, boardInput)
				}
				return pickedSquares, boardInputs
		}
	}
}

func processBoardInput(rawInputStrings []string) [][]string {
	var boardInput [][]string
	for _, inputLine := range rawInputStrings {
		splitInputRow := strings.Split(inputLine, " ")
		var filteredInputRow []string
		for _,input := range splitInputRow {
			if input != "" {
				filteredInputRow = append(filteredInputRow, input)
			}
		}
		boardInput = append(boardInput, filteredInputRow)
	}
	return boardInput
}

func play(boards []board.Board, inputs []int) (int, int) {
	winningBoards := make(map[int]struct{})
	for _, pickedNumber := range inputs {
		for i, b := range boards {
			b.MarkNumber(pickedNumber)
			if b.CheckWin() {
				log.Infof("Board %d won! Number called %d", i, pickedNumber)
				winningBoards[i] = struct{}{}
				if len(winningBoards) == len(boards) {
					return i, pickedNumber
				}
			}
		}
	}
	return 0, 0
}
