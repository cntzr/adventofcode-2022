package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: Maximum of calories: %d\n", part_one())         // 69528 Calories
	fmt.Printf("Part 2: Calories of the top 3 elves: %d\n", part_two()) // 206152 Calories
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

	maxCalories := 0
	curCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if curCalories > maxCalories {
				maxCalories = curCalories
			}
			curCalories = 0
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			curCalories += calories
		}
	}
	return maxCalories
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
	elves := make([]int, 0)

	maxCalories := 0
	curCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			if curCalories > maxCalories {
				maxCalories = curCalories
			}
			elves = append(elves, curCalories)
			curCalories = 0
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			curCalories += calories
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	topElves := 0
	for i := 0; i < 3; i++ {
		topElves += elves[i]
	}
	return topElves
}
