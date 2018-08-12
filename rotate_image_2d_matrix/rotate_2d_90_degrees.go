func rotate(matrix [][]int)  {
	//first swap elements along diagonal    
	//You only traverse one half about the diagonal - else it makes no change. 
	// i goes 0 to one less than last row
	// j goes i+1 (imagine triangle shape about diagonal) to last column
	for i:=0; i< len(matrix) -1; i++ {
	    for j:=i+1; j< len(matrix[0]); j++{
		if i != j {
	 //           fmt.Printf("swapping: %v %v a:%v b:%v\n", i,j, matrix[i][j], matrix[j][i])
		    matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	    }
	}
	
	//Now reverse each row to get the answer
	//fmt.Printf("matrix %v \n", matrix)
	for _, row := range matrix {
	    for i, j:=0, len(row)-1 ; i < j;i, j = i+1, j-1 {
		    row[i], row[j] = row[j], row[i]
	    }
	}  
	//fmt.Printf("matrix %v \n", matrix)
}
