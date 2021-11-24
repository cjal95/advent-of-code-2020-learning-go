package main

import (
	"bufio"
	"fmt"
	//	"fmt"
	"log"
	"os"
	"strconv"
	//	"strings"
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
			if result[i]+result[j] == 2020{
				fmt.Println(result[i])
				fmt.Println(result[j])
				fmt.Println(result[i]*result[j])
			}
		}
	}
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
}