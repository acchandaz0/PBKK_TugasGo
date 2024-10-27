package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Split(input, " ")

	var reversedWords []string
	for _, word := range words {
		reversedWords = append(reversedWords, reverseWord(word))
	}

	output := strings.Join(reversedWords, " ")
	fmt.Println("Output:", output)
}
