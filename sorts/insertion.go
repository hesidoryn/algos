package main

func insertion(arr []int) {
	l := len(arr)
	if l < 2 {
		return
	}

	for i := 0; i < l-1; i++ {
		for j := i; j >= 0 && arr[j+1] < arr[j]; j-- {
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}

	return
}
