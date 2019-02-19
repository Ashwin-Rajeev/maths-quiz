package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Quiz is a abstract struct for storing
// the quiz information.
type Quiz struct {
	QuestionAndAnswer map[string]string
	Marks             int
	Right             int
	Wrong             int
}

// newQuiz returns a new Quiz instance
// with initialized map parameter.
func newQuiz() *Quiz {
	return &Quiz{
		QuestionAndAnswer: make(map[string]string, 100),
	}
}

func main() {
	q := newQuiz()
	q.readFromCSV("problems.csv")
	var input int
	for {
		fmt.Printf("\n\t\tWelcome to Online quiz \n\n")
		fmt.Printf("**********************************************\n")
		fmt.Printf("\n\t\t1.Start quiz\n")
		fmt.Printf("\t\t2.Exit\n\n")
		fmt.Printf("\t\tEnter your choice:\n\n")
		fmt.Scan(&input)
		switch input {
		case 1:
			q.quizGame()
		default:
			os.Exit(0)
		}
	}
}

// readFromCSV read data from a csv file based on our input.
func (q *Quiz) readFromCSV(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(file))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, k := range records {
		for j, l := range k {
			if j == 0 {
				q.QuestionAndAnswer[l] = records[i][j+1]
				continue
			}
		}
	}
}

// quizGame game is a function where the
// actual game logic is taken place.
func (q *Quiz) quizGame() {
	q.Marks = 0
	q.Right = 0
	q.Wrong = 0
	var input string
	for question, answer := range q.QuestionAndAnswer {
	k:
		fmt.Printf("\nwhat is the result of %s : ", question)
		fmt.Scan(&input)
		inp, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input!")
			goto k
		}
		ans, err := strconv.Atoi(answer)
		if err != nil {
			log.Fatal(err)
		}
		if inp == ans {
			fmt.Println("Right answer!")
			q.Right++
			q.Marks++
			continue
		}
		q.Wrong++
		fmt.Println("Wrong answer!")
		fmt.Printf("Right answer is %s \n", answer)
	}
	fmt.Printf("\nRight answers: %d, Wrong answers: %d", q.Right, q.Wrong)
	fmt.Printf("\n\tYour total points: %d", q.Marks)
}
