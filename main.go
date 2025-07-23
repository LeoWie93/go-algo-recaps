package main

import (
	"fmt"
	"math/rand"

	arraylist "github.com/leowie93/go-alog-recap/array-list"
	binarysearch "github.com/leowie93/go-alog-recap/binary-search"
	linearsearch "github.com/leowie93/go-alog-recap/linear-search"
	quicksort "github.com/leowie93/go-alog-recap/quick-sort"
	ringerbuffer2 "github.com/leowie93/go-alog-recap/ringer-buffer2"
	twocrystalballs "github.com/leowie93/go-alog-recap/twocrystalballs"
)

func main() {
	// LinearSearch
	haystack := []int{9, 14, 15, 2200}
	resLinear := linearsearch.Linear(haystack, 11)
	fmt.Printf("Linear search result: %t\n", resLinear)

	// BinarySearch
	resBinary := binarysearch.Binary(haystack, 15)
	fmt.Printf("Binary search result: %d\n", resBinary)

	// TwoCrystalBalls simple example
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

	// ArrayList
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
	a.Remove(1)

	fmt.Println(a.Size())
	fmt.Println(a)

	a.Clear()

	fmt.Println(a)

	//RingBuffer
	fmt.Println("")
	fmt.Println("")
	fmt.Println("ringerbuffer2")

	r2 := ringerbuffer2.New(10, true)
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())

	item, _ := r2.Read()
	fmt.Println(item)
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())

	fmt.Println(r2)

	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())
	r2.Add(ringerbuffer2.NewItem())

	fmt.Println(r2)

	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())
	fmt.Println(r2.Read())

	fmt.Println(r2)

	testQuickSort()
}

func testQuickSort() {
	scrambled := rand.Perm(5000000)
	// fmt.Println(scrambled)

	quicksort.Sort(&scrambled)

	// fmt.Println(scrambled)
}
