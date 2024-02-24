// read.go

package week_1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Fname string
	Lname string
}

func Read() {

	// display request and get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the file name: ")
	filename, _ := reader.ReadString('\n')

	// Trim new lines
	filename = strings.TrimSpace(filename)

	// Open file
	file, err := os.Open(filename)

	// Handle file opening error
	if err != nil {
		panic(err)
	}

	// Initialise struct and slice
	var name User
	var UsersSli []User

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			// Extract first and last names from the line
			name.Fname = parts[0]
			name.Lname = parts[1]
			// Append struct to slice
			UsersSli = append(UsersSli, name)
		}

	}
	// Print each struct from slice
	for _, j := range UsersSli {
		fmt.Print(j, "\n")
	}
	// close file
	file.Close()

}
