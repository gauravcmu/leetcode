package main

import (
	"fmt"
)

func main() {
	input := []int{38, 27, 43, 3, 9, 82, 10}
	mergesort(input)
	fmt.Printf("after sort: %v\n", input)
}

func mergesort(input []int) {
	low := 0
	high := len(input) - 1
	sort(&input, low, high)
}

func sort(input *[]int, low, high int) {
	fmt.Printf("sort: low: %v high: %v\n", low, high)
	if low < high {
		mid := (high + low) / 2
		fmt.Printf(" mid: %v\n", mid)
		sort(input, low, mid)
		sort(input, mid+1, high)

		merge(input, low, mid, high)
	}
}

func merge(input *[]int, low, mid, high int) {
	//copy to two arrays A and B.
	if low == high {
		return
	}
	
	fmt.Printf("Lets merge low: %v mid: %v high: %v\n", low, mid, high)
	A := make([]int, mid-low+1) // IMPORTANT - watch the lengths: from 0 to mid (including mid).
	B := make([]int, high-mid)  // from mid+1 to high - excludes mid.

	for i := 0; i < mid-low+1; i++ {
		A[i] = (*input)[low+i]
	}
	for i := 0; i < high-mid; i++ {
		B[i] = (*input)[mid+1+i]
	}

	fmt.Printf("A: %v B: %v\n", A, B)
	i := 0
	j := 0
	k := low // IMPORTANT: remember we are not writing at 0 index all the time. This is the input array.

	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			(*input)[k] = A[i]
			i++
		} else {
			(*input)[k] = B[j]
			j++
		}
		k++
	}

	// we came here that means either of the arrays A and B are exhausted. Or both.

	for i < len(A) {
		(*input)[k] = A[i]
		i++
		k++
	}

	for j < len(B) {
		(*input)[k] = B[j]
		j++
		k++
	}
}

