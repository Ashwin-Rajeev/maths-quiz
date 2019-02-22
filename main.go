package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Quiz is a abstract struct for storing
// the quiz information.
type Quiz struct {
	Timer             *time.Timer
	CustomInputTime   time.Duration
	QuestionAndAnswer map[string]string
	TotalQuestions    int
	QuestionsAnswered int
	Marks             int
	Right             int
	Wrong             int
}

// newQuiz returns a new Quiz instance
// with initialized map parameter.
func newQuiz() *Quiz {
	return &Quiz{
		QuestionAndAnswer: make(map[string]string, 100),
		Timer:             time.NewTimer(5 * time.Minute),
	}
}

func main() {
	q := newQuiz()
	file := flag.String("file", "problems.csv", "Give the path of input csv file")
	t := flag.Int("time", 30, "Give custom game duration time(In Seconds)")
	flag.Parse()
	s := fmt.Sprintf("%ds", *t)
	duration, err := time.ParseDuration(s)
	if err != nil {
		fmt.Println(err)
	}
	q.CustomInputTime = duration
	q.readFromCSV(*file)
	q.TotalQuestions = len(q.QuestionAndAnswer)
	q.userInterface()
}

func (q *Quiz) userInterface() {
	var input int
	for {
		if !q.Timer.Reset(5 * time.Minute) {
			fmt.Println("Timer reseted failed")
		}
		fmt.Printf("\n\t\tWelcome to Online quiz \n\n")
		fmt.Printf("********************************************\n")
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
			}
		}
	}
}

// quizGame game is a function where the
// actual game logic is taken place.
func (q *Quiz) quizGame() {
	var input, flag string
	q.Marks = 0
	q.Right = 0
	q.Wrong = 0
e:
	fmt.Println("Shall we start the game? (y/n)")
	fmt.Scan(&flag)
	if flag != "y" && flag != "n" {
		fmt.Println("Invalid entry")
		goto e
	} else if flag == "n" {
		return
	}
	if q.Timer.Reset(q.CustomInputTime) {
		fmt.Printf("\nYour time start now...You have total of %d Seconds", q.CustomInputTime/1000000000)
	}
	// Creating a go routine for checking the time
	go func() {
		select {
		case <-q.Timer.C:
			q.Timer.Stop()
			fmt.Println("\nSorry time out, Please try again later!!!")
			fmt.Printf("\nTotal questions: %d , You Answered: %d", q.TotalQuestions, q.QuestionsAnswered)
			fmt.Printf("\nRight answers: %d, Wrong answers: %d", q.Right, q.Wrong)
			fmt.Println("\n\tYour total points: ", q.Marks)
			os.Exit(0)
		}
	}()

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
			q.QuestionsAnswered++
			continue
		}
		q.Wrong++
		q.QuestionsAnswered++
		fmt.Println("Wrong answer!")
		fmt.Printf("Right answer is %s \n", answer)
	}
	fmt.Printf("\nTotal questions: %d , You Answered: %d", q.TotalQuestions, q.QuestionsAnswered)
	fmt.Printf("\nRight answers: %d, Wrong answers: %d", q.Right, q.Wrong)
	fmt.Println("\n\tYour total points: ", q.Marks)
}
