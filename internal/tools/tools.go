package tools

import (
	"encoding/json"
	"github/cleverunemployed/ToDoGo/internal/models"
	"math/rand"
	"os"
)

func LoadTasks(filename string) ([]models.Task, error) {
	var tasks []models.Task

	file, err := os.Open(filename)
	if err != nil {
		return tasks, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&tasks)
	return tasks, err
}

func SaveTasks(filename string, tasks []models.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(tasks)
}

func GetRandomString() string {
	var length int8 = 12
	var symbols string = "qwertyuiop[]asdfghjkl;'zxcvbnm,./1234567890-=QWERTYUIOPLKJHGFDSAZXCVBNM"
	var lengthSymbols int = len(symbols)
	var resultString []rune

	for i := 0; i < int(length); i++ {
		resultString[i] = rune(symbols[rand.Intn(lengthSymbols)])
	}

	return string(resultString)
}
