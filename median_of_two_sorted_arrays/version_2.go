/*

Merge the two arrays until mid. 
Keep previous value for even number of elements case - as we need to find average of prev and current element. 

nums1 = [1, 3]
            ^        
nums2 = [2]
         ^
     prev = 1 
     curr = 2 
l1 = 2, l2 = 1 odd = true mid = 1 

k = 2 


nums1 [1, 2]
          ^ 
nums2 [3, 4]
       ^ 
       
       mid = 2 
       prev = 2 
       curr = 3 
       k = 3
       
*/


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    k := 0
    i := 0 
    j := 0 
    curr := 0 
    prev := 0 
    l1 := len(nums1)
    l2 := len(nums2)
    
    odd := (l1 + l2 )%2 != 0 
    
    mid := (l1 + l2) / 2 
    
    
    for i, j =0, 0; i< len(nums1) && j < len(nums2) && k <= mid; {
        prev = curr 
        if nums1[i] < nums2[j] {
            curr = nums1[i]
            i++
        } else {
            curr = nums2[j]
            j++
        }
        k++
    } 
    
    for i < len(nums1) && k <= mid {
            prev = curr 
            curr = nums1[i]
            i++
            k++ 
        
    }
    for j < len(nums2) && k <= mid {
            prev = curr 
            curr = nums2[j]
            j++
            k++         
    }
    if odd == true {
        return float64(curr) 
    } else {
        return float64(prev + curr)/ 2
    }
}


/*
nums1 [1, 2]
       ^ 
nums2 [3, 4]
       ^ 

find the lenghts of the two. 
if even. then iterate until end - start + 1 / 2 merges. 
else 
iterate until start + 1 / 2 merges and then average out previous and this element. 

for 

median for odd number of elements is the element at 
0-4 (5 elements) - (4-0)+1/2 = index 2 
6 elements 0-5 - 5-0+1/2 = 3 
0, 1, 2, 3, 4, 5 
index 2 and 3 
      

*/
