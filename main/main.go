package main

import (
	"fmt"
	"github.com/freesinger/go4aride/structs"
)

var testArray = []int{4, 123, 13, 214, 33, -1, 344, 222, 3, 9, 220}
var demo = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	/*
		Test for create/remove/trace single list
	*/
	//head := tools.CreateLinkedList(testArray)
	//head = list.RemoveKth(head, 4)
	//tools.TraceList(head)

	/*
		Test for construct/trace btree
	*/
	//root := tools.ConstructBTree(testArray)
	//tools.TraceBTree(root)

	/*
		Test for stack opts
	*/
	var s = structs.NewCapStack(20)
	for _, e := range testArray {
		s.Push(e)
	}
	size := s.Size()
	fmt.Println(size)
	fmt.Printf("%d\n", s.Peek())
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", s.Pop())
	}

	//var q = structs.NewCapQueue(20)
	//for _, e := range testArray {
	//	q.Offer(e)
	//}
	//size := q.Size()
	//fmt.Println(size)
	//q.Dequeue()
	//fmt.Printf("%d\n", q.Peek())
	//size = q.Size()
	//for i := 0; i < size; i++ {
	//	fmt.Printf("%d ", q.Dequeue())
	//}
 }
