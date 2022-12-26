package main

import "fmt"

func main() {
	fmt.Println("Structure Demo")

	Employee := User{"Kavan", 10, "kavan.sa@gmail.com", false}

	fmt.Println(Employee)

}

type User struct {
	Name     string
	Age      int
	Email    string
	Avaiable bool
}
