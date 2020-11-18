package concurrent

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func ConPrint(content string, times int) {
	for i := 0; i < times; i++ {
		log.Println(content)
		d := time.Second * time.Duration(rand.Intn(5))
		time.Sleep(d)
	}
	wg.Done()
}

func Print() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2)
	go ConPrint("Shane", 10)
	go ConPrint("Ann", 10)
	wg.Wait()
}
