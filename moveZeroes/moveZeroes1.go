/*
	Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
	
	Example:
	
	Input: [0,1,0,3,12]
	Output: [1,3,12,0,0]
	Note:
	
	You must do this in-place without making a copy of the array.
	Minimize the total number of operations.
*/
func moveZeroes(nums []int)  {
    for i:=len(nums)-1; i>=0;i-- {
        if nums[i] == 0 {
            for j:=i; j<len(nums)-1;j++ {
                temp := nums[j] 
                nums[j] = nums[j+1]
                nums[j+1] = temp
            }
        }
    }
}

