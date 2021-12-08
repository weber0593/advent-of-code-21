package binaryconvert

import (
	"math"
)

func BigEndianSliceToInt(slice []int) int {
	sum := 0
	power := 0
	for i:=len(slice)-1; i>=0; i-- {
		if slice[i] == 1 {
			sum += int(math.Pow(2, float64(power)))
		}
		power++
	}
	return sum
}
