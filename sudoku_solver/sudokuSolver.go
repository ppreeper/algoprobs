package main

// https://iq.opengenus.org/backtracking-sudoku/

import "fmt"

// printGrid function to print grid
func printGrid(grid *[][]string) {
	fmt.Println("___________________")

	for _, x := range *grid {
		fmt.Printf("|")
		for _, y := range x {
			fmt.Printf("%s|", y)
		}
		fmt.Println()
	}
	fmt.Println("-------------------")
}

// function to check if the value to be assigned to a cell already exists in that row of that cell
// it returns true if 'val' can be placed in a cell with row number as 'row'
func rowCheck(grid *[][]string, row int, num string) bool {
	// iterate through that row
	for y := 0; y < 9; y++ {
		// if value to be assigned is found then
		// it can't be placed in that cell
		if (*grid)[row][y] == num {
			return false
		}
	}
	return true
}

// function to check if the value to be assigned to a cell already exists in that column of that cell
// it returns true if 'val' can be placed in a cell with column number as 'col'
func colCheck(grid *[][]string, col int, num string) bool {
	// iterate through that column
	for x := 0; x < 9; x++ {
		// if value to be assigned is found then
		// it can't be placed in that cell
		if (*grid)[x][col] == num {
			return false
		}
	}
	return true
}

// function to check "box" condition
func boxCheck(grid *[][]string, row, col int, num string) bool {
	// get the center cell(r,c) of the box
	// with simple formula below
	r := (row/3)*3 + 1
	c := (col/3)*3 + 1
	// iterate through each cell of the box
	boxRange := []int{-1, 0, 1}
	for _, i := range boxRange {
		for _, j := range boxRange {
			// for each cell of the box check if the value to be placed exists
			// if exits then it can't be placed in that cell
			if (*grid)[r+i][c+j] == num {
				return false
			}
		}
	}
	return true
}

// function to find unassigned cell(a cell which doesn't contain a value) in the grid
func findUnassigned(grid *[][]string) (int, int) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// cell which contains 0 is unassigned
			if (*grid)[x][y] == "0" {
				return x, y
			}
		}
	}
	// if no cell left unassigned
	return -1, -1
}

// function to complete the sudoku
func sudokuSolver(grid *[][]string) bool {
	// find an unassigned cell in the grid
	x, y := findUnassigned(grid)
	// if no cell remain unassigned then the grid is filled completely and is valid
	if x == -1 && y == -1 {
		return true
	}

	// for each 'num' ranging from 1 to 9 check if it can be placed at '(i,j)'
	for _, num := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		// 'num' can be placed at '(i,j)' if
		// any value in row 'i' is not equal to 'num'
		// any value in column 'j' is not equal to 'num'
		// any value in the box it belongs to is not equal to 'num'
		if rowCheck(grid, x, num) && colCheck(grid, y, num) && boxCheck(grid, x, y, num) {
			// all the conditions are satisfied

			// place 'num' at '(i,j)' temporarily
			(*grid)[x][y] = num

			// check for the next cells, recursively
			if sudokuSolver(grid) {
				return true
			}

			// we've reached here because the choice we made by putting some 'num' here was wrong
			// hence now leave the cell unassigned to try another possibilities
			(*grid)[x][y] = "0"
		}
	}
	// putting any value doesn't solve the grid that means we've made a wrong choice earlier
	// if no value can be placed at '(i,j)' then returns false meaning that the current
	// sudoku is not feasible and try for another possibilities
	return false
}

func main() {
	var grid = [][]string{
		{"5", "0", "0", "9", "0", "3", "0", "0", "4"},
		{"9", "0", "0", "7", "0", "0", "0", "0", "0"},
		{"6", "0", "3", "0", "0", "1", "0", "9", "0"},

		{"0", "2", "0", "0", "0", "0", "5", "0", "0"},
		{"0", "9", "0", "4", "0", "8", "0", "3", "0"},
		{"0", "0", "7", "0", "0", "0", "0", "8", "0"},

		{"0", "5", "0", "6", "0", "0", "3", "0", "1"},
		{"0", "0", "0", "0", "0", "5", "0", "0", "7"},
		{"7", "0", "0", "2", "0", "9", "0", "0", "8"},
	}
	printGrid(&grid)
	fmt.Println()

	if ok := sudokuSolver(&grid); !ok {
		fmt.Println("No solution exists!")
		return
	}
	printGrid(&grid)
}
