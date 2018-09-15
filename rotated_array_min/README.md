# problem

Find min element in a rotated array. 

# Solution approach. 
Since array is sorted, there is possiblity of using modified binary search. 

1. If start == end 
	return nums[start]
2. if mid == 0 
	if mid is less than mid+1 value (1, 2) - return value at mid. 
3. if mid > 0 
	check if mid is smaller than previous element i.e. mid-1, return value at mid. 
4. if value at mid > nums[end]
	the answer lies to the right. 
	[ 4 5 6 7 8 1 2 3 ]
            ^       ^
           mid     high 
    else 
    	answer lies to the left. 