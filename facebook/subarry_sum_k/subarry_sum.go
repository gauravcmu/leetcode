import "fmt"

func subarraySum(nums []int, k int) int {
	sums := make([]int, 0)

	prev := 0
	count := 0
	for key, value := range nums {
		fmt.Printf("%v", key)
		sums = append(sums, prev+value)
		prev = sums[key]
		if value == k {
			count++
		}
	}
	fmt.Printf("sums: %v", sums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum_ij := sums[j] - sums[i] + nums[i]
			if sum_ij == k {
				count++
			}
		}
	}
	return count
}