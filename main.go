package main

import (
	"fmt"
	"log"

	"github.com/PauloJunior5/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Paulo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
