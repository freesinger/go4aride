package main

import (
	"github.com/freesinger/go4aride/algo/list"
	"github.com/freesinger/go4aride/tools"
)

var testArray = []int{4, 123, 13, 214, 33, 344, 222, 3, 9, 220}

func main() {
	head := tools.CreateLinkedList(testArray)
	head = list.RemoveKth(head, 4)
	tools.TraceList(head)
}
