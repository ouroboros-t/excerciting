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
case "ten":
    return 10
case "jack":
    return 10
case "queen":
    return 10
case "king":
    return 10
default:
    return 0
    }

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    
	var card1Val = ParseCard(card1)
    var card2Val = ParseCard(card2)
    var dealerCardVal = ParseCard(dealerCard)
    
    var sum = card1Val + card2Val
	var pairOfAces = (sum == 22)
    var dealerHasTenOrAce = (dealerCardVal == 11) || (dealerCardVal == 10)
    var hasBlackJack = (sum == 21)
    var hasToStand = hasBlackJack && dealerHasTenOrAce
    var autoWin = hasBlackJack && !dealerHasTenOrAce
    var inStandRange = sum >= 17 && sum <= 20
    var inHitRange = sum >=12 && sum <=16
    var shouldAlwaysStand = inStandRange || (inHitRange && (dealerCardVal < 7)) 
    var shouldAlwaysHit = (inHitRange && dealerCardVal >= 7) || sum <= 11
    
    var action = ""
    switch {
        case pairOfAces:
        	action = "P"
        case autoWin:
        	action = "W"
        case hasToStand:
        	action = "S"
        case shouldAlwaysStand:
        	action = "S"
        case shouldAlwaysHit:
        	action = "H"
        
    }
    return action
   
}
