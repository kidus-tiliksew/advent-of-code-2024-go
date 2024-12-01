package day1

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Part1(input io.Reader) (int, error) {
	left, right, err := GetLeftRightCols(input)
	if err != nil {
		return 0, err
	}

	slices.Sort(left)
	slices.Sort(right)

	ans := 0
	for i, e := range left {
		ans += abs(e - right[i])
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	left, right, err := GetLeftRightCols(input)
	if err != nil {
		return 0, err
	}

	counter := make(map[int]int)
	for _, e := range right {
		counter[e]++
	}

	ans := 0
	for _, e := range left {
		ans += e * counter[e]
	}

	return ans, nil
}

func GetLeftRightCols(input io.Reader) ([]int, []int, error) {
	scanner := bufio.NewScanner(input)

	left, right := []int{}, []int{}
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)

		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid input format: %s", text)
		}

		i, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("could not convert left column integer for text %s: ", text)
		}
		left = append(left, i)

		i, err = strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("could not convert right column integer for text %s: ", text)
		}
		right = append(right, i)
	}

	return left, right, nil
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
