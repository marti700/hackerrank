package main

import (
	"fmt"
)

// finds the absolute value of a number
func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}

// converts  the columns of a 3 by 3 matrix to rows, the rightmost colum becomes the first row
// and the leftmost the last
// Ej:
// The matrix
// 6 1 2
// 7 2 6
// 5 6 2

// will become
// 2 6 2
// 1 2 6
// 6 7 5

func rotateCols(m [3][3]int) [3][3]int {
	var matrix [3][3]int
	var tempMatrix [3]int

	row := 0
	col := 2

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if col == 2 {
				tempMatrix[j] = m[row][col]
				row++
			} else if col == 1 {
				tempMatrix[j] = m[row][col]
				row++
			} else {
				tempMatrix[j] = m[row][col]
				row++
			}
		}
		matrix[i] = tempMatrix
		tempMatrix = clear(tempMatrix)
		row = 0
		col--
	}

	return matrix
}

// Swaps the matrix first and last rows
func swapRows(m [3][3]int) [3][3]int {
	m[0], m[2] = m[2], m[0]
	return m
}

// resets a 3 element int array by setting its values to zero
func clear(arr [3]int) [3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

// there are only 8 magicquare that can be form with a 3X3 matrix
// this function starts with one 3X3 magic quare and rotates its
// columns and swap its rows until the remaining 7 magic squares are obtained
func generateMagicSquares() [8][3][3]int {

	magicSquares := [8][3][3]int{
		{
			{2, 7, 6},
			{9, 5, 1},
			{4, 3, 8},
		},
	}

	// get all the magic squares that can be obtained by rotating colums
	for k := 1; k < 4; k++ {
		magicSquares[k] = rotateCols(magicSquares[k-1])
	}

	// get the remaining magic squares by swapping the first and last row of the already created (by rotating rows) magic squares
	for k := 4; k < 8; k++ {
		magicSquares[k] = swapRows(magicSquares[k-4])
	}

	return magicSquares
}

// calculates the cost of the magic square by by comparing the input with all 8 3 by 3 magic scquare and selecting
// the one wiht less cost
func formingMagicSquare(s [3][3]int) int {
	magiSquares := generateMagicSquares()
	var temp int
	cost := 200000

	for i := 0; i < 8; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				temp += abs(magiSquares[i][j][k] - s[j][k])
			}
		}
		if cost > temp {
			cost = temp
		}
		temp = 0
	}

	return cost
}

func main() {

	var input [3][3]int

	// takes input
	fmt.Scanf("%d %d %d\n %d %d %d\n %d %d %d\n",
		&input[0][0], &input[0][1], &input[0][2],
		&input[1][0], &input[1][1], &input[1][2],
		&input[2][0], &input[2][1], &input[2][2])

	fmt.Println(formingMagicSquare(input))
}
