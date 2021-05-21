package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	quizQuestions := csvReader()
	userFlagTime := timeout()
	total := len(quizQuestions)
	correct := 0
	f := func() {
		printResults(correct, total)
	}
	time.AfterFunc(time.Duration(userFlagTime), f)
	questionAsker(quizQuestions, &correct, f)
}

func printResults(correct int, total int) {
	stringToPrint := fmt.Sprintln("You got", correct, "out of", total)
	fmt.Println(stringToPrint)
}

func timeout() int {
	fmt.Println("How many seconds should you be given?")
	userInputReader := bufio.NewReader(os.Stdin)
	input, err := userInputReader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	num, numErr := strconv.Atoi(input)

	if err != nil || numErr != nil || num <= 0 {
		return 30
	}
	return num
}

func questionAsker(questions [][]string, correct *int, completionMessage func()) {
	for _, item := range questions {
		question := item[0]
		answer := item[1]

		fmt.Println("What is: " + question)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		input = strings.TrimSuffix(input, "\n")

		if err != nil {
			fmt.Println("Error reading input answer")
			*correct = 0
		}

		if _, err := strconv.Atoi(input); err != nil {
			fmt.Println("not a number you have failed the quiz")
			*correct = 0
		}

		input = strings.TrimSuffix(input, "\n")

		if input == answer {
			*correct++
		}
	}
	completionMessage()
}

func csvReader() [][]string {
	fmt.Println("Type in a file to read")
	userInputReader := bufio.NewReader(os.Stdin)
	input, err := userInputReader.ReadString('\n')

	var recordFile *os.File
	var fileError error

	if err == nil {
		testPath := "./" + strings.TrimSuffix(input, "\n") + ".csv"
		fmt.Println(testPath)
		recordFile, fileError = os.Open(testPath)
	} else {
		fmt.Println("An error occured trying to read problems csv file defaulting to default questions")
		recordFile, fileError = os.Open("./problems.csv")
	}

	if fileError != nil {
		fmt.Println("An error occured trying to read problems csv file defaulting to default questions")
		recordFile, fileError = os.Open("./problems.csv")
	}

	if fileError != nil {
		fmt.Println("Could not run program")
		return nil
	}

	reader := csv.NewReader(recordFile)

	records, _ := reader.ReadAll()
	return records
}
