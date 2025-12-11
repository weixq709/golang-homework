package task2

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func TestSafetyIncrNumber(t *testing.T) {
	num := 10
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				incrementNumber(&num, 1)
			}
		}()
	}

	wg.Wait()
	if num != 10010 {
		t.Errorf("num != 10010")
	}
}

func TestAtomicIncrNumber(t *testing.T) {
	num := 10
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64((*int64)(unsafe.Pointer(&num)), 1)
			}
		}()
	}
	wg.Wait()
	wg.Wait()
	if num != 10010 {
		t.Errorf("num != 10010")
	}
}
