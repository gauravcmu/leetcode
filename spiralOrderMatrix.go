func spiralOrder(matrix [][]int) []int {
    lr := len(matrix)
    if lr== 0 {
        return []int {}
    }
    lc := len(matrix[0])
    rowstart := 0 
    rowend := lr - 1 //index of last row
    
    colstart := 0 
    colend := lc - 1 // index of last column
    
    direction := 0 
   
    res := make([]int, 0)
    var i int 
    var j int 
    var count int 
    for  {
        if direction == 0 {
           //left to right  
           //colstart to colend j  
            for j=colstart; j<= colend && count < lr*lc; j++ {
                res = append(res, matrix[rowstart][j]) 
                count++
            }
            //row read - remove row - update rowstart
            rowstart++
            //update direction 
            direction = 1
        } 
        if direction == 1 {
            //top to bottom
            //rowstart to rowend i
            for i=rowstart; i<= rowend && count < lr*lc; i++ {
                res = append(res, matrix[i][colend]) 
                count++
            }
            //column read - remove column - colend
            colend--
            //update direction 
            direction = 2
        }
        if direction == 2 {
           //right to left 
           //colend to colstart
            for j=colend; j>= colstart && count < lr*lc; j-- {
                res = append(res, matrix[rowend][j]) 
                count++
            }
            
            //row read - rowend --
            rowend--
            //update direction 
            direction = 3
        }
        if direction == 3 {
           //bottom to top  
           //rowend to rostart 
            for i=rowend; (i>= rowstart) && (count < lr*lc); i-- {
                res = append(res, matrix[i][colstart]) 
                count++
            }
            //column read - colstart--
            colstart++
            //update direction 
            direction = 0
        }
        if !(count < lr*lc) {
            break
        }
        fmt.Printf("count:%v\n", count)
    }
    return res
}
