/*
 use two pointers
	- Z goes to find zero element. 
	- NZ start from Z index and find non-zero element. 
	- swap
	- repeat
*/
func moveZeroes(nums []int)  {
    z := 0
    nz := 0

    for {
        for z < len(nums) -1 {
            if nums[z] == 0 {
               break 
            }
            z++
        }
        if z >= len(nums) -1  {
            return
        }
        
        nz = z 
        for nz < len(nums) {
            if nums[nz] != 0 {
                break
            } 
            nz++
        }
        
        if nz >= len(nums) {
            return
        }
        
        fmt.Printf("z:%v nz:%v\n", z, nz)
        nums[z], nums[nz] = nums[nz], nums[z]
        z++
        nz++
    } 
}
