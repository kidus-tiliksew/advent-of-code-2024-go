package day8

import (
	"bufio"
	"io"
)

func Part1(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	matrix := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		for _, e := range scanner.Text() {
			row = append(row, e)
		}

		matrix = append(matrix, row)
	}

	locations := make(map[rune][][]int)
	seen := make(map[*rune]bool)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '.' {
				continue
			}

			val := matrix[row][col]
			if _, ok := locations[val]; !ok {
				locations[val] = [][]int{}
			}

			locations[val] = append(locations[val], []int{row, col})
		}
	}

	isValid := func(row, col int) bool {
		return 0 <= row && row < len(matrix) && 0 <= col && col < len(matrix[0])
	}

	ans := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '.' {
				continue
			}

			for _, cell := range locations[matrix[row][col]] {
				currRow, currCol := cell[0], cell[1]

				if currRow == row && currCol == col {
					continue
				}

				x, y := currRow-row, currCol-col
				if isValid(currRow+x, currCol+y) && !seen[&matrix[currRow+x][currCol+y]] {
					seen[&matrix[currRow+x][currCol+y]] = true
					ans++
				}
			}
		}
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	matrix := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		for _, e := range scanner.Text() {
			row = append(row, e)
		}

		matrix = append(matrix, row)
	}

	locations := make(map[rune][][]int)
	seen := make(map[*rune]bool)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '.' {
				continue
			}

			val := matrix[row][col]
			if _, ok := locations[val]; !ok {
				locations[val] = [][]int{}
			}

			locations[val] = append(locations[val], []int{row, col})
		}
	}

	isValid := func(row, col int) bool {
		return 0 <= row && row < len(matrix) && 0 <= col && col < len(matrix[0])
	}

	ans := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '.' {
				continue
			}

			for _, cell := range locations[matrix[row][col]] {
				currRow, currCol := cell[0], cell[1]

				if currRow == row && currCol == col {
					continue
				}

				if isValid(currRow, currCol) && !seen[&matrix[currRow][currCol]] {
					seen[&matrix[currRow][currCol]] = true
					ans++
				}

				x, y := currRow-row, currCol-col
				currRow += x
				currCol += y

				for isValid(currRow, currCol) {
					if !seen[&matrix[currRow][currCol]] {
						seen[&matrix[currRow][currCol]] = true
						ans++
					}

					currRow += x
					currCol += y
				}
			}
		}
	}

	return ans, nil
}
