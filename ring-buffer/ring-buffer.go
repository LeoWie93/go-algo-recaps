package ringbuffer

import "errors"

// N is len of ring
// similar to arraylist. But everyting not between head and tail is nil\
// [0,0,head,5,20,2,13,tail,0,0]

// [0,0,4,0,0]
// [0,0,0,0,0]
// [0,0,4,1,0]

// can wrap around. tail%len gets wrap around index if N is reached
type Ringbuffer struct {
	start    int
	end      int
	buffer   []int
	overflow bool
	len      int
}

func New(lenght int, overflow bool) *Ringbuffer {
	ring := make([]int, lenght)
	return &Ringbuffer{
		0,
		0,
		ring,
		overflow,
		len(ring),
	}
}

func (r *Ringbuffer) Add(value int) error {
	if !r.overflow && (r.end+1)%r.len == r.start {
		return errors.New("overflow error")
	}

	r.buffer[r.end] = value
	r.end = (r.end + 1) % r.len

	return nil
}

func (r *Ringbuffer) Read() (int, error) {
	if r.start == r.end {
		return 0, errors.New("nothing to read")
	}

	item := r.buffer[r.start]
	r.start = (r.start + 1) % r.len

	return item, nil
}

// i get the feeling it is really similar to an array list. but with wrapping around
// special cases to take into account. tail not catching up to head = need to grow
// do some steps and print out current state for debugging, testing and showcasing
