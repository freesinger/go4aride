package tools

import (
	"fmt"
	"github.com/freesinger/go4aride/structs"
	"syscall"
)

func ConstructBTreeRecurs(a []int, l, r int) *structs.BTreeNode {
	if len(a) == 0 {
		fmt.Println("Null array")
		syscall.Exit(-1)
	}
	if l > r {
		var node *structs.BTreeNode = nil
		return node
	}
	mid := l + (r-l)/2
	root := &structs.BTreeNode{
		Val: a[mid],
	}
	root.Left = ConstructBTreeRecurs(a, l, mid-1)
	root.Right = ConstructBTreeRecurs(a, mid+1, r)

	return root
}

/**
Construct a binary tree
-1 note nil child
*/
func ConstructBTree(a []int) *structs.BTreeNode {
	if len(a) == 0 {
		fmt.Println("Null array")
		syscall.Exit(-1)
	}

	// init all nodes
	nodes := make([]*structs.BTreeNode, 0, len(a))
	for _, i := range a {
		node := &structs.BTreeNode{
			Left:  nil,
			Right: nil,
			Val:   i,
		}
		// fmt.Printf("idx: %d, i: %d\n", idx, i)
		nodes = append(nodes, node)
	}
	// 层序构建二叉树
	for i := 0; i < len(a)/2; i++ {
		left, right := i*2+1, i*2+2
		root := nodes[i]
		if left < len(a) && a[left] != -1 {
			root.Left = nodes[left]
		}
		if right < len(a) && a[right] != -1 {
			root.Right = nodes[right]
		}
	}

	return nodes[0]
}

/**
层序遍历二叉树
用slice实现fifo queue
*/
func TraceBTree(root *structs.BTreeNode) []int {
	list := make([]int, 0)
	queue := make([]*structs.BTreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			curNode := queue[0]
			queue = queue[1:]
			size--
			list = append(list, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}
	fmt.Println(list)
	return list
}
