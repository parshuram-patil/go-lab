package main

import (
	"fmt"
	"strconv"
	"time"
)

func IsClosed(ch <-chan string) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func reader(c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Reader - in count ", i)
		time.Sleep(time.Millisecond * 100)
	}

	/*	fmt.Println("in Reader part 1 = ", <-c)
		fmt.Println("in Reader part 2 = ", <-c)
		fmt.Println("in Reader part 3 = ", <-c)
	*/
	for msg := range c {
		fmt.Println("in reader ", msg)
	}

	fmt.Println("--->", IsClosed(c))
	fmt.Println("*** reader End")
}

func writer(c chan string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Writer - in count ", i)
		time.Sleep(time.Millisecond * 100)
		c <- "Writer count - " + strconv.Itoa(i)

	}

	fmt.Println("*** writer End")
	fmt.Println("--->", IsClosed(c))
	close(c)

}
func main() {
	c := make(chan string, 5)
	go writer(c)
	go reader(c)

	for {
	}
}
