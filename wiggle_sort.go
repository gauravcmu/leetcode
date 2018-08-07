//just ensure next element follows order
func wiggleSort(nums []int)  {
    sense := "smaller" //smaller - 0, bigger - 1
    for i:=0; i<len(nums);i++ {
        if sense == "smaller" {
            if i+1 < len(nums) {
                if nums[i] > nums[i+1] {
                    //swap
                    nums[i], nums[i+1] = nums[i+1], nums[i]
                }
            } 
            sense = invert(sense)
        } else {
            if i+1 < len(nums) {
                if nums[i] < nums[i+1] {
                    //swap
                    nums[i], nums[i+1] = nums[i+1], nums[i]
                }
            } 
            sense = invert(sense)
        }
    }
}
func invert(sense string) string {
    if sense == "smaller" {
        return "bigger"
    }
    return "smaller"
}
