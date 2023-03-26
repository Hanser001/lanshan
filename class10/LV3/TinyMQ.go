package main

import (
	"fmt"
	"time"
)

type MessageQueue interface {
	Send(message interface{})

	Pull(size int, timeout time.Duration) []interface{}

	Size() int

	Capacity() int
}

type MyMessageQueue struct {
	queue    chan interface{}
	capacity int
}

func (mq *MyMessageQueue) Send(message interface{}) {
	select {
	case mq.queue <- message:
		fmt.Printf("recevie message")
	default:
		time.Sleep(1 * time.Second)
	}
}

func (mq *MyMessageQueue) Pull(size int, timeout time.Duration) []interface{} {
	ret := make([]interface{}, 8)
	for i := 0; i < size; i++ {
		select {
		case msg := <-mq.queue:
			ret = append(ret, msg)
		case <-time.After(timeout):
			return ret
		}
	}
	return ret
}

func (mq *MyMessageQueue) Size() int {
	return len(mq.queue)
}

func (mq *MyMessageQueue) Capacity() int {
	return mq.capacity
}

func NewMessageQueue(capacity int) MessageQueue {
	var mq MessageQueue
	mq = &MyMessageQueue{
		queue:    make(chan interface{}, capacity),
		capacity: capacity,
	}
	return mq
}
