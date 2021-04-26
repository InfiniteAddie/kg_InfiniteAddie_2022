package main

import "fmt"
import "os"
import "strconv"

func main() {
	/* Get list of numbers from cmd line */
	args := os.Args[1:]
	
	/* For each item in args, loop through each digit, and phonetically append to phoneticArgs list */
	var phoneticNums = []string{}
	//iterate number-by-number
	for i := 0; i < len(args); i ++ {
		var phoneticNum string = ""
		//iterate digit-by-digit
		for _, c := range args[i] {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			switch num {
			case 0:
				phoneticNum += "Zero"
			case 1:
				phoneticNum += "One"
			case 2:
				phoneticNum += "Two"
			case 3:
				phoneticNum += "Three"
			case 4:
				phoneticNum += "Four"
			case 5:
				phoneticNum += "Five"
			case 6:
				phoneticNum += "Six"
			case 7:
				phoneticNum += "Seven"
			case 8:
				phoneticNum += "Eight"
			case 9:
				phoneticNum += "Nine"
			}
		}
		phoneticNums = append(phoneticNums, phoneticNum)
	}

	//output resulting list of phonetic numbers
	for i := 0; i < len(phoneticNums); i ++ {
		if i != len(phoneticNums) - 1 {
			fmt.Print(phoneticNums[i], ", ")
		} else {
			fmt.Println(phoneticNums[i])
		}
	}
}
