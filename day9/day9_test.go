package day9_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day9"
	"github.com/stretchr/testify/assert"
)

var data string = "2333133121414131402"

func TestPart1(t *testing.T) {
	res, err := day9.Part1(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1928, res)
}

func TestPart2(t *testing.T) {
	res, err := day9.Part2(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2858, res)
}
