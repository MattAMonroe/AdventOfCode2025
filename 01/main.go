package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	aoclib "github.com/MattAMonroe/AdventOfCode2025/AOCLib"
)

func main() {
	content := aoclib.ReadFile("sample.txt")
	_ = CountZeros(content)
}

type Direction int

func DirectionFromInt(a int) Direction {
	if a == DirLeft {
		return DirLeft
	} else {
		return DirRight
	}
}

const (
	DirLeft = iota
	DirRight
)

type Command struct {
	direction Direction
	clicks    int
}

func CountZeros(content string) int {
	position := 50
	commands := ProcessInput(content)
	numZero := 0
	for _, command := range commands {
		if command.direction == DirLeft {
			// subtraction
			clicks := command.clicks % 100
			position -= clicks
			if position < 0 {
				position += 100
			}
		} else {
			// addition
			position += command.clicks
			position %= 100
		}

		// check for 0
		if position == 0 {
			numZero++
		}
	}
	return numZero
}

func ProcessInput(content string) []Command {
	splits := strings.Split(content, "\n")
	commands := []Command{}
	for _, split := range splits {
		if split == "" {
			continue
		}
		var direction Direction = DirRight
		first := string(split[0])
		rest := split[1:]
		if first == "L" {
			direction = DirLeft
		}
		clicks, err := strconv.Atoi(rest)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse clicks: %s\n", rest)
		}

		commands = append(commands, Command{direction, clicks})
	}

	return commands
}
