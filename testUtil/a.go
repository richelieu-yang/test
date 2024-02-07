package testUtil

import "sync"

func GetNumber() (num int) {
	s := []int{0, 1, 2, 3}

	var wg sync.WaitGroup
	for _, ele := range s {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num += ele
		}()
	}

	wg.Wait()
	return
}
