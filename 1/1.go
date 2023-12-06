package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	for {
		scanner.Scan()
		text := scanner.Text()
		var firstChar int
		var lastChar int
		if len(text) != 0 {
			for _, char := range text {
				if val, err := strconv.Atoi(string(char)); err == nil {
					// Doesn't correctly deal with the case where the string
					// has 0 as the first digit... but the input
					// file doesn't have any of those.
					if firstChar == 0 {
						firstChar = val
					}
					lastChar = val
				}
			}
			answer += 10*firstChar + lastChar
		} else {
			break
		}
	}
	fmt.Println(answer)
}
