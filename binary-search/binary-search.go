package binarysearch

func Binary(haystack []int, value int) int {
	var lo int = 0
	var hi int = len(haystack) - 1

	var m int = 0
	for next := true; next; next = m < hi {

		m := lo + (hi-lo)/2

		mValue := haystack[m]
		if mValue == value {
			return m
		}

		if mValue < value {
			lo = m + 1
		} else {
			hi = m - 1
		}
	}

	return -1
}
