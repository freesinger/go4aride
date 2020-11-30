package concurrent

import "fmt"

var naturals = make(chan int)
var squares = make(chan int)

func Square() {
	//bidirectionalPipe()
	unidirectionalPipe()
}

// 2
func counter(out chan<- int) {
	for x := 0; x <= 30; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func unidirectionalPipe() {
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// 1
func bidirectionalPipe() {
	go func() {
		for x := 0; x <= 20; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
