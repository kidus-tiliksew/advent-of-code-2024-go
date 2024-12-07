package day7

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type equation struct {
	value   int64
	numbers []int64
}

func Part1(input io.Reader) (int64, error) {
	scanner := bufio.NewScanner(input)

	equations := []equation{}
	for scanner.Scan() {
		e := strings.Split(scanner.Text(), ":")

		value, err := strconv.ParseInt(e[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("could not parse equation %s: %w", e, err)
		}

		numbers := []int64{}
		for _, n := range strings.Fields(e[1]) {
			i, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("could not parse equation: %s: %w", n, err)
			}

			numbers = append(numbers, i)
		}

		equations = append(equations, equation{value, numbers})
	}

	validIndex := make(map[int]bool)

	var backtrack func([]int64, int, int64, int64, int)
	backtrack = func(numbers []int64, currIndex int, currVal int64, target int64, i int) {
		if currIndex == len(numbers) {
			if currVal == target {
				validIndex[i] = true
			}
			return
		}

		for _, op := range []rune{'+', '*'} {
			nextVal := numbers[currIndex]
			var newVal int64
			if op == '+' {
				newVal = currVal + nextVal
			} else {
				newVal = currVal * nextVal
			}
			backtrack(numbers, currIndex+1, newVal, target, i)
		}
	}

	for i, e := range equations {
		backtrack(e.numbers, 1, e.numbers[0], e.value, i)
	}

	ans := int64(0)
	for i := range validIndex {
		ans += equations[i].value
	}

	return ans, nil
}

func Part2(input io.Reader) (int64, error) {
	scanner := bufio.NewScanner(input)

	equations := []equation{}
	for scanner.Scan() {
		e := strings.Split(scanner.Text(), ":")

		value, err := strconv.ParseInt(e[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("could not parse equation %s: %w", e, err)
		}

		numbers := []int64{}
		for _, n := range strings.Fields(e[1]) {
			i, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("could not parse equation: %s: %w", n, err)
			}

			numbers = append(numbers, i)
		}

		equations = append(equations, equation{value, numbers})
	}

	validIndex := make(map[int]bool)

	var backtrack func([]int64, int, int64, int64, int)
	backtrack = func(numbers []int64, currIndex int, currVal int64, target int64, i int) {
		if currIndex == len(numbers) {
			if currVal == target {
				validIndex[i] = true
			}
			return
		}

		for _, op := range []rune{'+', '*', '|'} {
			nextVal := numbers[currIndex]
			var newVal int64
			if op == '+' {
				newVal = currVal + nextVal
			} else if op == '*' {
				newVal = currVal * nextVal
			} else {
				newVal, _ = strconv.ParseInt(fmt.Sprint(currVal)+fmt.Sprint(nextVal), 10, 64)
			}
			backtrack(numbers, currIndex+1, newVal, target, i)
		}
	}

	for i, e := range equations {
		backtrack(e.numbers, 1, e.numbers[0], e.value, i)
	}

	ans := int64(0)
	for i := range validIndex {
		ans += equations[i].value
	}

	return ans, nil
}
