package structs

/**
Typical binary search tree node
 */
type BTreeNode struct {
	Left  *BTreeNode
	Right *BTreeNode
	Val   int
}


