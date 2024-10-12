package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"math/rand"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	shuffle := flag.Bool("shuffle", false, "a boolean value to shuffle the questions")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	if *shuffle {
		// Shuffle the questions
		problems = shuffleLines(problems)
	}

	correct := 0

problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		// Create a channel to handle the user input
		answerCh := make(chan string)
		// Create a goroutine to handle the user input
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			// Send the answer to the channel
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("")
			break problemLoop
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == problem.answer {
				correct++
			}
		}

	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
}

func shuffleLines(problems []problem) []problem {
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})
	return problems
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
