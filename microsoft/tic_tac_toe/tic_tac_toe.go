package tic_tac_toe

type TicTacToe struct {
	grid [][]int
	n int
}

/*
	Diagonals - 	row = column
	Antidiagonals - row+column = n-1

	^ based on the above, you can build loops to check.
*/


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	grid := make([][]int, n)
	for i:=0; i< n;i++ {
		grid[i] = make([]int, n)
	}
	return TicTacToe{
		grid: grid,
		n: n,
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
	count := 0
	n := this.n
	this.grid[row][col] = player
	for j:=0;j<n;j++ {
		if this.grid[row][j] == player {
			count++
		}
	}
	if count == n {
		return player
	}

	count = 0
	for i:=0; i<n;i++ {
		if this.grid[i][col] == player {
			count++
		}
	}
	if count == n {
		return player
	}

	count = 0


	for i, j :=0, 0; i<n && j<n; i,j = i+1, j+1 {
		if this.grid[i][j] == player {
			count++
		}
	}
	if count == n {
		return player
	}

	count = 0
	for i, j := 0, (n-1); i<n && j>=0; i, j = i+1, j-1 {
		if this.grid[i][j] == player {
			count++
		}
	}

	if count == n {
		return player
	}

	return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
