#Introduction

This is a simple ring queue written in golang as a component of the GoSpell.

`NewQueue(maxsize int)` returns a new queue, u should set up the size of this queue. And this size won't change.

`Push(data interface{}) error` will push an item into the queue. if queue is full, this function will return ErrQueueFull. Otherwise, nil.

`Pop() (interface{},error)` will return the 1st item in queue. if queue is empty, this function will return nil,ErrQueueEmpty. Otherwise, data,nil. Notice you may need type assertion x.(type) to convert the result into the type u need.