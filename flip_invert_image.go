/*
Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]

Solution:
	- go over row by row. 
	- in each row swap ith element with len-ith element and xor it as well. 	
*/
func flipAndInvertImage(A [][]int) [][]int {
    for _, row := range A {
        //go over row by row in A
        for i, j:=0, len(row)-1; i<=j; i,j = i+1, j-1 {
            //go over elements in a row.  
            row[i], row[j] = 1^row[j], 1^row[i]    
        }
        fmt.Printf("row:%v\n", row)
    }
    return A 
}
