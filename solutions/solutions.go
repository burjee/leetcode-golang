package solutions

import (
	"leetcode/solutions/s101_200"
	"leetcode/solutions/s1_100"
	"log"
)

func Run(n string) {
	switch n {
	case "1":
		s1_100.Run_1()
	case "100":
		s1_100.Run_100()
	case "101":
		s101_200.Run_101()
	default:
		log.Println("Solution not found.")
	}
}
