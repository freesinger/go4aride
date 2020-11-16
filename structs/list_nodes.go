package structs

/**
ListNode
*/
type ListNode struct {
	Next *ListNode
	Val  int
}

/**
Double linked ListNode
*/
type DLinkedListNode struct {
	Prev *DLinkedListNode
	Next *DLinkedListNode
	Val  int
}
