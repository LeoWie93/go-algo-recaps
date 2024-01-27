package ringerbuffer2

import (
	"errors"
	"math/rand"
)

type Item struct {
	value int
}

func NewItem() *Item {
	return &Item{
		value: rand.Int(),
	}
}

func (item *Item) GetValue() int {
	return item.value
}

type RingerBuffer2 struct {
	head     int
	tail     int
	buffer   []*Item
	length   int
	overflow bool
}

func New(length int8, overflow bool) *RingerBuffer2 {
	buffer := make([]*Item, length)
	return &RingerBuffer2{
		head:     0,
		tail:     0,
		buffer:   buffer,
		length:   len(buffer),
		overflow: overflow,
	}
}

func (r *RingerBuffer2) Add(item *Item) error {
	nextHead := (r.head + 1) % r.length
	if !r.overflow && nextHead == r.tail {
		return errors.New("overflow")
	}

	r.buffer[r.head] = item
	r.head = nextHead

	return nil
}

func (r *RingerBuffer2) Read() (*Item, error) {
	currentValue := r.buffer[r.tail]

	if currentValue == nil {
		return nil, errors.New("empty")
	}

	r.buffer[r.tail] = nil

	nextTail := (r.tail + 1) % r.length
	if nextValue := r.buffer[nextTail]; nextValue != nil {
		r.tail = nextTail
	}

	return currentValue, nil
}
