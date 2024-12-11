package day11

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Part1(input io.Reader, blinks int) (int, error) {
	scanner := bufio.NewScanner(input)

	stones := []string{}
	for scanner.Scan() {
		stones = strings.Split(scanner.Text(), " ")
	}

	var score func(string, int) int
	score = func(stone string, blinks int) int {
		if blinks == 0 {
			return 1
		}

		if stone == "0" {
			return score("1", blinks-1)
		}

		if len(stone)%2 == 0 {
			s := []rune(stone)
			middle := len(stone) / 2

			left, right := s[:middle], s[middle:]

			return score(trimLeadingZeros(left), blinks-1) + score(trimLeadingZeros(right), blinks-1)
		}

		i, _ := strconv.Atoi(stone)
		i *= 2024

		return score(strconv.Itoa(i), blinks-1)
	}

	ans := 0
	for _, e := range stones {
		ans += score(e, blinks)
	}

	return ans, nil
}

func Part2(input io.Reader, blinks int) (int, error) {
	scanner := bufio.NewScanner(input)

	stones := []string{}
	for scanner.Scan() {
		stones = strings.Split(scanner.Text(), " ")
	}

	type memoStruct struct {
		stone  string
		blinks int
	}

	memo := make(map[memoStruct]int)

	var score func(string, int) int
	score = func(stone string, blinks int) int {
		if blinks == 0 {
			return 1
		}

		if _, ok := memo[memoStruct{stone, blinks}]; ok {
			return memo[memoStruct{stone, blinks}]
		}

		if stone == "0" {
			memo[memoStruct{stone, blinks}] = score("1", blinks-1)
			return memo[memoStruct{stone, blinks}]
		}

		if len(stone)%2 == 0 {
			s := []rune(stone)
			middle := len(stone) / 2

			left, right := s[:middle], s[middle:]
			memo[memoStruct{stone, blinks}] = score(trimLeadingZeros(left), blinks-1) + score(trimLeadingZeros(right), blinks-1)
			return memo[memoStruct{stone, blinks}]
		}

		i, _ := strconv.Atoi(stone)
		i *= 2024

		memo[memoStruct{stone, blinks}] = score(strconv.Itoa(i), blinks-1)
		return memo[memoStruct{stone, blinks}]
	}

	ans := 0
	for _, e := range stones {
		ans += score(e, blinks)
	}

	return ans, nil
}

func trimLeadingZeros(b []rune) string {
	str := strings.TrimLeft(string(b), "0")
	if str == "" {
		return "0"
	}

	return str
}
