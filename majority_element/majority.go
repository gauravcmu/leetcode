func majorityElement(nums []int) int {
	l := len(nums)
	maj_len := l / 2

	mapping := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := mapping[v]; ok {
			mapping[v] += 1
		} else {
			mapping[v] = 1
		}
		if mapping[v] > maj_len {
			return v
		}
	}
	return -1
}