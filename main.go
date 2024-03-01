package main

import (
	"fmt"
	"learning_go/Exercism/go/greeting"
	"learning_go/Exercism/go/lasagna"
	"learning_go/Exercism/go/partyrobot"
)

func main() {
	fmt.Println(greeting.HelloWorld())
	fmt.Println(lasagna.ElapsedTime(1, 30))
	fmt.Println(partyrobot.AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
