package list

import (
	"fmt"
	"github.com/freesinger/go4aride/structs"
	"syscall"
)

/**
Remove Kth number in a single list
 */
func RemoveKth(head *structs.ListNode, k int) *structs.ListNode {
	if k <= 0 {
		fmt.Println("Invalid k")
		syscall.Exit(-1)
	}

	pre, fast, dummy := &structs.ListNode{}, &structs.ListNode{}, &structs.ListNode{}
	pre.Next = head
	fast = head
	dummy.Next = head
	k--

	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast.Next != nil {
		fast = fast.Next
		head = head.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next

	return dummy.Next
}
