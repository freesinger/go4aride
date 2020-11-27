package tools

import (
	"github.com/freesinger/go4aride/algo/list"
	"github.com/freesinger/go4aride/consts"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	head := CreateLinkedList(consts.TestArray)
	head = list.RemoveKth(head, 4)
	TraceList(head)
}
