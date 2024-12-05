package day5

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Part1(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	rules := make(map[int][]int)
	updates := [][]int{}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		s := strings.Split(text, "|")
		x, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, fmt.Errorf("failed to parse rule %s: %w", text, err)
		}
		y, err := strconv.Atoi(s[1])
		if err != nil {
			return 0, fmt.Errorf("failed to parse rule %s: %w", text, err)
		}

		if _, ok := rules[x]; !ok {
			rules[x] = []int{}
		}

		rules[x] = append(rules[x], y)
	}

	for scanner.Scan() {
		text := scanner.Text()

		m := []int{}
		for _, s := range strings.Split(text, ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return 0, fmt.Errorf("failed to parse update %s: %w", s, err)
			}
			m = append(m, i)
		}

		updates = append(updates, m)
	}

	ans := 0
	for _, e := range updates {
		isValid := true

	Exit:
		for i := 1; i < len(e); i++ {
			for j := 0; j < i; j++ {
				if slices.Contains(rules[e[i]], e[j]) {
					isValid = false
					break Exit
				}
			}
		}

		if isValid {
			index := (len(e) - 1) / 2
			ans += e[index]
		}
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	rules := make(map[int][]int)
	updates := [][]int{}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		s := strings.Split(text, "|")
		x, err := strconv.Atoi(s[0])
		if err != nil {
			return 0, fmt.Errorf("failed to parse rule %s: %w", text, err)
		}
		y, err := strconv.Atoi(s[1])
		if err != nil {
			return 0, fmt.Errorf("failed to parse rule %s: %w", text, err)
		}

		if _, ok := rules[x]; !ok {
			rules[x] = []int{}
		}

		rules[x] = append(rules[x], y)
	}

	for scanner.Scan() {
		text := scanner.Text()

		m := []int{}
		for _, s := range strings.Split(text, ",") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return 0, fmt.Errorf("failed to parse update %s: %w", s, err)
			}
			m = append(m, i)
		}

		updates = append(updates, m)
	}

	ans := 0
	for k, e := range updates {
		isValid := true

		for i := 1; i < len(e); i++ {
			for j := 0; j < i; j++ {
				if slices.Contains(rules[e[i]], e[j]) {
					updates[k][i], updates[k][j] = updates[k][j], updates[k][i]
					isValid = false
				}
			}
		}

		if !isValid {
			index := (len(e) - 1) / 2
			ans += updates[k][index]
		}
	}

	return ans, nil
}
