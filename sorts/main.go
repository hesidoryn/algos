package main

import "fmt"

func main() {
	ins := []int{13, 12, 3, 4, 1123, 321, 32, 432, 65, 8}
	fmt.Println("Unsorted array:", ins)
	insertion(ins)
	fmt.Println("Sorted array:", ins)

	bub := []int{13, 12, 3, 4, 1123, 321, 32, 432, 65, 8}
	fmt.Println("Unsorted array:", bub)
	bubble(bub)
	fmt.Println("Sorted array:", bub)

	sel := []int{13, 12, 3, 4, 1123, 321, 32, 432, 65, 8}
	fmt.Println("Unsorted array:", sel)
	selection(sel)
	fmt.Println("Sorted array:", sel)

	sh := []int{13, 12, 3, 4, 1123, 321, 32, 432, 65, 8}
	fmt.Println("Unsorted array:", sh)
	shell(sh)
	fmt.Println("Sorted array:", sh)

	mer := []int{13, 12, 3, 4, 1123, 321, 32, 432, 65, 8}
	fmt.Println("Unsorted array:", mer)
	res := merge(mer)
	fmt.Println("Sorted array:", res)
}
