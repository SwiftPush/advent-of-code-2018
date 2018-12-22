package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type event struct {
	minute int
	action int
	guard  int
}

const (
	beginShift = iota
	fallAsleep
	wakeUp
)

func parseLines(lines []string) []event {
	// Parse lines into events
	events := []event{}
	currentGuard := -1
	for _, line := range lines {
		event := event{}

		var year, month, day, hour, minute int
		_, _ = fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)

		event.minute = minute

		endOfDateIndex := strings.Index(line, "]")
		line := line[endOfDateIndex+2:]
		firstWord := strings.Split(line, " ")[0]

		switch {
		case firstWord == "Guard":
			_, _ = fmt.Sscanf(line, "Guard #%d begins shift", &currentGuard)
			event.action = beginShift
		case firstWord == "wakes":
			event.action = wakeUp
		case firstWord == "falls":
			event.action = fallAsleep
		default:
			fmt.Println("Error")
		}
		event.guard = currentGuard
		events = append(events, event)
	}

	return events
}

func sleepiestGuard(events []event) int {
	// Calculate how much time each guard spends asleep
	guardSleepTime := map[int]int{}
	guardAsleepFrom := -1
	guardAsleepTo := -1
	currentGuard := 1
	for _, event := range events {
		switch {
		case event.action == beginShift:
			currentGuard = event.guard
		case event.action == wakeUp:
			guardAsleepTo = event.minute
			guardSleepTime[currentGuard] += guardAsleepTo - guardAsleepFrom
		case event.action == fallAsleep:
			guardAsleepFrom = event.minute
		default:
			fmt.Println("Error")
		}
	}

	// Find the sleepiest guard
	sleepiestGuard := 0
	for idx, elem := range guardSleepTime {
		if elem > guardSleepTime[sleepiestGuard] {
			sleepiestGuard = idx
		}
	}
	return sleepiestGuard
}

func sleepiestMinute(events []event, guard int) int {
	// Build an array of how often the guard is aleep for a given minute
	minutes := [60]int{}
	guardAsleepFrom := -1
	guardAsleepTo := -1
	currentGuard := -1
	for _, event := range events {
		if event.action == beginShift {
			currentGuard = event.guard
			continue
		}
		if currentGuard != guard {
			continue
		}

		switch {
		case event.action == wakeUp:
			guardAsleepTo = event.minute
			for i := guardAsleepFrom; i < guardAsleepTo; i++ {
				minutes[i]++
			}
		case event.action == fallAsleep:
			guardAsleepFrom = event.minute
		default:
			fmt.Println("Error")
		}
	}

	// Find the minute the guard is most commonly asleep
	sleepiestMinute := 0
	for idx, elem := range minutes {
		if elem > minutes[sleepiestMinute] {
			sleepiestMinute = idx
		}
	}

	return sleepiestMinute
}

func main() {
	// Check Command Line Arguments
	if len(os.Args) != 2 {
		log.Print("Expected filename")
		os.Exit(1)
	}

	// Open File
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Read file
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Sort by time
	sort.Strings(lines)

	events := parseLines(lines)

	sleepiestGuard := sleepiestGuard(events)
	sleepiestMinute := sleepiestMinute(events, sleepiestGuard)
	fmt.Printf("The sleepiest guard is %d\n", sleepiestGuard)
	fmt.Printf("The sleepiest minute is %d\n", sleepiestMinute)
	fmt.Printf("The answer is %d * %d = %d\n", sleepiestGuard, sleepiestMinute, sleepiestGuard*sleepiestMinute)
}
