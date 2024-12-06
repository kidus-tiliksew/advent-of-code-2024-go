package day6

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

	m, n := len(matrix), len(matrix[0])
	seen := make(map[*rune]bool)

	startingRow, startingCol, startingDir := 0, 0, []int{}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			curr := matrix[row][col]
			if curr == '^' || curr == '<' || curr == '>' || curr == 'v' {
				startingRow, startingCol = row, col

				switch curr {
				case '^':
					startingDir = []int{-1, 0}
				case '<':
					startingDir = []int{0, -1}
				case '>':
					startingDir = []int{0, 1}
				case 'v':
					startingDir = []int{1, 0}
				}
			}
		}
	}

	isValid := func(row, col int) bool {
		return 0 <= row && row < m && 0 <= col && col < n
	}

	ans := 0
	var dfs func(int, int, []int)
	dfs = func(row, col int, dir []int) {
		if !seen[&matrix[row][col]] {
			seen[&matrix[row][col]] = true
			ans++
		}

		nextRow, nextCol := row+dir[0], col+dir[1]
		if !isValid(nextRow, nextCol) {
			return
		}

		for matrix[nextRow][nextCol] == '#' {
			if dir[0] == -1 && dir[1] == 0 {
				// Move right
				dir = []int{0, 1}
			} else if dir[0] == 0 && dir[1] == -1 {
				// Move up
				dir = []int{-1, 0}
			} else if dir[0] == 0 && dir[1] == 1 {
				// Move down
				dir = []int{1, 0}
			} else if dir[0] == 1 && dir[1] == 0 {
				// Move left
				dir = []int{0, -1}
			}

			nextRow, nextCol = row+dir[0], col+dir[1]
			if !isValid(nextRow, nextCol) {
				return
			}
		}

		dfs(nextRow, nextCol, dir)
	}

	dfs(startingRow, startingCol, startingDir)

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

	m, n := len(matrix), len(matrix[0])

	startingRow, startingCol, startingDir := 0, 0, []int{}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			curr := matrix[row][col]
			if curr == '^' || curr == '<' || curr == '>' || curr == 'v' {
				startingRow, startingCol = row, col

				switch curr {
				case '^':
					startingDir = []int{-1, 0}
				case '<':
					startingDir = []int{0, -1}
				case '>':
					startingDir = []int{0, 1}
				case 'v':
					startingDir = []int{1, 0}
				}
			}
		}
	}

	isValid := func(row, col int) bool {
		return 0 <= row && row < m && 0 <= col && col < n
	}

	ans := 0
	var dfs func(int, int, []int, map[*rune]int)
	dfs = func(row, col int, dir []int, seen map[*rune]int) {
		seen[&matrix[row][col]]++

		if seen[&matrix[row][col]] > 100 {
			ans++
			return
		}

		nextRow, nextCol := row+dir[0], col+dir[1]
		if !isValid(nextRow, nextCol) {
			return
		}

		for matrix[nextRow][nextCol] == '#' || matrix[nextRow][nextCol] == 'O' {
			if dir[0] == -1 && dir[1] == 0 {
				// Move right
				dir = []int{0, 1}
			} else if dir[0] == 0 && dir[1] == -1 {
				// Move up
				dir = []int{-1, 0}
			} else if dir[0] == 0 && dir[1] == 1 {
				// Move down
				dir = []int{1, 0}
			} else if dir[0] == 1 && dir[1] == 0 {
				// Move left
				dir = []int{0, -1}
			}

			nextRow, nextCol = row+dir[0], col+dir[1]
			if !isValid(nextRow, nextCol) {
				return
			}
		}

		dfs(nextRow, nextCol, dir, seen)
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == startingRow && col == startingCol {
				continue
			}

			seen := make(map[*rune]int)
			curr := matrix[row][col]
			matrix[row][col] = '#'
			dfs(startingRow, startingCol, startingDir, seen)
			matrix[row][col] = curr
		}
	}

	return ans, nil
}
