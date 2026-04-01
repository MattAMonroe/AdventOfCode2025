package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello World!")
}

func ReadBatteryBanks(content string) [][]int {
	batteries := [][]int{}

	splits := strings.Split(content, "\n")
	for _, split := range splits {
		if split == "" {
			continue
		}
		bank := []int{}
		for i := range split {
			v, err := strconv.Atoi(split[i : i+1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to process character %s: %v\n", split[i:i+1], err)
				continue
			}
			bank = append(bank, v)
		}
		batteries = append(batteries, bank)
	}

	return batteries
}

func FindBestCombo(bank []int) int {
	firstDigitIdx, secondDigitIdx := 0, 0

	cursor := 0
	for i, v := range bank {
		if i == len(bank)-1 {
			break
		}
		if v > cursor {
			firstDigitIdx = i
			cursor = v
		}
	}

	searchLen := len(bank) - firstDigitIdx - 1
	cursor = 0
	for i := range searchLen {
		loc := firstDigitIdx + i + 1
		if bank[loc] > cursor {
			secondDigitIdx = loc
			cursor = bank[loc]
		}
	}

	return bank[firstDigitIdx]*10 + bank[secondDigitIdx]
}

func SumBestCombos(batteries [][]int) int {
	sum := 0
	for _, bank := range batteries {

		sum += FindBestCombo(bank)
	}

	return sum
}
