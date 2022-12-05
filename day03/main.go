package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: Priority sum is %d\n", part_one()) // Priority of 7872
	fmt.Printf("Part 2: Badge sum is %d\n", part_two())    // Badge sum of 2497
}

const (
	PRIO = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

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

	prioSum := 0
	for scanner.Scan() {
		completePack := scanner.Text()
		halfLength := len(completePack) / 2
		leftBag := completePack[:halfLength]
		rightBag := completePack[halfLength:]

		doubleChar := ""
		for _, ch := range leftBag {
			found := strings.ContainsRune(rightBag, ch)
			if found {
				doubleChar = string(ch)
				break
			}
		}

		priority := strings.Index(PRIO, doubleChar)
		if priority == -1 {
			fmt.Println("ILLEAL character found!")
			os.Exit(1)
		}
		prioSum += priority

		// fmt.Printf("%s = %s + %s ... Double = %s (%d)\n", completePack, leftBag, rightBag, doubleChar, priority)
	}

	return prioSum
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
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	badgeSum := 0
	i := 0
	entries := len(lines)
	for {
		if i >= entries {
			break
		}
		a := lines[i]
		i++
		if i >= entries {
			break
		}
		b := lines[i]
		i++
		if i >= entries {
			break
		}
		c := lines[i]
		i++

		badge := ""
		for _, ch := range a {
			foundInB := strings.ContainsRune(b, ch)
			foundInC := strings.ContainsRune(c, ch)
			if foundInB && foundInC {
				badge = string(ch)
				break
			}
		}
		if badge == "" {
			fmt.Println("No badge found!")
			os.Exit(1)
		}
		priority := strings.Index(PRIO, badge)
		if priority == -1 {
			fmt.Println("ILLEAL character found!")
			os.Exit(1)
		}
		badgeSum += priority
	}

	return badgeSum
}
