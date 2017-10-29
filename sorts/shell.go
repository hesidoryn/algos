package main

func shell(arr []int) {
	l := len(arr)
	if l < 2 {
		return
	}

	div := 3
	for d := l / div; d > 0; d /= div {
		for i := 0; i < l-d; i++ {
			for j := i; j >= 0 && arr[j+d] < arr[j]; j -= d {
				arr[j+d], arr[j] = arr[j], arr[j+d]
			}
		}
	}

	return
}
