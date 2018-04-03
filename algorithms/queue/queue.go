package queue

import (
	"errors"
	"fmt"
	"reflect"
)

type Element interface {
}

type Queue struct {
	array     []Element
	queueType reflect.Type
}

func New() *Queue {
	q := Queue{}
	return &q
}

func (q *Queue) IsEmpty() bool {
	return q.Lenght() == 0
}

func (q *Queue) Lenght() int {
	return len(q.array)
}

func (q *Queue) Capacity() int {
	return cap(q.array)
}

func (q *Queue) Enqueue(element Element) {
	if q.IsEmpty() && q.queueType == nil {
		q.queueType = reflect.TypeOf(element)
	}
	if elementType := reflect.TypeOf(element); elementType != q.queueType {
		msg := "Incompatible type "
		msg += elementType.String()
		panic(msg)
	}
	q.array = append(q.array, element)
}

func (q *Queue) Dequeue() (*Element, error) {
	if q.IsEmpty() {
		return nil, errors.New("Cannot dequeue, queue is empty")
	}
	var el Element
	el, q.array = q.array[0], q.array[1:]
	return &el, nil
}

func (q *Queue) Info() {
	fmt.Printf("len = %d, cap = %d\n", q.Lenght(), q.Capacity())
}
