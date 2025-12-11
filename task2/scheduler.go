package task2

import (
	"fmt"
	"time"
)

type Scheduler struct{}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Execute(taskName string, task func()) {
	go func() {
		begin := time.Now()
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("[%s] panic occurred: %v \n", taskName, err)
				return
			}
			useTime := time.Now().Sub(begin)
			fmt.Printf("Execute took %fs \n", useTime.Seconds())
		}()

		task()
	}()
}
