package task2

import (
	"fmt"
	"sync"
	"testing"
)

func TestPrintNumber(t *testing.T) {
	fmt.Println("starting...")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Print("odd numbers: ")
		printOddNumber()
		fmt.Println()
	}()

	go func() {
		defer wg.Done()
		fmt.Print("even numbers: ")
		printEvenNumber()
		fmt.Println()
	}()

	wg.Wait()
	fmt.Println("ended")
}
