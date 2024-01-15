package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(text string) int {
	return 0
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
