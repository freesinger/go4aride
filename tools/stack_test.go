package tools

import (
	"fmt"
	"github.com/freesinger/go4aride/consts"
	"testing"
)

func TestStackOpts(t *testing.T) {
	var s = NewCapStack(20)

	for _, e := range consts.TestArray {
		s.Push(e)
	}
	size := s.Size()
	fmt.Println(size)
	fmt.Printf("%d\n", s.Peek())
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", s.Pop())
	}
}
