package main

import (
	"flag"
	"fmt"
	"log"
)

func isPalindrome(input string) bool {

	forwardRunes := []rune(input)
	reverseRunes := make([]rune, len(input))

	for index := range input {
		reverseRunes[index] = forwardRunes[len(input)-1-index]
	}

	if string(reverseRunes) != string(forwardRunes) {
		return false
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
