package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

	frequencies := make(map[int]bool)
	frequencies[0] = true

	sum := 0
	scanner := bufio.NewScanner(file)
	for true {
		for scanner.Scan() {
			temp, _ := strconv.Atoi(scanner.Text())
			sum += temp
			if frequencies[sum] == true {
				fmt.Println(sum)
				os.Exit(0)
			} else {
				frequencies[sum] = true
			}
		}
		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
	}

	fmt.Println("No duplicate frequency found")
}
