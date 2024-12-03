package day3_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day3"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	res, err := day3.Part1(strings.NewReader(str))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 161, res)
}

func TestPart2(t *testing.T) {
	str := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	res, err := day3.Part2(strings.NewReader(str))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 48, res)
}

func TestFindMatches(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	matches := day3.FindMatches(str)
	assert.Equal(t, 4, len(matches))
}

func TestCaptureNumbers(t *testing.T) {
	str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	matches := day3.FindMatches(str)

	assert.Equal(t, 4, len(matches))

	assert.Equal(t, "2", matches[0][1])
	assert.Equal(t, "4", matches[0][2])

	assert.Equal(t, "5", matches[1][1])
	assert.Equal(t, "5", matches[1][2])
}

func TestFindMatchesWithConditions(t *testing.T) {
	str := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	matches := day3.FindMatchesWithConditions(str)
	assert.Equal(t, 6, len(matches))
}
