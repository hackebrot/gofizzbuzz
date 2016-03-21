package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hackebrot/gofizzbuzz/gofizzbuzz"
)

func main() {
	stringNumbers := os.Args[1:]
	for _, stringNumber := range stringNumbers {
		intNumber, err := strconv.Atoi(stringNumber)
		if err == nil {
			word := gofizzbuzz.GoFizzBuzz(intNumber)
			fmt.Println(word)
		}
	}
}
