package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nums = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	for {
		scanner.Scan()
		text := scanner.Text()
		var firstChar int
		var lastChar int
		var lowestIdx int
		var highestIdx int
		var lowNumReplacement string
		var highNumReplacement string
		if len(text) != 0 {
			// First pass - go through the text and find the first occurrence
			// of each of the string numbers
			// Take the one with the lowest index and do a single replace
			// on the text on it.
			// Need to make sure a case like eightwo1 is handled correctly
			// This should transform first into
			// 8wo1
			// and then we pick out 81
			// important not to convert the "two" first into eigh21
			lowestIdx = len(text)
			for numSpelt, _ := range nums {
				index := strings.Index(text, numSpelt)
				if index == -1 {
					continue
				}
				if index < lowestIdx {
					lowestIdx = index
					lowNumReplacement = numSpelt
				}
			}
			replacedText := strings.Replace(text, lowNumReplacement, nums[lowNumReplacement], 1)
			// Second pass - now do the second replacement at the end of
			// the string
			// Once we've done this, the algorithm for the first part of the
			// problem should work as usual.
			// In this pass, we can replace all instances.
			for numSpelt := range nums {
				index := strings.LastIndex(replacedText, numSpelt)
				if index == -1 {
					continue
				}
				if index > highestIdx {
					highestIdx = index
					highNumReplacement = numSpelt
				}
			}
			replacedText = strings.Replace(replacedText, highNumReplacement, nums[highNumReplacement], -1)
			for _, char := range replacedText {
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
			// fmt.Println(text, replacedText, firstChar, lastChar)
		} else {
			break
		}
	}
	fmt.Println(answer)
}
