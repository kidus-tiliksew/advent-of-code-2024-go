package day11_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day11"
	"github.com/stretchr/testify/assert"
)

var data string = "125 17"

func TestPart1(t *testing.T) {
	res, err := day11.Part1(strings.NewReader(data), 25)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 55312, res)
}

func TestPart2(t *testing.T) {
	res, err := day11.Part2(strings.NewReader(data), 25)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 55312, res)
}
