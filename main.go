package main

import (
	"fmt"
	"os"
	"strings"

	// alias imported library
	flag "github.com/ogier/pflag"
)

// declare variable to hold user input
var user string

func main() {
	fmt.Println("Welcome to github profile viewer")
	// parse supplied flags
	flag.Parse()

	// Print usage instruction if user does not supply a command line argument
	if flag.NFlag() == 0 {
		printUsageInstruction()
	}

	// if multiple users are passed separated by commas, store them in a list
	users := strings.Split(user, ",")

	/*
		function: pluralize
		description: pluralize a given word based on the amount of entries supplied
	*/
	pluralize := func(contentList []string, word string) (newWord string) {
		contentCount := len(contentList)
		newWord = word
		if contentCount > 1 {
			newWord = word + "s"
		}
		return
	}
	// Pluralize word if more than one
	pluralizedWord := pluralize(users, "user")
	fmt.Printf("Searching %s: %s\n", pluralizedWord, users)

	// "for... range" loop in GO allows us to iterate over each element of the array.
	// "range" keyword can return the index of the element (e.g. 0, 1, 2, 3 ...etc)
	// and it can return the actual value of the element.
	// Since GO does not allow unused variables, we use the "_" character to tell GO we don't care about the index, but
	// we want to get the actual user we're looping over to pass to the function.
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`Username:	`, result.Login)
		fmt.Println(`Name:		`, result.Name)
		fmt.Println(`Email:		`, result.Email)
		fmt.Println(`Bio:		`, result.Bio)
		fmt.Println("")
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "user1[, user2]")
}

func printUsageInstruction() {
	if flag.NFlag() == 0 {
		fmt.Println("You did not supply any flags.")
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
