package main

import (
	"fmt"
	"leetcode/solutions"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run . <problem_number>")
		return
	}

	solutions.Run(os.Args[1])
}
