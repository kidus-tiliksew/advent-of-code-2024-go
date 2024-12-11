package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kidus-tiliksew/advent-of-code-2024-go/day1"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day10"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day11"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day2"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day3"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day4"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day5"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day6"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day7"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day8"
	"github.com/kidus-tiliksew/advent-of-code-2024-go/day9"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "input",
				Usage: "input file path",
			},
			&cli.StringFlag{
				Name:  "day",
				Usage: "day number",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.String("input") == "" {
				return fmt.Errorf("input file path is required")
			}

			if cmd.String("day") == "" {
				return fmt.Errorf("day number is required")
			}

			input := cmd.String("input")
			day := cmd.String("day")

			file, err := os.Open(input)
			if err != nil {
				return fmt.Errorf("could not open input: %w", err)
			}
			handleError(err, ErrorContext{day, day, "could not open input"})

			switch day {
			case "1":
				res, err := day1.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 1: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day1.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 1: part 2: %d", res)
			case "2":
				res, err := day2.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 2: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day2.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 2: part 2: %d", res)

			case "3":
				res, err := day3.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 3: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day3.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 3: part 2: %d", res)

			case "4":
				res, err := day4.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 4: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day4.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 4: part 2: %d", res)

			case "5":
				res, err := day5.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 5: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day5.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 5: part 2: %d", res)

			case "6":
				res, err := day6.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 6: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day6.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 6: part 2: %d", res)
			case "7":
				res, err := day7.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 7: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day7.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 7: part 2: %d", res)
			case "8":
				res, err := day8.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 8: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day8.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 8: part 2: %d", res)
			case "9":
				res, err := day9.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 9: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day9.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 9: part 2: %d", res)
			case "10":
				res, err := day10.Part1(file)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 10: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day10.Part2(file)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 10: part 2: %d", res)
			case "11":
				res, err := day11.Part1(file, 25)
				handleError(err, ErrorContext{day, "1", ""})
				log.Printf("day 11: part 1: %d", res)

				file, _ := os.Open(input)

				res, err = day11.Part2(file, 75)
				handleError(err, ErrorContext{day, "2", ""})
				log.Printf("day 11: part 2: %d", res)
			default:
				return fmt.Errorf("day %s is not implemented", day)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
