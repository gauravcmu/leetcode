/*
You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

Example:

Input: [5,2,6,1]
Output: [2,1,1,0] 
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.

Solution:
NOTE - this is not accepted due to execution time for large input. 

	- go over the array.
	- count smaller than current in rest of the array 
		- sort rest of the array for less comparison
*/
func countSmaller(nums []int) []int {
   res := []int{}
    
    for key, value := range nums {
        if key == len(nums)-1 {
           //last element = add 0 for this guy 
            res = append(res, 0) 
            continue
        }
       //for each number in array, sort the rest of the array and match. 
        temp := make([]int, len(nums[key+1:]))
        copy(temp, nums[key+1:])
        res = append(res, findSmallerThanX(value, temp))
    }
    
    return res 
}
func findSmallerThanX(x int, nums[]int) int {
    count := 0 
    temp := nums
    
    sort.Slice(temp, func (a, b int)bool {
        return temp[a] < temp[b]
    })
    i:=0
    for (i <= len(nums)-1) { 
        if (temp[i] < x) {
           count++ 
            i++
        } else {
            break
        }
    }
    return count
}
