package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameResult struct {
	ID    int
	Count map[string]int
}

func solve(text string) int {
	parts := strings.Split(text, ":")
	gameID, _ := strconv.Atoi(strings.TrimSpace(parts[0][5:])) // Extract the ID from "Game X"
	counts := strings.Split(parts[1], ";")

	game := GameResult{
		ID:    gameID,
		Count: make(map[string]int),
	}

	for _, countStr := range counts {
		countsByColor := strings.Split(strings.TrimSpace(countStr), ",")
		for _, countPart := range countsByColor {
			countParts := strings.Split(strings.TrimSpace(countPart), " ")
			if len(countParts) == 2 {
				color := countParts[1]
				num, _ := strconv.Atoi(countParts[0])
				if num > game.Count[color] {
					game.Count[color] = num
				}
			}
		}
	}

	// fmt.Println(text, game.Count)

	return game.Count["red"] * game.Count["green"] * game.Count["blue"]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			answer += solve(text)
		} else {
			break
		}
	}
	fmt.Println(answer)
}
