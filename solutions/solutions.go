package solutions

import (
	"leetcode/solutions/s0"
	"log"
)

func Run(n string) {
	switch n {
	case "1":
		s0.Run_1()
	default:
		log.Println("Solution not found.")
	}
}
