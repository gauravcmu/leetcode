
/* 
Input: [0,1,0,3,12]

Iterate over the array and count zeros_so_far. 
    - for each element - if non-zero. 
        - copy that element to index i - zeroes_so_far position. 
        if element is zero - then update zeroes_so_far and keep going. 
          4
[0,1,0,3,12] 
   ^ 
[1, 1, 0, 3, 12] 
zsf = 1

[0,1,0,3,12] 
       ^ 
[1, 3, 0, 3, 12] 
zsf = 2

[0,1,0,3,12] 
          ^ 
[1, 3, 12, 3, 12] 
zsf = 2

Once first pass is done. Do second pass from reverse, copying zsf number of zeroes. 

*/
func moveZeroes(nums []int)  {
    zeroes_so_far := 0 
    
    for index, value := range nums {
        if value == 0 {
            zeroes_so_far++
        } else {
            nums[index - zeroes_so_far] = value
        }
    }
    
    for ; zeroes_so_far > 0; zeroes_so_far-- {
        nums[len(nums)-zeroes_so_far] = 0 
    }
}

