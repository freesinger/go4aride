package main

import (
	"github.com/freesinger/go4aride/tools"
)

var testArray = []int{4, 123, 13, 214, 33, 344, 222, 3, 9, 220}

func main() {
	head := tools.CreateLinkedList(testArray)
	tools.TraceList(head)
}
