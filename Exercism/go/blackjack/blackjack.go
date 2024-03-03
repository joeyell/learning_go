package blackjack

var CardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return CardValues[card]
}

// getSum
func getSum(card1 string, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardValue := getSum(card1, card2)
	dealerValue := ParseCard(dealerCard)
	switch {
	case cardValue == 21 && dealerValue < 10:
		return "W"
	case cardValue == 21 && dealerValue >= 10:
		return "S"
	case card1 == "ace" && card2 == "ace":
		return "P"
	case 17 <= cardValue && cardValue <= 20:
		return "S"
	case 12 <= cardValue && cardValue <= 16 && dealerValue >= 7:
		return "H"
	case 12 <= cardValue && cardValue <= 16:
		return "S"
	case cardValue <= 11:
		return "H"
	default:
		return "cheese"
	}
}
