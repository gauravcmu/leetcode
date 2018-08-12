# Problem Description

Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.

## Solution approach:
- Sort the input array by the start field 
- Insert first element to result array
- iterate over input array, merge each element into result array. 

## Complexity
- O(N) worst case space complexity. 
- Since we sort the array - O(nLogn) and insert is O(n*n)
- To make O(nlogn) - compare element with only the last element inserted into result.  
