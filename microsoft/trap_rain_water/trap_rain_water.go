package main

func trap(height []int) int {
	//keep two pointers - left and right.
	//keep two values - leftMaxHeight and rightMaxHeight
	//out of left and right, whichever has smaller tower, do the following for it:
	// if it is left. Chek if value at left is greater than leftMaxValue.
	// if so, then we found a new max on left side. Also, it being tallest, wont store water.
	// else it will store leftMaxValue - height[left] amount of water
	// similarly for right..
	//
	//when left and right cross, you have seen all the input.
	total := 0

	left := 0
	right := len(height) -1

	if left == right || len(height) == 0 {
		return 0
	}

	leftMaxHeight := height[left]
	rightMaxHeight := height[right]

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMaxHeight {
				leftMaxHeight = height[left]
			} else {
				total += leftMaxHeight - height[left]
			}
			left++
		} else {
			if height[right] > rightMaxHeight {
				rightMaxHeight = height[right]
			} else {
				total += rightMaxHeight - height[right]
			}
			right--
		}
	}
	return total
}
