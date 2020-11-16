package tools

import (
	"fmt"
	"github.com/freesinger/go4aride/structs"
	"syscall"
)

func ConstructBTree(a []int) *structs.BTreeNode {
	if len(a) == 0 {
		fmt.Println("Null array")
		syscall.Exit(-1)
	}
	root := &structs.BTreeNode{
		Left:  nil,
		Right: nil,
		Val:   a[0],
	}
	// TODO: add lch/rch
}
