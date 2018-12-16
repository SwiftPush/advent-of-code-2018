package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func distance(a string, b string) int {
	// assume a and b are same length
	result := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result++
		}
	}
	return result
}

func commonLetters(a string, b string) string {
	// assume a and b are same length
	result := ""
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result += string(a[i])
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Expected filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ids := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	for _, id1 := range ids {
		for _, id2 := range ids {
			d := distance(id1, id2)
			if d == 1 {
				fmt.Println(commonLetters(id1, id2))
				os.Exit(0)
			}
		}
	}
	//fmt.Println(checksum)
}
