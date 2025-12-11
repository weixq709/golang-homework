package task2

import (
	"sync"
	"testing"
)

func TestBasicChannel(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		producer(10, ch)
	}()
	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		producer(100, ch)
	}()
	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()
}
