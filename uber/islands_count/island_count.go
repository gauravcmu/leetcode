package main

import (
"fmt"
)
func GetNumberOfIslands(binaryMatrix [][]int) int {
	// your code goes here
	visited := make([][]bool, len(binaryMatrix))
	for i, _ := range visited {
		visited[i] = make([]bool, len(binaryMatrix[0]))
	}
	count := 0
	for i:=0; i < len(binaryMatrix);i++{
		for j:=0; j<len(binaryMatrix[0]); j++ {
			if visited[i][j] != true && binaryMatrix[i][j] != 0 {
				dfs("  ", binaryMatrix, i, j, &visited)
				count++
			}
		}
	}
	return count
}

func main() {

	var bm [][]int = [][]int{
		{0,    1,    0,    1,    0},
		{0,    0,    1,    1,    1},
		{1,    0,    0,    1,    0},
		{0,    1,    1,    0,    0},
		{1,    0,    1,    0,    1},
	}
	fmt.Printf("Test1 - number of islands: %v \n", GetNumberOfIslands(bm))

	var bm2 [][]int = [][]int{
		{0,    1,    0,    1,    0},
		{0,    0,    1,    0,    1},
		{1,    0,    0,    1,    0},
		{0,    1,    1,    0,    0},
		{1,    0,    1,    0,    1},
	}
	fmt.Printf("Test2: number of islands: %v \n", GetNumberOfIslands(bm2))

	var bm3 [][]int = [][]int{}

	fmt.Printf("Test3: number of islands: %v \n", GetNumberOfIslands(bm3))
}

func dfs(path string, binaryMatrix[][]int, row, col int, visited *[][]bool) {
	//fmt.Printf("%v visiting: row:%v col:%v\n", path, row, col)
	if (*visited)[row][col] {
		return
	}

	(*visited)[row][col] = true

	if row > 0 && binaryMatrix[row-1][col] == 1{
		dfs(path + "  ",  binaryMatrix, row-1, col, visited)
	}
	if row < len(binaryMatrix)-1 && (binaryMatrix[row+1][col] == 1) {
		dfs(path + "  ", binaryMatrix, row+1, col, visited)
	}
	if col > 0 && (binaryMatrix[row][col-1] == 1) {
		dfs(path + "  ", binaryMatrix, row, col-1, visited)
	}
	if col < len(binaryMatrix[0])-1 && (binaryMatrix[row][col+1] == 1){
		dfs(path + "  ", binaryMatrix, row, col+1, visited)
	}
}
