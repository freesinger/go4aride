package tools

import (
	"fmt"
	"github.com/freesinger/go4aride/consts"
	"testing"
)

func TestQueueOpts(t *testing.T) {
	var q = NewCapQueue(20)
	for _, e := range consts.TestArray {
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
