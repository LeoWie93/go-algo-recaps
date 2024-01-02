package twocrystalballs

import "math"

func TwoCrystalBalls(breaks []bool) int {
	var jumpAmount int = int(math.Round(math.Sqrt(float64(len(breaks)))))
	var index int = 0

	for ; index < len(breaks); index += jumpAmount {
		if breaks[index] {
			break
		}
	}

	index -= jumpAmount

	for ; index < len(breaks); index++ {
		if breaks[index] {
			return index
		}
	}

	return -1
}
