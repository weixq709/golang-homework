package task2

import "fmt"

func producer(num int, ch chan<- int) {
	for i := 0; i < num; i++ {
		ch <- i
		fmt.Printf("producer %d done\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Printf("consumer: %d\n", v)
	}
}
