package main

import (
	"flag"
	"fmt"
	"log"
)

func isPalindrome(input string) bool {

	forwardRunes := []rune(input)
	reverseRunes := make([]rune, len(forwardRunes))

	for index := range forwardRunes {
		reverseRunes[index] = forwardRunes[len(forwardRunes)-index-1]
	}

	if string(reverseRunes) != string(forwardRunes) {
		return false
	}

	return true
}

// This approach should be faster than the above
func isPalindromeAlt(input string) bool {
	forwardRunes := []rune(input)
	for index, fr := range forwardRunes {
		rr := forwardRunes[len(forwardRunes)-1-index]
		if fr != rr {
			return false
		}
	}
	return true
}

func main() {
	wordPtr := flag.String("input", "", "a word or a number")

	flag.Parse()

	if len(*wordPtr) == 0 {
		log.Fatalf("Missing words")
	}

	result := isPalindrome(*wordPtr)
	if result != true {
		fmt.Printf("The word: %s is not a palindrome", *wordPtr)
	} else {
		fmt.Printf("The word: %s is a palindrome", *wordPtr)
	}

}
