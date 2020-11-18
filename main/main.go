package main

import (
	"fmt"
	"github.com/freesinger/go4aride/tools"
	"time"
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
	   Test for construct bst
	*/
	s := time.Now().Nanosecond()
	root := tools.ConstructBST(testArray, 0, len(testArray)-1)
	fmt.Printf("Elapsed (recursive): %d ms\n", (time.Now().Nanosecond() - s) / 1000)

	s = time.Now().Nanosecond()
	root = tools.ConstructBTree(testArray)
	fmt.Printf("Elapsed (level): %d ms\n", (time.Now().Nanosecond() - s) / 1000)

	tools.TraceBTree(root)
	tools.TraceBtreeUsingQueue(root)

	/*
		Test for stack opts
	*/
	//var s = structs.NewCapStack(20)
	//for _, e := range testArray {
	//	s.Push(e)
	//}
	//size := s.Size()
	//fmt.Println(size)
	//fmt.Printf("%d\n", s.Peek())
	//for i := 0; i < size; i++ {
	//	fmt.Printf("%d ", s.Pop())
	//}

	/*
		Test for queue opts
	*/
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
