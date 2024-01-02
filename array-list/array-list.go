package arraylist

import (
	"errors"
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

	// copy it on top instead of interating. test this one day how big the impact is between implementations
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
}

func (a *ArrayList) Remove(idx int) (int, error) {
	if idx > a.size-1 {
		return 0, errors.New("index not defined")
	}

	removed := a.list[idx]

	for i := idx; i < a.size-1; i++ {
		a.list[i] = a.list[i+1]
	}

	a.list[a.size-1] = 0

	//check if we could shrink the array back down
	// is difference between cap and size equal or more than 25%?
	if (a.cap-a.size)/a.cap*100 > 25 {

	}

	a.size--

	return removed, nil
}
