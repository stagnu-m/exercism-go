package blackjack

var mapOfCards = map[string]int{
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
	return mapOfCards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if ParseCard(card1) == mapOfCards["ace"] && ParseCard(card2) == mapOfCards["ace"] {
		return "P"
	}
	dealer := ParseCard(dealerCard)
	hand := ParseCard(card1) + ParseCard(card2)
	switch {
	case hand == 21:
		if dealer == mapOfCards["ace"] || dealer == mapOfCards["ten"] {
			return "S"
		}
		return "W"
	case hand >= 17 && hand <= 20:
		return "S"
	case hand >= 12 && hand <= 16:
		if dealer >= 7 {
			return "H"
		}
		return "S"
	case hand <= 11:
		return "H"
	default:
		return ""
	}

}
