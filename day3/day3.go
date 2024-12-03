package day3

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func Part1(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	ans := 0
	matches := FindMatches(str)
	for _, match := range matches {
		a, b := match[1], match[2]
		x, err := strconv.Atoi(a)
		if err != nil {
			return 0, fmt.Errorf("could not convert str %s: %w", a, err)
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			return 0, fmt.Errorf("could not convert str %s: %w", b, err)
		}

		ans += (x * y)
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	ans, enabled := 0, true
	matches := FindMatchesWithConditions(str)

	for _, match := range matches {
		if match[0] == "don't()" {
			enabled = false
			continue
		}

		if match[0] == "do()" {
			enabled = true
			continue
		}

		if !enabled {
			continue
		}

		a, b := match[1], match[2]
		x, err := strconv.Atoi(a)
		if err != nil {
			return 0, fmt.Errorf("could not convert str %s: %w", a, err)
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			return 0, fmt.Errorf("could not convert str %s: %w", b, err)
		}

		ans += (x * y)
	}

	return ans, nil
}

func FindMatches(str string) [][]string {
	var valid = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	return valid.FindAllStringSubmatch(str, -1)
}

func FindMatchesWithConditions(str string) [][]string {
	var valid = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))`)
	return valid.FindAllStringSubmatch(str, -1)
}
