package queue

import "errors"

type queue struct {
	h,count,size int
	buf []interface{}
}

var ErrQueueFull = errors.New("Queue: queue is full")
var ErrQueueEmpty = errors.New("Queue: queue is empty")

func NewQueue(maxsize int) *queue {
	return &queue{
		h : 0,
		count : 0,
		buf : make([]interface{}, maxsize),
		size : maxsize,
	}
}

func (this *queue) Push(data interface{}) error {
	if this.count == this.size {
		return ErrQueueFull
	}
	this.count++
	this.buf[(this.h+this.count)%this.size] = data
	return nil
}

func (this *queue) Pop() (interface{},error) {
	rect,err := this.Peek()
	if this.count > 0 {
		this.h++
		if this.h == this.size {
			this.h = 0
		}
		this.count--
	}
	return rect, err
}

func (this *queue) Peek() (interface{}, error) {
	if this.count == 0 {
		return nil, ErrQueueEmpty
	}
	return this.buf[this.h], nil
}

	
