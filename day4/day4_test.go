package day4_test

import (
	"strings"
	"testing"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day4"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	data := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	res, err := day4.Part1(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 18, res)
}

func TestPart2(t *testing.T) {
	data := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	res, err := day4.Part2(strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 9, res)
}