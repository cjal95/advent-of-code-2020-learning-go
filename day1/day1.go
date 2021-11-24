package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var target int = 2020

func part1(numbers []int) int {
	for  i := 0; i < len(numbers)-1; i++{
		for j := i+1; j < len(numbers)-1; j++{
			if numbers[i]+numbers[j] == target{
				return (numbers[i]*numbers[j])
			}
		}
	}
	return -1
}

func part2(numbers []int) int {
	for  i := 0; i < len(numbers)-1; i++{
		for j := i+1; j < len(numbers)-1; j++{
			for k := j+1; k < len(numbers)-1; k++{
				if numbers[i]+numbers[j] == target{
					return (numbers[i]*numbers[j]*numbers[k])
				}				
			}
		}
	}
	return -1
}

func main() {

	var result []int
	ints, err := os.Open("input.txt")
	
	if err != nil{
		log.Fatal(err) //If error in opening file
	}
	defer ints.Close()

	scanner := bufio.NewScanner(ints)

	for scanner.Scan(){
		x, err := strconv.Atoi(scanner.Text())
		if err != nil{
			log.Fatal(err)
		}
		result = append(result, x)
	}
	fmt.Println(part1(result))
	fmt.Println(part2(result))
}