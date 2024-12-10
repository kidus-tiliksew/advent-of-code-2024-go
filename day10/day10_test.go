package day10_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day10"
	"github.com/stretchr/testify/assert"
)

var data string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1(t *testing.T) {
	res, err := day10.Part1(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 36, res)
}

func TestPart2(t *testing.T) {
	res, err := day10.Part2(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 81, res)
}
