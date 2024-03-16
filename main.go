package main

import (
	"fmt"

	"github.com/PauloJunior5/greetings"
)

func main() {
	message := greetings.Hello("teste")
	fmt.Println(message)
}
