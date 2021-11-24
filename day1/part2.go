package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
	for  i := 0; i < len(result)-1; i++{
		for j := i; j < len(result)-1; j++{
			for k := j; k < len(result)-1; k++{
				if result[i]+result[j]+result[k] == 2020{
					fmt.Println(result[i])
					fmt.Println(result[j])
					fmt.Println(result[k])
					fmt.Println(result[i]*result[j]*result[k])
				}	
			}
		}
	}
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
}