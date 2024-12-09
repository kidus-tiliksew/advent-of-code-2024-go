package day9

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(input io.Reader) (int, error) {
	diskMap := ""

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		diskMap = scanner.Text()
	}

	blocks, count := GetBlocks(diskMap)

	right := len(blocks) - 1
	for i, e := range blocks {
		if i >= count {
			break
		}

		if e != '.' {
			continue
		}

		for blocks[right] == '.' {
			right--
		}

		blocks[i], blocks[right] = blocks[right], '.'
		right--
	}

	ans := 0
	for i, e := range blocks {
		if e == '.' {
			break
		}

		ans += ((int(e) - '0') * i)
	}

	return ans, nil
}

func Part2(input io.Reader) (int, error) {
	diskMap := ""

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		diskMap = scanner.Text()
	}

	blocks, _ := GetBlocks(diskMap)

	for right := len(blocks) - 1; right > 0; {
		if blocks[right] == '.' {
			right--
			continue
		}

		count := 0
		for j := right; j > 0 && blocks[j] == blocks[right]; j-- {
			count++
		}

		idx := -1

		// Search for blocks with good available space
		for i := 0; i < len(blocks) && i < right; i++ {
			if blocks[i] != '.' {
				continue
			}

			availableSpaces := 0
			for j := i; j < len(blocks) && blocks[j] == '.'; j++ {
				availableSpaces++
			}

			if availableSpaces >= count {
				idx = i
				break
			}
		}

		// If no good space is found, move to the next block
		if idx == -1 {
			curr := blocks[right]
			for right >= 0 && blocks[right] == curr {
				right--
			}
			continue
		}

		// Move the blocks in the available space
		for range count {
			blocks[idx], blocks[right] = blocks[right], '.'
			idx++
			right--
		}
	}

	ans := 0
	for i, e := range blocks {
		if e == '.' {
			continue
		}

		ans += ((int(e) - '0') * i)
	}

	return ans, nil
}

func GetBlocks(diskMap string) ([]rune, int) {
	blocks := []rune{}
	freeSpace, id, count := false, 0, 0
	for _, e := range diskMap {
		k, _ := strconv.Atoi(string(e))

		if freeSpace {
			for i := 0; i < k; i++ {
				blocks = append(blocks, '.')
			}
		} else {
			for i := 0; i < k; i++ {
				blocks = append(blocks, rune('0'+id))
				count++
			}

			id++
		}

		freeSpace = !freeSpace
	}

	return blocks, count
}
