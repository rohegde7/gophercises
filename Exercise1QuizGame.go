package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bufio"
	"io"
)

func main() {

	problemsCsvFile, err := ioutil.ReadFile("problems.csv")
	problemsCsvFileInString := string(problemsCsvFile)

	keyboardReader := bufio.NewReader(os.Stdin)

	noOfCorrectAnswers := 0

	if err != nil {
		panic(err)
	}

	fmt.Println("__________my program starts here_____________")

	currentReader := csv.NewReader(strings.NewReader(problemsCsvFileInString))

	for {
		currentRow, err := currentReader.Read()

		if err == io.EOF {
			fmt.Println("Total correct answers: ", noOfCorrectAnswers)
			break
		}

		question := currentRow[0]
		correctAnswer := currentRow[1]

		fmt.Println(question)

		userAnswer, _, _ := keyboardReader.ReadLine()
		//userAnswer = strings.TrimSuffix(userAnswer, "\n")

		if string(userAnswer) == correctAnswer {
			noOfCorrectAnswers++
		}
	}
}
