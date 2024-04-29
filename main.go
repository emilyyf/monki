package main

import (
	"fmt"
	"monki/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! This is a basic REPL for the Monkey progamming language!\n", user.Username)
	fmt.Printf("Feel free to type some commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
