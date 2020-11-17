package structs

import "container/list"

/**
list容器实现的queue
 */
type Queue struct {
	container *list.List
	capacity  int
}

func NewQueue() *Queue {
	return NewCapQueue(-1)
}

func NewCapQueue(cap int) *Queue {
	return &Queue{
		container: list.New(),
		capacity:  cap,
	}
}

func (q *Queue) Offer(item interface{}) bool {
	// TODO: 动态扩容
	if q.capacity < 0 || q.container.Len() < q.capacity {
		q.container.PushBack(item)
		return true
	}

	return false
}

func (q *Queue) Add(item interface{}) {
	q.Offer(item)
}

func (q *Queue) Enqueue(item interface{}) {
	q.Offer(item)
}

func (q *Queue) Dequeue() interface{} {
	var head interface{} = nil
	var itemToDelete *list.Element = nil

	itemToDelete = q.container.Front()
	if itemToDelete != nil {
		head = q.container.Remove(itemToDelete)
	}

	return head
}

func (q *Queue) Size() int {
	return q.container.Len()
}

func (q *Queue) Peek() interface{} {
	var firstItem *list.Element = nil

	if q.container.Front() != nil {
		firstItem = q.container.Front()
	}

	return firstItem.Value
 }
