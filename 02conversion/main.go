package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, Give rating between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is: ", input)

	//rating, err := strconv.ParseFloat(input, 64)
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Your new rating is: ", rating+1)
	}

}
