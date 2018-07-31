type TicTacToe struct {
   T [][]int 
   Capacity int 
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
    t := make([][]int, n)
    for k, _ := range t {
        t[k] = make([]int, n)
    } 
    
    return TicTacToe {
        T: t, 
        Capacity: n*n,
    } 
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
    if this.Capacity ==  0 {
       return 0 
    }  
    
    this.T[row][col] = player
    this.Capacity-- 
    
    n := len(this.T[0])
    if n*n - this.Capacity < n {
        return 0
    }
    if row == col {
        //diagonal element being added
        return this.IsWinner(row, col, player, true, true) 
    } else if row + col == n-1 {
        //on antidiag     
        return this.IsWinner(row, col, player, false, true) 
    } else {
        return this.IsWinner(row, col, player, false, false) 
    }
    
    return this.IsWinner(row, col, player, false, false)
}

func (this *TicTacToe)IsWinner(row, col, player int, diag, antidiag bool) int {
    //given the row / col, we know the possible winning moves could be 
    //entire row
    //entire column
    //diagonal going through the point
    var found bool 
    n := len(this.T[0])
    
    if this.T[row][0] == player && player == this.T[row][n-1] {
        for c:=0; c<=n-1;c++ {
            if this.T[row][c] != player {
               found = false  
               break 
            } else {
               found = true 
            }
        }
    
        if found {
            return player
        } 
        found = false
    }
    
    if this.T[0][col] == player && player == this.T[n-1][col] {
        for r:=0; r<=n-1;r++ {
            if this.T[r][col] != player {
               found = false  
               break 
            } else {
               found = true 
            }
        }
    
        if found {
            return player
        } 
        found = false
    } 
    
    
    if this.T[0][0] == player && player ==this.T[n-1][n-1] {
        if diag {
            //move on diagonal
            for r:=0; r<= n-1;r++ {
                if this.T[r][r] != player {
                   found = false  
                   break 
                } else {
                   found = true 
                }
            }
            if found {
                return player
            } 
            found = false
        }
    }

    if this.T[n-1][0] == player && player ==this.T[0][n-1] {
        if antidiag {
            for r, c:=n-1, 0; r>= 0 && c < n;r, c = r-1, c+1  {
                //r = 1, c = 0 
                // r = 0, c = 1 
                if this.T[r][c] != player {
                   found = false  
                   break 
                } else {
                   found = true 
                }
            }
        }
    }
    
    if found {
        return player
    } else {
        return 0
    } 
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
