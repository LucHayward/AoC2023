package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputTxt string

/*
Bag of cubes -> RGB
Play a game:
- Hide N cubes of each colour (Nr, Ng, Nb) in the bag
- Elf shows you K random cubes J times
*/
func main() {
	lines := strings.Split(inputTxt, "\n")

	// part1(lines)
	part2(lines)
}

// In Part2:
// Determine the _fewest number of cubes of each colour_ that could have been in the bag
// ie: Given all the groups find max([rgb])
func part2(lines []string) {
	powerSum := 0
	for _, line := range lines {
		cubes := strings.Split(line, " ")[2:]

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for i := 0; i < len(cubes); i += 2 { // iterate over the game 2 at a time for (num, colour)
			num, _ := strconv.Atoi(cubes[i])
			colour := rune(cubes[i+1][0])

			switch colour {
			case 'r':
				maxRed = max(maxRed, num)
			case 'g':
				maxGreen = max(maxGreen, num)
			case 'b':
				maxBlue = max(maxBlue, num)
			}
		}
		powerSum += maxRed * maxGreen * maxBlue

	}
	fmt.Println(powerSum)
}

// Given many games, determine which games would have been possible if the bag contained only 12R, 13G, 14B
// ie: determine which games R<=12, G<=13, B<=14
// Add the IDs of all these games
func part1(lines []string) {
	var goodIdsCnt int = 0
	for game, line := range lines {
		game = game
		cubes := strings.Split(line, " ")
		id, _ := strconv.Atoi(cubes[1][:len(cubes[1])-1])
		cubes = cubes[2:]

		goodGame := true
		for i := 0; i < len(cubes); i += 2 {
			num, _ := strconv.Atoi(cubes[i])
			colour := rune(cubes[i+1][0])

			if (num > 12 && colour == 'r') || (num > 13 && colour == 'g') || (num > 14 && colour == 'b') {
				goodGame = false
				fmt.Println(id)
				break
			}
		}
		if goodGame {
			goodIdsCnt += id
		}
	}
	fmt.Println(goodIdsCnt)
}
