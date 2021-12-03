package ringbuffer

import "fmt"

type RingBuffer struct {
	buffer []int
	index int
}

func NewRingBuffer(cap int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, 0, cap),
		index: 0,
	}
}

func (r *RingBuffer) Set(value int) {
	if len(r.buffer) < cap(r.buffer) {
		r.buffer = append(r.buffer, value)
	} else {
		r.buffer[r.index] = value
	}
	r.index++
	if r.index == cap(r.buffer) {
		r.index = 0
	}
}

func (r *RingBuffer) IsFull() bool {
	return len(r.buffer) == cap(r.buffer)
}

func (r *RingBuffer) GetSum() int {
	sum := 0
	for _, value := range r.buffer {
		sum += value
	}
	return sum
}

func (r *RingBuffer) Dump() string {
	formattedString := "["
	for _, value := range r.buffer {
		formattedString += fmt.Sprintf("%d,", value)
	}
	formattedString = formattedString[:len(formattedString)-1]
	formattedString += "]"
	return formattedString
}