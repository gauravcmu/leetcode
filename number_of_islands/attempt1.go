/*
 *This solution fails for a test case, needs improvement to handle that. 
 */
type point struct {
    row int 
    col int 
}


func numIslands(grid [][]byte) int {
    islands := [][][]point{}
    for i, value:= range grid {
        for j, v := range value {
            fmt.Printf("i:%v, j:%v, value:%v v:%v\n",i,j, value, v)
            //v is value at index i,j - 0 or 1. 
            if string(v) == "1" {
                fmt.Printf("Calling Add to Islands: %v\n", islands)
                addToIslands(i, j, &islands)
            } 
        }
    }
    return len(islands)
}

func addToIslands(i, j int, islands *[][][]point) {
    fmt.Printf("--- i:%v, j:%v, islands:%v\n",i,j,*islands)
    for _, island := range *islands {
        fmt.Printf("here 1")
       //going over island by island. 
        for key, land := range island {
            fmt.Printf("here 2")
            //go over indexes in island
            for _, pt := range land {
                fmt.Printf("here 3")
                if (i+1 == pt.row) && (j ==pt.col) || 
                (i-1 == pt.row) && (j == pt.col) {
                    //vertically neighbor
                    land = append(land, point{row:i, col:j})
                    island[key] = land
                    fmt.Printf("return 1 %v\n")
                    return 
                } 
                if (j+1 == pt.col) && (i == pt.row) || 
                (j-1 == pt.col) && (i == pt.row) {
                    //vertically neighbor
                    land = append(land, point{row:i, col:j})
                    island[key] = land
                    fmt.Printf("return 2 %v\n")
                    return 
                } 
            }
    }
}
                fmt.Printf("here 4")
        p := point{row:i, col:j}
        is:= make([][]point, 0) 
        is = append(is, []point{p})
        
        *islands = append(*islands, is) 
        fmt.Printf("**** adding %v\ni\n\n", *islands)
}
