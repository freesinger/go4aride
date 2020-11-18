package tools

import "container/list"

/*
使用list容器简单实现的stack
*/
type Stack struct {
	container *list.List
	capacity  int
}

func NewStack() *Stack {
	return NewCapStack(-1)
}

func NewCapStack(cap int) *Stack {
	return &Stack{
		container: list.New(),
		capacity:  cap,
	}
}

func (s *Stack) Push(item interface{}) bool {
	// TODO: 动态扩容
	if s.capacity < 0 || s.container.Len() < s.capacity {
		s.container.PushBack(item)
		return true
	}

	return false
}

func (s *Stack) Pop() interface{} {
	var item interface{} = nil
	var lastContainerItem *list.Element = nil

	lastContainerItem = s.container.Back()
	if lastContainerItem != nil {
		item = s.container.Remove(lastContainerItem)
	}

	return item
}

func (s *Stack) Peek() interface{} {
	var item *list.Element = nil

	if s.container.Len() > 0 {
		item = s.container.Back()
	}

	return item.Value
}

func (s *Stack) Size() int {
	return s.container.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.Size() <= 0
}