package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan string, 10)

	for i := 0; i < 10; i++ {
		fmt.Println("posted to c1...")
		c1 <- i
	}

	go worker(c1, c2)
	go worker(c1, c2)
	go worker(c1, c2)
	go worker(c1, c2)
	go worker(c1, c2)

	for i := 0; i < 10; i++ {
		msg := <-c2
		fmt.Println(msg)
	}

}

func worker(read <-chan int, post chan<- string) {
	for j := 0; j < 10; j++ {
		num := <-read
		str := ""
		for i := num; i < 10; i++ {
			str += fmt.Sprint(", ", i)
		}
		post <- str
    time.Sleep(1 * time.Second)
	}

}
