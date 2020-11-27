package tools

import (
	"github.com/freesinger/go4aride/consts"
	"testing"
)

func TestConstructBST(t *testing.T) {
	root := ConstructBST(consts.TestArray, 0, len(consts.TestArray)-1)
	TraceBTree(root)
	TraceBtreeUsingQueue(root)
}

func TestConstructBTree(t *testing.T) {
	root := ConstructBTree(consts.TestArray)
	TraceBTree(root)
	TraceBtreeUsingQueue(root)
}
