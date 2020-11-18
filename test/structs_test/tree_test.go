package structs_test

import (
	"github.com/freesinger/go4aride/test"
	"github.com/freesinger/go4aride/tools"
	"testing"
)

func TestConstructBST(t *testing.T) {
	root := tools.ConstructBST(test.TestArray, 0, len(test.TestArray)-1)
	tools.TraceBTree(root)
	tools.TraceBtreeUsingQueue(root)
}

func TestConstructBTree(t *testing.T) {
	root := tools.ConstructBTree(test.TestArray)
	tools.TraceBTree(root)
	tools.TraceBtreeUsingQueue(root)
}
