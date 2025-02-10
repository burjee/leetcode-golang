package main

import (
	"leetcode/solutions"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("go run . <problem_number>")
		return
	}

	solutions.Run(os.Args[1])
}
