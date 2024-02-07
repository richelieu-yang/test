package testUtil

func GetNumber() (num int) {
	s := []int{0, 1, 2, 3}

	for _, ele := range s {
		go func() {
			num += ele
		}()
	}
	return
}
