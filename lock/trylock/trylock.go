package main


import "sync"

// Lock -> try lock struct
type Lock struct {
	c chan struct{}
}

// NewLock -> generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock try lock and return lock result
func (l Lock) Lock() bool {
	res := false
	select {
	case <-l.c:
		res = true
	default:
	}
	return res
}

// Unlock -> Unlock the try lock
func (l Lock) Unlock() {
    l.c <- struct{}{}
}

var counter int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			println("current counter", counter)
			l.Unlock()
		}()
	}
    wg.Wait()
    print(counter)
}
