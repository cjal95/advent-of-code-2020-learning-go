package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Policy -
type Policy struct {
	Min      int
	Max      int
	Letter   string
	Password string
}

func checkPolicyPartOne(policies []Policy) int {
	result := 0
	for _, policy := range policies {
		count := len(strings.Split(policy.Password, policy.Letter)) - 1
		if count >= policy.Min && count <= policy.Max {
			result++
		}
	}
	return result
}

func checkPolicyPartTwo(policies []Policy) int {
	result := 0
	for _, policy := range policies {
		if policy.Min <= 0 || policy.Max > len(policy.Password) {
			continue
		}
		minLetter := policy.Password[policy.Min-1 : policy.Min]
		maxLetter := policy.Password[policy.Max-1 : policy.Max]
		count := 0
		if minLetter == policy.Letter {
			count++
		}
		if maxLetter == policy.Letter {
			count++
		}
		if count == 1 {
			result++
		}
	}
	return result
}

func getInt(val string) int {
	result, _ := strconv.Atoi(val)
	return result
}

func getInput() []string {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var passwords []string
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}
	return passwords
}

func getPolicies(rows []string) []Policy {
	var policies []Policy
	for _, row := range rows {
		minEnd := strings.Index(row, "-")
		min := getInt(row[:minEnd])
		row = row[minEnd+1:]
		maxEnd := strings.Index(row, " ")
		max := getInt(row[:maxEnd])
		row = row[maxEnd+1:]
		letter := row[:1]
		password := row[3:]
		policies = append(policies, Policy{
			Letter:   letter,
			Password: password,
			Min:      min,
			Max:      max,
		})
	}
	return policies
}

func main() {
	rows := getInput()
	polices := getPolicies(rows)
	fmt.Println("PartOne", checkPolicyPartOne(polices))
	fmt.Println("PartTwo", checkPolicyPartTwo(polices))
}
