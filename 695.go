package main

import (
	"fmt"
)

func maxAreaOfIsland(grid [][]int) int {
	var ans int = 0
	var m = make([][]int, len(grid))
	for i := range m {
		m[i] = make([]int, len(grid[i]))
	}
	for i := range grid {
		for j, value := range grid[i] {
			if value != 0 {
				p := calculateMaxArea(grid, i, j, 0, m)
				if p > ans {
					ans = p
				}
			}
		}
	}
	return ans
}

func calculateMaxArea(grid [][]int, row int, column int, result int, m [][]int) int {
	if grid[row][column] == 0 {
		return result
	} else if m[row][column] == 1 {
		return result
	}
	m[row][column] = 1
	result++
	if row < len(grid)-1 {
		result = calculateMaxArea(grid, row+1, column, result, m)
	}
	if row > 0 {
		result = calculateMaxArea(grid, row-1, column, result, m)
	}
	if column < len(grid[row])-1 {
		result = calculateMaxArea(grid, row, column+1, result, m)
	}
	if column > 0 {
		result = calculateMaxArea(grid, row, column-1, result, m)
	}
	return result
}

func main() {
	data := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Printf("value: %d", maxAreaOfIsland(data))
}
