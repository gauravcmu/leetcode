type Solution struct {
    nums []int  
    shuf []int 
}


func Constructor(nums []int) Solution {
    n := make([]int, len(nums))
    copy (n, nums)
    shuf := make([]int, len(nums))
    copy (shuf, nums)
    
    return Solution{
        nums: n,   
        shuf: shuf, 
    } 
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
   return this.nums 
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    var r int 
    for i:=1; i<len(this.shuf);i++ {
        r  = rand.Intn(i)
        this.shuf[r], this.shuf[i] = this.shuf[i], this.shuf[r]
    } 
    return this.shuf
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
