package main

import (
	"fmt"
	"go4aride/algo"
)

var testArray = []int{4, 123, 13, 214, 33, 344, 222, 3, 9, 220}

func main() {
	fmt.Println("Index:", algo.BinarySearch(testArray, 222))
}
