package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Question struct represents a question-answer pair
type Question struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// loadQuestions loads questions from the JSON file
func loadQuestions(filename string) ([]Question, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var questions []Question
	err = json.Unmarshal(file, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

// getRandomQuestion selects a random question from the list
func getRandomQuestion(questions []Question) Question {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return questions[r.Intn(len(questions))]
}

func main() {
	// Load questions from the quiz.json file
	questions, err := loadQuestions("/Users/kyle/Documents/Code/go/quiz-me/quiz.json")
	if err != nil {
		fmt.Println("Error loading questions:", err)
		os.Exit(1)
	}

	// Select a random question
	randomQuestion := getRandomQuestion(questions)

	// Display the question
	fmt.Println("Question:")
	fmt.Println(randomQuestion.Question)

	// Get the user's answer
	var userAnswer string
	fmt.Print("Your answer: ")
	fmt.Scanln(&userAnswer)

	// Clean the answer (to lowercase) and check if it's correct
	userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))

	// Check if the answer is correct
	if userAnswer == randomQuestion.Answer {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Wrong! Correct answer is: %s\n", randomQuestion.Answer)
	}
}
