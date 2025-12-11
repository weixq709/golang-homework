package task2

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	scheduler := NewScheduler()
	var wg sync.WaitGroup
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 5; i++ {
		t := rand.Intn(10)

		wg.Add(1)
		taskName := fmt.Sprintf("task%d", i+1)
		scheduler.Execute(taskName, func() {
			fmt.Printf("[%s] starting...\n", taskName)
			select {
			// 设置随机睡眠时间
			case <-time.After(time.Duration(t) * time.Second):
				fmt.Printf("[%s] finished \n", taskName)
				wg.Done()
			}
		})
	}

	fmt.Println("[main] waiting...")
	wg.Wait()
	fmt.Println("[main] done")
}
