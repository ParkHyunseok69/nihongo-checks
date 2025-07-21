# Nihongo Checks 🈶

A simple Japanese-to-English translation quiz app built with Go and the Fyne GUI toolkit.

## ✨ Features

- Loads Japanese-English sentence pairs from a JSON file.
- Asks the user to translate Japanese to English.
- Tracks number of items to quiz and provides feedback per item.
- Displays total score at the end.
- Fully GUI-based with Fyne.

## 🛠️ Requirements

- Go 1.18+ (due to use of `math/rand/v2`)
- [Fyne v2](https://developer.fyne.io/started/)

## 📁 Project Structure


nihongo-checks/
├── main.go
├── go.mod
├── data/
│   └── references.json


> Make sure you place `references.json` inside a `data/` folder. This file should contain an array of objects like:
```json
[
  { "english": "Good morning", "japanese": "おはよう" },
  { "english": "Thank you", "japanese": "ありがとう" }
]
````

## 🚀 Running the App

go run main.go


Make sure your Go environment is properly set up and you’ve installed the Fyne library.

## 📦 Installation

go get fyne.io/fyne/v2



