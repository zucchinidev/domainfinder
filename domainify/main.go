package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var topLevelDomains = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		var newText []rune
		text := strings.ToLower(s.Text())
		for _, letter := range text {
			if unicode.IsSpace(letter) {
				letter = '-'
			}
			if !strings.ContainsRune(allowedChars, letter) {
				continue
			}
			newText = append(newText, letter)
		}
		if newText[len(newText)-1] == '-' {
			newText = newText[:len(newText)-1]
		}
		fmt.Println(string(newText) + "." + topLevelDomains[rand.Intn(len(topLevelDomains))])
	}
}
