package task2

import "fmt"

func printOddNumber() {
	for i := 1; i < 10; i++ {
		if (i % 2) == 1 {
			fmt.Printf("%d ", i)
		}
	}
}

func printEvenNumber() {
	for i := 1; i < 10; i++ {
		if (i % 2) == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
