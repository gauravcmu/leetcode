import (
	"errors"
	"fmt"
)

func findMin(nums []int) int {
	n, err := findMinHelper(nums, 0, len(nums)-1)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	return n
}

//[3,4,5,1,2]
// if mid less than end element, go right.
// else go left.
func findMinHelper(nums []int, start, end int) (int, error) {
	mid := start + (end-start)/2
	fmt.Printf("nums:%v start:%v mid:%v end:%v\n", nums, start, mid, end)

	if start == end {
		return nums[start], nil
	}

	if (mid == 0 && nums[mid] < nums[mid+1]) || (mid > 0 && nums[mid] < nums[mid-1]) {
		return nums[mid], nil
	}

	if nums[mid] < nums[end] {
		return findMinHelper(nums, start, mid-1)
	} else if nums[mid] > nums[end] {
		return findMinHelper(nums, mid+1, end)
	}
	return -1, errors.New("failed")
} 