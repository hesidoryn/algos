package main

func bubble(arr []int) {
	l := len(arr)
	if l < 2 {
		return
	}

	for i := 1; i < l-2; i++ {
		isSorted := true
		for k := 0; k < l-i; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}

	return
}
