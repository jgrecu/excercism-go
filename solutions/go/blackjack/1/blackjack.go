package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Calculate card values
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	// totalValue value of player's cards
	totalValue := value1 + value2

	// Check for pair of aces
	if card1 == "ace" && card2 == "ace" {
		return "P" // Always split
	}

	// Check for Blackjack
	if totalValue == 21 {
		// If dealer doesn't have ace, 10 or face card, automatically win
		if dealerValue != 11 && dealerValue != 10 {
			return "W"
		}
		return "S" // Otherwise, stand
	}

	// Strategy based on total value
	if totalValue >= 17 && totalValue <= 20 {
		return "S" // Always stand
	}

	if totalValue >= 12 && totalValue <= 16 {
		// Stand if dealer's card is less than 7
		if dealerValue < 7 {
			return "S"
		}
		return "H" // Hit if dealer has 7 or higher
	}

	// Always hit if total is 11 or lower
	return "H"
}
