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
		log.Print("Expected filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		sum += temp
	}

	fmt.Println(sum)
}
