package main

import (
	"fmt"
	"math"
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

	secondDigitIdx = FindNextDigitLoc(bank[firstDigitIdx+1:])

	return bank[firstDigitIdx]*10 + bank[secondDigitIdx+1+firstDigitIdx]
}

func SumBestCombos(batteries [][]int) int {
	sum := 0
	for _, bank := range batteries {

		sum += FindBestCombo(bank)
	}

	return sum
}

func FindFirstDigitLoc(bank []int) int {
	foundIdx := 0
	cursor := 0
	for i := range len(bank) {
		checkPos := len(bank) - i - 1
		if bank[checkPos] >= cursor {
			cursor = bank[checkPos]
			foundIdx = checkPos
		}
	}

	return foundIdx
}

func FindNextDigitLoc(bank []int) int {
	digitIdx := 0

	searchLen := len(bank)
	cursor := 0

	for i := range searchLen {
		if bank[i] > cursor {
			digitIdx = i
			cursor = bank[i]
		}
	}

	return digitIdx
}

func SumBestLargeCombos(batteries [][]int) int {
	sum := 0

	for _, bank := range batteries {
		start := 0
		jolt := 0
		for i := range 12 {
			digitSliceIdx := FindNextDigitLoc(bank[start : len(bank)-11+i])
			digitIdx := digitSliceIdx + start
			digit := bank[digitIdx]
			jolt += int(math.Pow(10.0, 11.0-float64(i))) * digit

			start = digitIdx + 1

		}
		sum += jolt
	}

	return sum
}
