package structs_test

import (
	"fmt"
	"github.com/freesinger/go4aride/test"
	"github.com/freesinger/go4aride/tools"
	"testing"
)

func TestStackOpts(t *testing.T) {
	var s = tools.NewCapStack(20)

	for _, e := range test.TestArray {
		s.Push(e)
	}
	size := s.Size()
	fmt.Println(size)
	fmt.Printf("%d\n", s.Peek())
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", s.Pop())
	}
}
