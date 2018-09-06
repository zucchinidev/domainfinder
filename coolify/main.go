package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	mustDuplicateVowel bool = true
	mustRemoveVowel    bool = false
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			index := getVowelIndex(word)
			isVowel := index > 0
			if isVowel {
				switch randBool() {
				case mustDuplicateVowel:
					word = append(word[:index+1], word[index:]...)
				case mustRemoveVowel:
					word = append(word[:index], word[index+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}

func randBool() bool {
	return rand.Intn(2) == 0
}

func getVowelIndex(word []byte) int {
	var index int = -1
	for i, char := range word {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if randBool() {
				index = i
			}
		}
	}
	return index
}
