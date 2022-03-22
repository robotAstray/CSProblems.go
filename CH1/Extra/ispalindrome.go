// Maintainer: RobotAstray
// Copyright (c) RobotAstray
// Usage: go build ispalindrome.go;  ./ispalindrome --phrase "enter your string here"

package main

import (
	"fmt"
	"log"
	"flag"
	"regexp"
	"os"
)


func main(){
	phrase := flag.String("phrase", "", "Enter a phrase here")
	flag.Parse()
	requiredFlags := []string{"phrase"}
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
	var s string 
	if *phrase != ""{

		r, err := regexp.Compile(`[A-Za-z]`)
		
		if err != nil {
			log.Println("Error, cannot regex compile: %v", err)
			os.Exit(1)
		}
		for _, i := range *phrase{
			if r.MatchString(string(i)){
				s += string(i)
			}
		}
		isPalindrome(s)
	}

}


func isPalindrome(s string){
	if s[0] == s[len(s)-1]{
		s = s[1:len(s)-1]
		if len(s) <=1{
			fmt.Printf("It's PALINDROME\n")
		}else{
			isPalindrome(s)
		}	
	} else{
		// return false 
		fmt.Printf("It's NOT PALINDROME\n")
	}
}