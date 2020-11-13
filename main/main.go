package main

import (
	"fmt"
	"go4aride/algo"
)

var test_array = []int{4, 123, 13, 214, 33, 344, 222, 3, 9, 220}

func main() {
	fmt.Println(algo.BinarySearch(test_array, 33))
}
