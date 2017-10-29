package main

func merge(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}

	c := l / 2
	left := arr[:c]
	right := arr[c:]
	return mergeHelp(merge(left), merge(right))
}

func mergeHelp(l, r []int) []int {
	res := []int{}
	li, ri, reslen := 0, 0, len(l)+len(r)
	for i := 0; i < reslen; i++ {
		if li == len(l) {
			res = append(res, r[ri:]...)
			break
		}
		if ri == len(r) {
			res = append(res, l[li:]...)
			break
		}

		if l[li] > r[ri] {
			res = append(res, r[ri])
			ri++
			continue
		}

		if l[li] <= r[ri] {
			res = append(res, l[li])
			li++
			continue
		}
	}

	return res
}
