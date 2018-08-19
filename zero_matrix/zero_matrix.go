//Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.
//a zero entry
type entry struct {
    row int 
    col int 
}

func setZeroes(matrix [][]int)  {
    z := make([]entry, 0) 
    
    for rn, rows := range matrix {
        for cn, colval := range rows {
            if colval == 0 {
                z = append(z, entry{row:rn, col:cn}) 
            }
        }
    }
    for _, en := range z {
        // for each entry - go over the entire row. 
        // go over the column and set it to 0 
        for j:=0; j< len(matrix[0]); j++ {
            matrix[en.row][j] = 0
        } 
        for i:=0; i<len(matrix); i++ {
            matrix[i][en.col] = 0  
        } 
    } 
}
