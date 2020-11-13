package algo

import (
	"fmt"
	"sort"
)

func BinarySearch(a []int, t int) int {
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		left, right := 0, len(a)-1
		mid := left + (right-left)/2
		if a[mid] > t {
			right = mid - 1
		} else if a[mid] < t {
			left = mid + 1
		} else {
			return i
		}
	}
	_ = fmt.Errorf("Target element not exist!\n")
	return -1
}
