package concurrent

import (
	"fmt"
	"time"
)

/*
channel的收和发必须成对出现
*/
func SentAndReceive() {
	c := make(chan int)
	// sent to channel
	// chan<- type
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second * 2)
		ch <- x * x
	}(c, 3)
	done := make(chan struct{})
	// receive from channel
	// <-chan type
	go func(ch <-chan int) {
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second * 2)
		done <- struct{}{} // sent
	}(c)
	<-done // receive
	// send and receive should be pair
	fmt.Println("bye")
}

func Football() {
	var ball = make(chan string)
	kickBall := func(playername string) {
		for ; ; {
			fmt.Println(<-ball, "passed ball")
			time.Sleep(time.Second)
			ball <- playername
		}
	}

	go kickBall("john")
	go kickBall("jack")
	go kickBall("ann")
	go kickBall("vickey")
	ball <- "Judge"
	var c chan float32
	<-c
}
