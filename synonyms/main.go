package main

import (
	"bufio"
	"fmt"
	"github.com/zucchidev/synonyms/thesaurus"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("BIG_HUGE_THESAURUS")
	if apiKey == "" {
		log.Fatalln("BIG_HUGE_THESAURUS environment variable is mandatory")
	}
	thes := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thes.Synonyms(word)
		if err != nil {
			log.Fatalln("Failer when looking for synonyms for "+word+" ", err.Error())
		}

		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for " + word + " ")
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
