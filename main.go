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

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
