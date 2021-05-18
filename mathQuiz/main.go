package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	quizQuestions := csvReader()
	fmt.Println(quizQuestions)
	total := len(quizQuestions)
	correct := questionAsker(quizQuestions)
	stringToPrint := fmt.Sprintln("You got", correct, "out of", total)
	fmt.Println(stringToPrint)
}

func questionAsker(questions [][]string) int {
	userCorrect := 0
	for _, item := range questions {
		question := item[0]
		answer := item[1]

		fmt.Println("What is: " + question)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		input = strings.TrimSuffix(input, "\n")

		if err != nil {
			fmt.Println("Error reading input answer")
			return 0
		}

		if _, err := strconv.Atoi(input); err != nil {
			fmt.Println("not a number you have failed the quiz")
			return 0
		}

		input = strings.TrimSuffix(input, "\n")

		if input == answer {
			userCorrect++
		}
	}
	return userCorrect
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
