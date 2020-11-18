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
}
