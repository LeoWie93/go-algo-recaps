package quicksort

func partition(arr *[]int, low int, high int) int {
	pivot := (*arr)[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if (*arr)[i] <= pivot {
			idx++
			tmp := (*arr)[idx]
			(*arr)[idx] = (*arr)[i]
			(*arr)[i] = tmp
		}
	}

	idx++
	(*arr)[high] = (*arr)[idx]
	(*arr)[idx] = pivot

	return idx
}

func qsort(arr *[]int, low int, high int) {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high)

	qsort(arr, low, pivot-1)
	qsort(arr, pivot+1, high)
}

func Sort(arr *[]int) {
	qsort(arr, 0, len(*arr)-1)
}
