package main

func selection(arr []int) {
	l := len(arr)
	if l < 2 {
		return
	}

	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return
}
