package algo

import (
	"fmt"
	"sort"
)

func BinarySearch(a []int, t int) int {
	maps := make(map[int]int)
	for i := 0; i < len(a); i++ {
		maps[a[i]] = i
	}
	sort.Ints(a)

	left, right := 0, len(a)-1
	for left < right {
		mid := left + (right-left)/2
		if a[mid] > t {
			right = mid - 1
		} else if a[mid] < t {
			left = mid + 1
		} else {
			return maps[a[mid]]
		}
	}
	fmt.Println("Target element not exist!")
	return -1
}
