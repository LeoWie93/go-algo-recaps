package main

import (
	"fmt"
	"math/rand"

	arraylist "github.com/leowie93/go-alog-recap/array-list"
	binarysearch "github.com/leowie93/go-alog-recap/binary-search"
	linearsearch "github.com/leowie93/go-alog-recap/linear-search"
	"github.com/leowie93/go-alog-recap/twocrystalballs"
)

func main() {
	haystack := []int{9, 14, 15, 2200}
	resLinear := linearsearch.Linear(haystack, 11)
	fmt.Printf("Linear search result: %t\n", resLinear)

	resBinary := binarysearch.Binary(haystack, 15)
	fmt.Printf("Binary search result: %d\n", resBinary)

	var breaks = make([]bool, 5000000)
	pattern := []bool{true}
	idx := rand.Intn(5000000)

	sliceToChange := breaks[idx:]
	copy(sliceToChange, pattern)

	for j := len(pattern); j < len(sliceToChange); j *= 2 {
		//j: => gets the slice after the given index
		//j: => gets the slice until the given index
		copy(sliceToChange[j:], sliceToChange[:j])
	}

	fmt.Println(idx)
	resBreaks := twocrystalballs.TwoCrystalBalls(breaks)
	fmt.Printf("TCB search result: %d\n", resBreaks)

	a := arraylist.NewArrayList(3)
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(5)
	a.Prepend(4)

	fmt.Println(a.Size())

	fmt.Println(a.Get(0))
	fmt.Println(a.Get(7))

	a.Append(2)
	a.Append(2)
	a.Append(3)
	a.Append(5)
	a.Append(3)
	a.Append(5)

	a.Remove(2)
	a.Remove(2)
	a.Remove(2)
	a.Remove(2)
	a.Remove(2)
	a.Remove(2)
	a.Remove(2)
	a.Remove(2)
	a.Remove(2)

	fmt.Println(a.Size())
	fmt.Println(a)

	a.Clear()

	fmt.Println(a)
}
