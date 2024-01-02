package linearsearch

func Linear(haystack []int, value int) bool {
	// go linear throw array
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == value {
			return true
		}
	}

	return false
}
