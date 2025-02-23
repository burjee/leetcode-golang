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
	case "21":
		s1_100.Run_21()
	case "23":
		s1_100.Run_23()
	case "24":
		s1_100.Run_24()
	case "26":
		s1_100.Run_26()
	case "27":
		s1_100.Run_27()
	case "28":
		s1_100.Run_28()
	case "33":
		s1_100.Run_33()
	case "39":
		s1_100.Run_39()
	case "42":
		s1_100.Run_42()
	case "45":
		s1_100.Run_45()
	case "48":
		s1_100.Run_48()
	case "49":
		s1_100.Run_49()
	case "100":
		s1_100.Run_100()
	case "101":
		s101_200.Run_101()
	default:
		fmt.Println("Solution not found.")
	}
}
