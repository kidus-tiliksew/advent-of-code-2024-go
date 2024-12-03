package day2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Part1(input io.Reader) (int, error) {
	levels, err := getLevels(input)
	if err != nil {
		return 0, err
	}

	ans := 0
	for _, level := range levels {
		if isSafe(true, level) || isSafe(false, level) {
			ans++
		}
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	levels, err := getLevels(input)
	if err != nil {
		return 0, err
	}

	ans := 0
	for _, level := range levels {
		safe := false

		for i := range level {
			data := append([]int{}, level[:i]...)
			data = append(data, level[i+1:]...)

			if isSafe(true, data) || isSafe(false, data) {
				safe = true
			}
		}

		if safe {
			ans++
		}
	}

	return ans, nil
}

func isSafe(increasing bool, level []int) bool {
	isValid := func(x, y int) bool {
		if increasing && x < y {
			return false
		}

		if !increasing && x > y {
			return false
		}

		diff := abs(x - y)
		if diff < 1 || diff > 3 {
			return false
		}

		return true
	}

	for i := 1; i < len(level); i++ {
		if !isValid(level[i], level[i-1]) {
			return false
		}
	}

	return true
}

func getLevels(input io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(input)

	levels := [][]int{}

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Fields(text)

		level := []int{}
		for _, e := range s {
			i, err := strconv.Atoi(e)
			if err != nil {
				return nil, fmt.Errorf("could not convert level to integers")
			}

			level = append(level, i)
		}

		levels = append(levels, level)
	}

	return levels, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
