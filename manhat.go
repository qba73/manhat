package manhat

import (
	"fmt"
	"math"
)

// CalculateDistance takes a number and calculates
// Manhattan Distance between the point and the center.
func CalculateDistance(n int) (int, error) {
	size, err := calculateMatrixSize(n)
	if err != nil {
		return 0, err
	}

	// We need only the map to lookup for position coordinates
	_, m, err := buildSpiralMatrix(int(size))
	if err != nil {
		return 0, fmt.Errorf("building spiral matrix %v", err)
	}

	// Center of the matrix
	center := point{
		x: m[1][0],
		y: m[1][1],
	}

	// Position of the Point of Interest
	poi := point{
		x: m[int(n)][0],
		y: m[int(n)][1],
	}

	return manhattanDistance(center, poi), nil
}

// point represents x, y coordinates in the 2D matrix.
type point struct {
	x int
	y int
}

// manhattanDistance takes two points (center and point of interest)
// from the matrix and calculates the 'manhattan' distance between them.
func manhattanDistance(center, poi point) int {
	return int(math.Abs(float64(poi.x-center.x)) + math.Abs(float64((poi.y - center.y))))
}

// buildSpiralMatrix takes a number, builds a spiral matrix
// and a map representing number mapped to its coordinates.
func buildSpiralMatrix(n int) ([][]int, map[int][]int, error) {
	if n <= 0 {
		return nil, nil, fmt.Errorf("invalid matrix dimention %v", n)
	}

	// coordiantes holds a map of numbers mapped to their
	// coordinates in the matrix
	coordinates := make(map[int][]int)

	// build 2D slice NxN
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// right side
	startCol := n - 1
	// left side
	endCol := 0
	// bottom side
	startRow := n - 1
	// top side
	endRow := 0

	counter := n * n

	for startCol >= endCol && startRow >= endRow {
		// bottom row
		for i := startCol; i >= endCol; i-- {
			matrix[startRow][i] = counter
			coordinates[counter] = []int{i, startRow}
			counter--
		}
		startRow--

		// right side of the matrix
		for i := startRow; i >= endRow; i-- {
			matrix[i][endCol] = counter
			coordinates[counter] = []int{endCol, i}
			counter--
		}
		endCol++

		// top row of the matrix
		for i := endCol; i <= startCol; i++ {
			matrix[endRow][i] = counter
			coordinates[counter] = []int{i, endRow}
			counter--
		}
		endRow++

		// left side of the matrix
		for i := endRow; i <= startRow; i++ {
			matrix[i][startCol] = counter
			coordinates[counter] = []int{startCol, i}
			counter--
		}
		startCol--

	}

	return matrix, coordinates, nil
}

// calculateMatrixSize takes a number and calculates
// minimal size of the matrix that will allow to take
// position x,y for the given point value.
func calculateMatrixSize(num int) (int, error) {
	if num <= 0 {
		return 0, fmt.Errorf("invalid lookup number %v, expecting number > 0", num)
	}

	x := int(math.Ceil(math.Sqrt(float64(num))))
	if x%2 == 0 {
		x++
	}
	return x, nil

}
