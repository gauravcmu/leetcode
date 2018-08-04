func islandPerimeter(grid [][]int) int {
    l := len(grid)
    w := len(grid[0])
    
    var sum int 
    for i  := 0 ; i<l; i++ {
        for j:=0; j< w; j++ {
        if grid[i][j] == 1 {
            if i-1 >= 0 {
             sum += grid[i-1][j] ^ 1     
            } else {
                sum += 1
            }
            if i+1 < l {
              sum +=  grid[i+1][j] ^ 1
            } else {
                sum += 1
            } 
            if j-1 >= 0 {
             sum += grid[i][j-1] ^ 1
            } else {
                sum += 1
            }
            
            if j+1 < w {
              sum +=  grid[i][j+1] ^ 1
            } else {
                sum += 1
            }
        }
        fmt.Printf("%v l:%v w:%v\n", sum, l, w)
        }
    } 
    return sum 
}
