package main

import "fmt"
import "os"
import "strconv"

func main() {
	/* Get list of numbers from cmd line */
	args := os.Args[1:]
	//fmt.Println(args)
	
	/* For each item in args, loop through each digit, and phonetically append to phoneticArgs list */
	//var phoneticArgs = []string{}
	for i := 0; i < len(args); i ++ {
		for j, c := range args[i] {
			//fmt.Println(j, "=>", string(c))
			num, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			fmt.Println(j, "=>", num)
		}
	}
}
