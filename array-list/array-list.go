package arraylist

import (
	"errors"
	"fmt"
	"math"
)

type ArrayList struct {
	size int
	cap  int
	list []int
}

func NewArrayList(cap int) *ArrayList {
	if cap == 0 {
		cap = 10
	}

	list := make([]int, cap, cap)

	return &ArrayList{
		size: 0,
		cap:  cap,
		list: list,
	}
}

func (a *ArrayList) Size() int {
	return a.size
}

func (a *ArrayList) Clear() {
	a.list = nil
	a.list = make([]int, a.cap, a.cap)
	a.size = 0
}

func (a *ArrayList) Get(idx int) (int, error) {
	if idx > a.size-1 {
		return 0, errors.New("index not defined")
	}

	return a.list[idx], nil
}

func (a *ArrayList) Prepend(item int) {
	if a.size == a.cap {
		a.Grow()
	}

	a.size++

	// copy it on top instead of iterating. test this on how big the performance impact is between implementations
	// copy(a.list[1:], a.list)
	// a.list[0] = item

	for i := a.size - 1; i > 0; i-- {
		a.list[i] = a.list[i-1]
	}

	a.list[0] = item
}

func (a *ArrayList) Append(item int) {
	if a.size == a.cap {
		a.Grow()
	}

	a.size++
	a.list[a.size-1] = item
}

func (a *ArrayList) Grow() {
	increased := make([]int, a.size*2, a.cap*2)
	copy(increased, a.list)
	a.list = increased
	a.cap = len(increased)
}

func (a *ArrayList) Shrink() {
	smallerSize := int(math.RoundToEven(float64(a.cap / 4 * 3)))

	smaller := make([]int, smallerSize, smallerSize)

	copy(smaller, a.list)
	a.list = smaller
	a.cap = len(smaller)
}

func (a *ArrayList) Remove(idx int) (int, error) {
	if idx > a.size-1 {
		return 0, errors.New("index not defined")
	}

	removed := a.list[idx]

	for i := idx; i < a.size-1; i++ {
		a.list[i] = a.list[i+1]
	}

	//set last value after moving values to 0
	a.list[a.size-1] = 0

	// check if we could shrink the array back down
	// is difference between cap and size equal or more than 25%?
	if float64(a.cap-a.size)/float64(a.cap)*100 > 25 {
		fmt.Print(a.list)
		a.Shrink()
		fmt.Print(a.list)
	}

	a.size--

	return removed, nil
}
