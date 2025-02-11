package solutions

import (
	"leetcode/solutions/s1_100"
	"log"
)

func Run(n string) {
	switch n {
	case "1":
		s1_100.Run_1()
	default:
		log.Println("Solution not found.")
	}
}
