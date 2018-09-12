import "sort"
func twoSum(nums []int, target int) []int {
    nums2 := make([][]int, 0)
    for index, v := range nums {
        nums2 = append(nums2, []int{index, v})
    }
    sort.Slice(nums2, func(a, b int) bool {
        return nums2[a][1] < nums2[b][1]
    })
    i:=0
    j := len(nums2)-1
    
    for (i<j) {
        if nums2[i][1] + nums2[j][1] == target {
            return []int {nums2[i][0], nums2[j][0]}
        } else if nums2[i][1] + nums2[j][1] < target {
            i++
        } else {
            j--
        }
    }
    return []int{}
}
