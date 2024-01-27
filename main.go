package main

import (
	"fmt"
	"math/rand"

	arraylist "github.com/leowie93/go-alog-recap/array-list"
	binarysearch "github.com/leowie93/go-alog-recap/binary-search"
	linearsearch "github.com/leowie93/go-alog-recap/linear-search"
	ringbuffer "github.com/leowie93/go-alog-recap/ring-buffer"
	ringerbuffer2 "github.com/leowie93/go-alog-recap/ringer-buffer2"
	"github.com/leowie93/go-alog-recap/twocrystalballs"
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
	r := ringbuffer.New(7, true)

	r.Add(1)
	r.Add(2)
	r.Add(3)
	r.Add(4)
	r.Add(5)

	fmt.Println(r.Read())
	fmt.Println(r.Read())
	fmt.Println(r.Read())

	r.Add(6)
	r.Add(7)
	r.Add(8)
	r.Add(9)

	fmt.Println(r)
	fmt.Println(r.Add(10))
	fmt.Println(r)

	fmt.Println("")
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
}
