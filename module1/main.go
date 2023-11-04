package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	q := make(chan int, 10)
	r := rand.New(rand.NewSource(99))
	for {
		producer(r.Int(), q)
		consume(q)
		time.Sleep(time.Second)
	}
}

func consume(q <-chan int) {
	i := <-q
	fmt.Println(i, time.Now().Format(time.ANSIC))
}

func producer(a int, q chan<- int) {
	q <- a
}
