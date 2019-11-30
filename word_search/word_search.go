package main

import (
	"fmt"
)

func main() {
/*
  Consider we search for ABCCED - this will fail in this current implementation as there are two choices for E after C. 
  Algorithm takes to the E on the first row 
*/
	board := [][]byte{{'A', 'B', 'C', 'E'},
			  {'S', 'F', 'C', 'S'},
			  {'A', 'D', 'E', 'E'},
	}
	fmt.Printf("%v\n", exist(board, "ABCCED"))
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, 0)

	for i := 0; i < len(board); i++ {
		row := make([]bool, len(board[0]))
		visited = append(visited, row)
	}

	fmt.Printf("visited: %v\n", visited)

	q := &queue{}

	q.enqueue(&info{row: 0, col: 0})
	i := 0 //i is the index within word that is to be matched next.

	for q.isEmpty() != true && i < len(word) {
		u := q.dequeue()

		if word[i] != board[u.row][u.col] {
			continue 
		}
		i++
		visited[u.row][u.col] = true
		fmt.Printf("visiting: %c\n", board[u.row][u.col])
		if u.row < len(board)-1 && visited[u.row+1][u.col] != true {
			q.enqueue(&info{row:u.row+1, col:u.col,})
		}
		if u.row > 0 && visited[u.row-1][u.col] != true {
			q.enqueue(&info{row:u.row-1, col:u.col,})
		}
		if u.col < len(board[0])-1 && visited[u.row][u.col+1] != true {
			q.enqueue(&info{row:u.row, col:u.col+1,})
		}
		if u.col > 0 && visited[u.row][u.col-1] != true {
			q.enqueue(&info{row:u.row, col:u.col-1,})
		}
	}

	return i == len(word)
}

type info struct {
	row int
	col int
}

type queue []*info

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) enqueue(v *info) {
	fmt.Printf("enqueue: %v\n", v)
	*q = append(*q, v)
}

func (q *queue) dequeue() *info {
	t := (*q)[0]
	(*q) = (*q)[1:]
	return t
}
