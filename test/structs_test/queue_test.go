package structs_test

import (
	"fmt"
	"github.com/freesinger/go4aride/test"
	"github.com/freesinger/go4aride/tools"
	"testing"
)

func TestQueueOpts(t *testing.T) {
	var q = tools.NewCapQueue(20)
	for _, e := range test.TestArray {
		q.Offer(e)
	}
	size := q.Size()
	fmt.Println(size)
	q.Dequeue()
	fmt.Printf("%d\n", q.Peek())
	size = q.Size()
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", q.Dequeue())
	}
}