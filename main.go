package main

import (
	"log"
	"os"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day1"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day2"
)

func main() {
	// Day 1
	{
		// Part 1
		input, err := os.Open("./day1/input.txt")
		handleError(err, ErrorContext{"1", "1", "could not open input"})

		res, err := day1.Part1(input)
		handleError(err, ErrorContext{"1", "1", ""})
		log.Printf("day 1: part 1: %d", res)

		// Part 2
		input, err = os.Open("./day1/input.txt")
		handleError(err, ErrorContext{"1", "2", "could not open input"})

		res, err = day1.Part2(input)
		handleError(err, ErrorContext{"1", "2", ""})
		log.Printf("day 1: part 2: %d", res)
	}

	// Day 2
	{
		// Part 1
		input, err := os.Open("./day2/input.txt")
		handleError(err, ErrorContext{"2", "1", "could not open input"})

		res, err := day2.Part1(input)
		handleError(err, ErrorContext{"2", "1", ""})
		log.Printf("day 2: part 1: %d", res)

		// Part 2
		input, err = os.Open("./day2/input.txt")
		handleError(err, ErrorContext{"2", "2", "could not open input"})

		res, err = day2.Part2(input)
		handleError(err, ErrorContext{"2", "2", ""})
		log.Printf("day 2: part 2: %d", res)
	}
}
