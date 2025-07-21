package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Initialization for custom variable types
type Sentence struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
}
type Answer struct {
	English   string
	Japanese  string
	Answer    string
	IsCorrect bool
}

func main() {
	//App Initialization & File Loading
	a := app.New()
	w := a.NewWindow("Nihongo Checks")
	file, err := os.ReadFile("data/references.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	sentences := []Sentence{}
	err = json.Unmarshal(file, &sentences)
	if err != nil {
		log.Fatalf("Error unmarshalling json: %v", err)
	}

	//Widget & Variables Initialization
	title := widget.NewLabel("Welcome to Nihongo Checks!")
	itemLabel := widget.NewLabel("Items:")
	itemEntry := widget.NewEntry()
	questionLabel := widget.NewLabel("Question: ")
	questionText := widget.NewLabel("")
	answerLabel := widget.NewLabel("Answer: ")
	answerEntry := widget.NewEntry()
	answerEntry.Disable()
	answerText := widget.NewLabel("")
	statusLabel := widget.NewLabel("")
	var itemInt int
	results := []Answer{}
	idx := rand.IntN(50) //Change based on how many sentences you have in your json
	score := 0
	x := 1

	//Button Initialization & Main Function of the App
	itemBtn := widget.NewButton("Enter", func() {
		itemText := itemEntry.Text
		val, err := strconv.Atoi(itemText)
		if err != nil {
			fmt.Println("Invalid number:", err)
			return
		}
		itemInt = val
		itemEntry.SetText("")
		itemEntry.Disable()
		answerEntry.Enable()
		questionText.SetText(sentences[idx].English)
	})
	ansBtn := widget.NewButton("Enter", func() {
		if itemInt > 0 {
			if x < itemInt {
				questionText.SetText(sentences[idx].English)
				ans := answerEntry.Text
				answerEntry.SetText("")
				answer := Answer{
					English:  sentences[idx].English,
					Japanese: sentences[idx].Japanese,
					Answer:   ans,
				}
				if answer.Japanese == answer.Answer {
					answer.IsCorrect = true
				} else if answer.Japanese != answer.Answer {
					answer.IsCorrect = false
				}
				results = append(results, answer)

				if answer.IsCorrect {
					answerText.SetText(fmt.Sprintf("English: %s\n CorrectJP: %s\n AnswerJP: %s\n Correct!\n",
						answer.English, answer.Japanese, answer.Answer))
					score += 1
				} else {
					answerText.SetText(fmt.Sprintf("English: %s\n CorrectJP: %s\n AnswerJP: %s\n Wrong!\n",
						answer.English, answer.Japanese, answer.Answer))
				}
				x++
				idx = rand.IntN(50)
				questionText.SetText(sentences[idx].English)
				statusLabel.SetText("")
			} else if x >= itemInt {
				statusLabel.SetText("All Done! Your score: " + strconv.Itoa(score) + "/" + strconv.Itoa(itemInt))
				answerEntry.Disable()
				answerText.SetText("")
				questionText.SetText("")
			}
		} else if itemInt <= 0 {
			statusLabel.SetText("Enter how many items first!")
		}

	})
	//Display all Widgets & Run App
	quiz := container.NewVBox(
		title,
		itemLabel,
		itemEntry,
		itemBtn,
		questionLabel,
		questionText,
		answerLabel,
		answerEntry,
		ansBtn,
		answerText,
		statusLabel,
	)
	w.SetContent(quiz)
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
