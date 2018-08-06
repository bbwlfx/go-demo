package main

import "fmt"

var ch chan int = make(chan int)

func sum(a int, b int, c chan int) {
	c <- a + b
}
func main() {
	go sum(5, 10, ch)

	fmt.Println(<-ch)

}
