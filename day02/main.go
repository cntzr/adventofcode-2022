package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: My score is %d\n", part_one()) // 15523 points
	fmt.Printf("Part 2: My score is %d\n", part_two()) // 15702 points
}

func part_one() int {
	data := "data.txt"
	f, err := os.Open(data)
	if err != nil {
		fmt.Printf("Can't open file %s. Got error %q\n", data, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	score := 0
	for scanner.Scan() {
		switch scanner.Text() {
		case "A X":
			score += 3 + 1
		case "A Y":
			score += 6 + 2
		case "A Z":
			score += 0 + 3
		case "B X":
			score += 0 + 1
		case "B Y":
			score += 3 + 2
		case "B Z":
			score += 6 + 3
		case "C X":
			score += 6 + 1
		case "C Y":
			score += 0 + 2
		case "C Z":
			score += 3 + 3
		}
	}
	return score
}

func part_two() int {
	data := "data.txt"
	f, err := os.Open(data)
	if err != nil {
		fmt.Printf("Can't open file %s. Got error %q\n", data, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	score := 0
	for scanner.Scan() {
		switch scanner.Text() {
		case "A X": // lose against Rock with Scissors
			score += 0 + 3
		case "A Y": // draw against Rock with Rock
			score += 3 + 1
		case "A Z": // win against Rock with Paper
			score += 6 + 2
		case "B X": // lose against Paper with Rock
			score += 0 + 1
		case "B Y": // draw against Paper with Paper
			score += 3 + 2
		case "B Z": // win against Paper with Scissors
			score += 6 + 3
		case "C X": // lose against Scissors with Paper
			score += 0 + 2
		case "C Y": // draw against Scissors with Scissors
			score += 3 + 3
		case "C Z": // win against Scissors with Rock
			score += 6 + 1
		}
	}
	return score
}
