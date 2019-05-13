func moveZeroes(nums []int) {
	i := 0
	index_to_write_at := 0

	//find non-zero elements and write them to index_to_write_at
	for i < len(nums) {
		if nums[i] != 0 {
			nums[index_to_write_at] = nums[i]
			index_to_write_at++
		}

		i++
	}
	//fill up rest of the spots with zeros
	for i := index_to_write_at; i < len(nums); i++ {
		nums[i] = 0
	}
}