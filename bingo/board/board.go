package board

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Square struct {
	value int
	marked bool
}

type Board [][]Square

func NewBoard(inputStrings [][]string) Board {
	board := make([][]Square, 5)
	for i := range board {
		board[i] = make([]Square, 5)
	}
	for i, row := range inputStrings {
		for j, value := range row {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				log.Errorf("String value could not be converted to int %s", value)
			}
			board[i][j] = Square{
				value: intValue,
				marked: false,
			}
		}
	}
	return board
}

func (b Board) String() string {
	formattedString := ""
	for i, row := range b {
		rowString := ""
		for j := range row {
			rowString += fmt.Sprintf("%02d", b[i][j].value)
			if b[i][j].marked {
				rowString += "*"
			} else {
				rowString += " "
			}
			if j != len(row) - 1 {
				rowString += ", "
			}
		}
		formattedString += fmt.Sprintf("%s\n", rowString)
	}

	return formattedString
}

func (b Board) MarkNumber(n int) {
	for i, row := range b {
		for j := range row {
			if b[i][j].value == n {
				b[i][j].marked = true
			}
		}
	}
}

func (b Board) CheckWin() bool {
	// Check every row
	for i, row := range b {
		win := true
		for j := range row {
			if !b[i][j].marked {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	// Check every column
	for i, row := range b {
		win := true
		for j := range row {
			if !b[j][i].marked {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	return false
}

func (b Board) SumUnMarked() int {
	sum := 0
	for i, row := range b {
		for j := range row {
			if !b[i][j].marked {
				sum += b[i][j].value
			}
		}
	}
	return sum
}