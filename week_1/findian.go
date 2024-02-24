//findian.go

//Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!”
//if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
//The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

package week_1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Findian() {

	// display request and get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input a string: ")
	userInput, _ := reader.ReadString('\n')

	// normalise input
	lowerCaseInput := strings.ToLower(userInput)
	normalisedInput := strings.TrimSpace(lowerCaseInput)

	// Test for each i, a and n
	iPrefix := strings.HasPrefix(normalisedInput, "i")
	iSuffix := strings.HasSuffix(normalisedInput, "n")
	aContains := strings.Contains(normalisedInput, "a")

	if iPrefix && iSuffix && aContains {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
