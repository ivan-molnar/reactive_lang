package main

import (
	"fmt"
	"os"
	"os/user"
	"reactive/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the new reactive programming language!\n", user.Username)
	fmt.Printf("for now this prompt should print out tokens generated from your input.\nFeel free to give it a try!")
	repl.Start(os.Stdin, os.Stdout)
}
