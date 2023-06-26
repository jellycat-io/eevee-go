package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jellycat-io/eevee/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Eevee!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
