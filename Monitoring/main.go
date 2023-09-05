package main

func RotateMatrix(rows, columns int, matrix [][]int) [][]int {
	resultingMatrix := make([][]int, columns)

	for i := 0; i < columns; i++ {
		resultingMatrix[i] = make([]int, rows)
	}

	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			resultingMatrix[i][j] = matrix[j][i]
		}
	}
	return resultingMatrix
}
