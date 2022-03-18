package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)



func main() {
	num := flag.String("num", "", "enter a number")
	flag.Parse()
	requiredFlags := []string{"num"}
	seenFlag := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		seenFlag[f.Name] = true
	})
	for _, reqFlag := range requiredFlags {
		if !seenFlag[reqFlag] {
			log.Printf("ERROR!!Missing mandatory flag: %s", reqFlag)
			os.Exit(1)
		}
	}
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		} else{
			return fibonacci(n-1) + fibonacci(n-2)
		}
	
	}
	if *num != "" {
		intNum, err := strconv.Atoi(*num)
	
		if err == nil {
			fmt.Println(intNum)
			fmt.Println(fibonacci(intNum))
		} else {
			fmt.Printf("You must enter an integer!\n")
		}

	}

}
