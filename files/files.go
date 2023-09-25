package files

import (
	"bufio"
	"fmt"

	"os"

	"github.com/chriscarsam/gofrom0/exercises"
)

var fileName string = "./files/txt/table.txt"

func SaveTable() {
	var text string = exercises.MultiplicationTable()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error when creating the file " + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func AddTable() {
	var text string = exercises.MultiplicationTable()
	if !Append(fileName, text) {
		fmt.Println("Error when concatenating content")
	}
}

func Append(fileN, text string) bool {
	file, err := os.OpenFile(fileN, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error during Append" + err.Error())
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error during WriteString" + err.Error())
		return false
	}

	file.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("There was an error reading the file " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		registration := scanner.Text()
		fmt.Println("> " + registration)
	}
	file.Close()
}
