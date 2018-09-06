package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var marks = map[bool]string{true: "✔", false: "✖"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net:43"
	conn, err := net.Dial("tcp", whoisServer)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	newLine := "\n"
	conn.Write([]byte(domain + newLine))
	for scanner.Scan() {
		lowerText := strings.ToLower(scanner.Text())
		if strings.Contains(lowerText, "no match") {
			return false, nil
		}
	}
	return true, nil
}
