package structs_test

import (
	"github.com/freesinger/go4aride/algo/list"
	"github.com/freesinger/go4aride/test"
	"github.com/freesinger/go4aride/tools"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	head := tools.CreateLinkedList(test.TestArray)
	head = list.RemoveKth(head, 4)
	tools.TraceList(head)
}
