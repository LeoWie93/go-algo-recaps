package binarysearch

func Binary(haystack []int, value int) int {
	var lo int = 0
	var hi int = len(haystack) - 1

	var i int = 0
	for next := true; next; next = i < hi {

		m := lo + (hi-lo)/2

		mValue := haystack[i]
		if mValue == value {
			return i
		}

		if mValue < value {
			lo = m + 1
		} else {
			hi = m - 1
		}

		i++
	}

	return -1
}
