package main

import (
	"fmt"
	"learning_go/Exercism/go/blackjack"
)

func main() {
	// fmt.Println(greeting.HelloWorld())
	// fmt.Println(lasagna.ElapsedTime(1, 30))
	// fmt.Println(partyrobot.AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
	fmt.Println(blackjack.FirstTurn("ten", "four", "six"))
}

// --- FAIL: TestFirstTurn (0.00s)
//     --- FAIL: TestFirstTurn/blackjack_with_ace_for_dealer (0.00s)
//         blackjack_test.go:297: FirstTurn(ace, jack, ace) = W, want S
//     --- FAIL: TestFirstTurn/blackjack_with_queen_for_dealer (0.00s)
//         blackjack_test.go:297: FirstTurn(king, ace, queen) = W, want S
//     --- FAIL: TestFirstTurn/score_of_16_with_six_for_dealer (0.00s)
//         blackjack_test.go:297: FirstTurn(ten, six, six) = H, want S
//     --- FAIL: TestFirstTurn/score_of_15_with_six_for_dealer (0.00s)
//         blackjack_test.go:297: FirstTurn(ten, five, six) = H, want S
//     --- FAIL: TestFirstTurn/score_of_14_with_six_for_dealer (0.00s)
//         blackjack_test.go:297: FirstTurn(ten, four, six) = H, want S
//     --- FAIL: TestFirstTurn/score_of_13_with_six_for_dealer (0.00s)
//         blackjack_test.go:297: FirstTurn(ten, three, six) = H, want S
//     --- FAIL: TestFirstTurn/score_of_12_with_six_for_dealer (0.00s)
