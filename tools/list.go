package tools

import (
	"fmt"
	"github.com/freesinger/go4aride/structs"
	"syscall"
)

/**
层序构建一个单链表
*/
func CreateLinkedList(a []int) *structs.ListNode {
	if len(a) == 0 {
		fmt.Println("Null array")
		syscall.Exit(-1)
	}
	head := &structs.ListNode{Val: -1}
	node := &structs.ListNode{Val: a[0]}
	head.Next = node
	for i := 1; i < len(a); i++ {
		nextNode := &structs.ListNode{Val: a[i]}
		node.Next = nextNode
		node = nextNode
	}

	return head
}

/**
遍历单链表
*/
func TraceList(head *structs.ListNode) {
	if head == nil {
		fmt.Println("Null list")
		syscall.Exit(-1)
	}
	next := head.Next
	for next != nil {
		fmt.Printf("%d ", next.Val)
		next = next.Next
	}
}
