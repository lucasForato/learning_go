package main

import (
	"fmt"
	"time"
)

// since we are waiting for 3 async calls, we can batch them together into one execution for each and then we
// can use the result of all of them together.
func main() {
  start := time.Now()
	c := make(chan string, 3)
	msg := ""
	go firstParagraph(c)
	go secondParagraph(c)
	go thirdParagraph(c)

	// we have three asynchronous calls
	for i := 0; i < 3; i++ {
		msg += <-c
	}
  close(c)

	fmt.Println(msg)
  end := time.Now()
  fmt.Printf("Execution time %v\n", end.Sub(start))
}

// imagine this is asynchronous
func firstParagraph(c chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println("firstParagraph")
	c <- "This is the first paragraph."
}

// another asynchronous call
func secondParagraph(c chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println("secondParagraph")
	c <- "This is the second paragraph."
}

// third paragraph
func thirdParagraph(c chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println("thirdParagraph")
	c <- "This is the third paragraph."
}
