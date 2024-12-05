package day5_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day5"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	data := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	res, err := day5.Part1(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 143, res)
}

func TestPart2(t *testing.T) {
	data := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	res, err := day5.Part2(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 123, res)
}