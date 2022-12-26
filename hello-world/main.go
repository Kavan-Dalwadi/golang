package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello from the world")

	var firstName string

	fmt.Scan(&firstName)
	fmt.Printf("My first name is: %v\n", firstName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	
	// comma ok | error ok

	input, _ := reader.ReadString('\n')
    fmt.Println("Name: ",input)

}
