package main

import (
	"log"
	"os"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day1"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day2"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day3"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day4"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day5"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day6"
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

	// Day 3
	{
		// Part 1
		input, err := os.Open("./day3/input.txt")
		handleError(err, ErrorContext{"3", "1", "could not open input"})

		res, err := day3.Part1(input)
		handleError(err, ErrorContext{"3", "1", ""})
		log.Printf("day 3: part 1: %d", res)

		// Part 2
		input, err = os.Open("./day3/input.txt")
		handleError(err, ErrorContext{"3", "2", "could not open input"})

		res, err = day3.Part2(input)
		handleError(err, ErrorContext{"3", "2", ""})
		log.Printf("day 3: part 3: %d", res)
	}

	// Day 4
	{
		// Part 1
		input, err := os.Open("./day4/input.txt")
		handleError(err, ErrorContext{"4", "1", "could not open input"})

		res, err := day4.Part1(input)
		handleError(err, ErrorContext{"4", "1", ""})
		log.Printf("day 4: part 1: %d", res)

		// Part 2
		input, err = os.Open("./day4/input.txt")
		handleError(err, ErrorContext{"4", "2", "could not open input"})

		res, err = day4.Part2(input)
		handleError(err, ErrorContext{"4", "2", ""})
		log.Printf("day 4: part 3: %d", res)
	}

	// Day 5
	{
		// Part 1
		input, err := os.Open("./day5/input.txt")
		handleError(err, ErrorContext{"5", "1", "could not open input"})

		res, err := day5.Part1(input)
		handleError(err, ErrorContext{"5", "1", ""})
		log.Printf("day 5: part 1: %d", res)

		// Part 2
		input, err = os.Open("./day5/input.txt")
		handleError(err, ErrorContext{"5", "2", "could not open input"})

		res, err = day5.Part2(input)
		handleError(err, ErrorContext{"5", "2", ""})
		log.Printf("day 5: part 3: %d", res)
	}

	// Day 6
	{
		// Part 1
		input, err := os.Open("./day6/input.txt")
		handleError(err, ErrorContext{"6", "1", "could not open input"})

		res, err := day6.Part1(input)
		handleError(err, ErrorContext{"6", "1", ""})
		log.Printf("day 5: part 1: %d", res)

		// Part 2
		input, err = os.Open("./day6/input.txt")
		handleError(err, ErrorContext{"6", "2", "could not open input"})

		res, err = day6.Part2(input)
		handleError(err, ErrorContext{"6", "2", ""})
		log.Printf("day 6: part 3: %d", res)
	}
}
