package task2

import "sync"

var lock sync.Mutex

func incrementNumber(num *int, delta int) {
	lock.Lock()
	defer lock.Unlock()
	*num += delta
}
