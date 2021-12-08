package main

import (
	"advent-of-code-21/pkg/binaryconvert"
	"advent-of-code-21/pkg/filereader"
	log "github.com/sirupsen/logrus"
)

func main() {
	fileReader, err := filereader.NewFileReader("./input.txt", make(chan string))
	if err != nil {
		log.Fatal(err)
	}
	digitArrays := createArrays(fileReader)
	gammaArray := calculateGammaArray(digitArrays)
	gammaRate := binaryconvert.BigEndianSliceToInt(gammaArray)
	log.Infof("gammaRate: %d",gammaRate)
	epsilonArray := calculateEpsilonArray(gammaArray)
	epsilonRate := binaryconvert.BigEndianSliceToInt(epsilonArray)
	log.Infof("epsilonRate: %d", epsilonRate)

	powerConsumption := gammaRate * epsilonRate

	log.Infof("Total Power Consumption: %d", powerConsumption)

	var copyArrays [][]int
	for _, innerArray := range digitArrays {
		copyInnerArray := make([]int, len(innerArray))
		copy(copyInnerArray, innerArray)
		copyArrays = append(copyArrays, copyInnerArray)
	}

	oxygenArray := calculateLifeSupportArray(digitArrays, true)
	oxygenValue := binaryconvert.BigEndianSliceToInt(oxygenArray)
	log.Infof("Oxygen Value %d", oxygenValue)

	co2Array := calculateLifeSupportArray(copyArrays, false)
	co2Value := binaryconvert.BigEndianSliceToInt(co2Array)
	log.Infof("CO2 Value %v", co2Value)

	lifeSupportRating := oxygenValue * co2Value

	log.Infof("Total Life Support Rating: %d", lifeSupportRating)

}

func createArrays(fileReader filereader.FileReader) [][]int {
	go fileReader.ReadAllToCh()
	var digitArrays [][]int
	for {
		select {
		case line := <-fileReader.OutCh:
			charArray := []rune(line)
			for index, char := range charArray {
				if char != '1' && char != '0' {
					log.Errorf("Invalid char: %v", char)
					continue
				}
				if len(digitArrays)-1 < index {
					digitArrays = append(digitArrays, []int{})
				}
				digitArrays[index] = append(digitArrays[index], int(char) - '0')
			}
		case <-fileReader.DoneCh:
			return digitArrays
		}
	}
}

func calculateGammaArray(digitArrays [][]int) []int {
	var gammaArray []int
	for _, digitArray := range digitArrays {
		sum := 0
		for _, digit := range digitArray {
			sum += digit
		}
		if sum > len(digitArray)/2 {
			gammaArray = append(gammaArray, 1)
		} else {
			gammaArray = append(gammaArray, 0)
		}
	}
	return gammaArray
}

func calculateEpsilonArray(gammaArray []int) []int {
	var epsilonArray []int
	for _, digit := range gammaArray {
		if digit == 1 {
			epsilonArray = append(epsilonArray, 0)
		} else {
			epsilonArray = append(epsilonArray, 1)
		}
	}
	return epsilonArray
}

func calculateLifeSupportArray(digitArrays [][]int, oxygen bool) []int {
	for i:=0; i<len(digitArrays); i++ {
		digitArray := digitArrays[i]
		if len(digitArray) == 1 {
			var finalValue []int
			for _, arrays := range digitArrays {
				finalValue = append(finalValue, arrays[0])
			}
			return finalValue
		}
		var sum float32 = 0
		mostCommonDigit := 1
		for _, digit := range digitArray {
			sum += float32(digit)
		}
		if sum < float32(len(digitArray))/float32(2) {
			mostCommonDigit = 0
		}
		if !oxygen {
			if mostCommonDigit == 1 {
				mostCommonDigit = 0
			} else {
				mostCommonDigit = 1
			}
		}

		digitArrays = trimArrays(digitArrays, i, mostCommonDigit)
	}
	if len(digitArrays[0]) == 1 {
		var finalValue []int
		for _, arrays := range digitArrays {
			finalValue = append(finalValue, arrays[0])
		}
		return finalValue
	}
	log.Errorf("Unable to narrow down the options to only one row after all arrays have been processed.  Final result %v", digitArrays)
	return []int{}
}

func trimArrays(digitArrays [][]int, column int, searchCriteria int) [][]int {
	for i := 0; i<len(digitArrays[column]); {
		if digitArrays[column][i] != searchCriteria {
			// Found a row to delete at i, loop over all columns and remove i
			for j:=0; j<len(digitArrays); j++ {
				digitArrays[j] = append(digitArrays[j][:i], digitArrays[j][i+1:]...)
			}
		} else {
			i++
		}
	}
	return digitArrays
}