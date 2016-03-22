package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hackebrot/gofizzbuzz/gofizzbuzz"
)

func main() {
	for _, s := range os.Args[1:] {
		i, err := strconv.Atoi(s)
		if err == nil {
			w := gofizzbuzz.GoFizzBuzz(i)
			fmt.Println(w)
		}
	}
}
