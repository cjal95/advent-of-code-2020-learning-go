package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"regexp"
)

type password struct {
	min, max         int
	letter, password string
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err) //If error in opening file
	}
	defer file.Close()
	passwords := []*password{}

	inputRegex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		min, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$1"))
		max, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$2"))
		passwords = append(passwords, &password{
			min:      min,
			max:      max,
			letter:   inputRegex.ReplaceAllString(text, "$3"),
			password: inputRegex.ReplaceAllString(text, "$4"),
		})
	}
	correct := 0
	for _, pw := range passwords {
		if count := strings.Count(pw.password, pw.letter); count >= pw.min && count <= pw.max {
		correct++
		}
	}
	fmt.Println(correct)
}