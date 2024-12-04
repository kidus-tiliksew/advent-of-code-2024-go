package day4

import (
	"bufio"
	"io"
	"slices"
)

func Part1(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	matrix := [][]byte{}
	for scanner.Scan() {
		text := scanner.Text()
		matrix = append(matrix, []byte(text))
	}

	m, n := len(matrix), len(matrix[0])

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	isValid := func(row, col int) bool {
		return 0 <= row && row < m && 0 <= col && col < n
	}

	ans := 0

	var dfs func(int, int, string, []int)
	dfs = func(row, col int, curr string, direction []int) {
		if curr == "XMAS" {
			ans++
		}

		if len(curr) >= 4 {
			return
		}

		nextRow, nextCol := row+direction[0], col+direction[1]
		if isValid(nextRow, nextCol) {
			curr += string(matrix[nextRow][nextCol])
			dfs(nextRow, nextCol, curr, direction)
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			for _, dir := range directions {
				dfs(row, col, string(matrix[row][col]), dir)
			}
		}
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	matrix := [][]byte{}
	for scanner.Scan() {
		text := scanner.Text()
		matrix = append(matrix, []byte(text))
	}

	m, n := len(matrix), len(matrix[0])

	directions := [][]int{
		{-1, -1}, // Top-left
		{-1, 1},  // Top-right
		{1, -1},  // Bottom-left
		{1, 1},   // Bottom-right
	}

	isValid := func(row, col int) bool {
		return 0 <= row && row < m && 0 <= col && col < n
	}

	ans := 0

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] != 'A' {
				continue
			}

			valid := true
			for _, dir := range directions {
				nextRow, nextCol := row+dir[0], col+dir[1]
				if !isValid(nextRow, nextCol) {
					valid = false
				}
			}

			if !valid {
				continue
			}

			topLeft := matrix[row+directions[0][0]][col+directions[0][1]]
			topRight := matrix[row+directions[1][0]][col+directions[1][1]]
			bottomLeft := matrix[row+directions[2][0]][col+directions[2][1]]
			bottomRight := matrix[row+directions[3][0]][col+directions[3][1]]

			mas := []byte("MAS")
			leftToRight := []byte{topLeft, matrix[row][col], bottomRight}
			rightToLeft := []byte{topRight, matrix[row][col], bottomLeft}

			slices.Sort(mas)
			slices.Sort(leftToRight)
			slices.Sort(rightToLeft)

			if string(leftToRight) == string(mas) && string(rightToLeft) == string(mas) {
				ans++
			}
		}
	}

	return ans, nil
}
