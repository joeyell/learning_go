// main.go

package week_2

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Create Person struct
type Person struct {
	Name    string
	Address string
}

func Makejson() {

	//Write a program which prompts the user to first enter a Name, and then enter an Address.
	//Your program should create a map and add the Name and Address to the map using the keys “Name” and “Address”, respectively.
	//Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

	// Initialise Person struct as User
	var User Person

	// display request and get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	User.Name, _ = reader.ReadString('\n')

	// display request and get user input
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your address: ")
	User.Address, _ = reader.ReadString('\n')

	// Trim new lines
	User.Name = strings.TrimSpace(User.Name)
	User.Address = strings.TrimSpace(User.Address)

	// Marshal User struct
	complete, err := json.Marshal(User)

	// Handle errors
	if err != nil {
		fmt.Println("You broke it?")
	}

	// Print JSON object
	fmt.Printf("Completed JSON is: %s", string(complete))
}
