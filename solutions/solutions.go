package solutions

import (
	"fmt"
	"leetcode/solutions/s101_200"
	"leetcode/solutions/s1_100"
)

func Run(n string) {
	switch n {
	case "1":
		s1_100.Run_1()
	case "2":
		s1_100.Run_2()
	case "3":
		s1_100.Run_3()
	case "5":
		s1_100.Run_5()
	case "6":
		s1_100.Run_6()
	case "9":
		s1_100.Run_9()
	case "11":
		s1_100.Run_11()
	case "12":
		s1_100.Run_12()
	case "13":
		s1_100.Run_13()
	case "14":
		s1_100.Run_14()
	case "15":
		s1_100.Run_15()
	case "17":
		s1_100.Run_17()
	case "19":
		s1_100.Run_19()
	case "20":
		s1_100.Run_20()
	case "100":
		s1_100.Run_100()
	case "101":
		s101_200.Run_101()
	default:
		fmt.Println("Solution not found.")
	}
}
