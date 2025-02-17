package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestLoadQuestions(t *testing.T) {
	// Create a temporary JSON file with test data
	testData := `[{"question": "What is 2+2?", "answer": "4"}, {"question": "What is the capital of France?", "answer": "paris"}]`
	tmpfile, err := ioutil.TempFile("", "testdata*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testData)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Load questions from the temporary file
	questions, err := loadQuestions(tmpfile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the questions are loaded correctly
	if len(questions) != 2 {
		t.Fatalf("Expected 2 questions, got %d", len(questions))
	}

	expectedQuestion := "What is 2+2?"
	if questions[0].Question != expectedQuestion {
		t.Errorf("Expected question %q, got %q", expectedQuestion, questions[0].Question)
	}

	expectedAnswer := "4"
	if questions[0].Answer != expectedAnswer {
		t.Errorf("Expected answer %q, got %q", expectedAnswer, questions[0].Answer)
	}
}

func TestGetRandomQuestion(t *testing.T) {
	questions := []Question{
		{Question: "What is 2+2?", Answer: "4"},
		{Question: "What is the capital of France?", Answer: "paris"},
	}

	// Seed the random number generator for reproducibility
	rand.Seed(time.Now().UnixNano())

	// Get a random question
	randomQuestion := getRandomQuestion(questions)

	// Check if the random question is one of the questions in the list
	found := false
	for _, q := range questions {
		if q.Question == randomQuestion.Question && q.Answer == randomQuestion.Answer {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Random question %v not found in the original list", randomQuestion)
	}
}
