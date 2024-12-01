package day1_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day1"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	data := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	res, err := day1.Part1(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 11, res)
}

func TestPart2(t *testing.T) {
	data := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	res, err := day1.Part2(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 31, res)
}
