package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var message = "Welcome to the Tech Palace"
    var customerUppercase = strings.ToUpper(customer)
    s := []string{message, customerUppercase}
    return strings.Join(s, ", ")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars = strings.Repeat("*", numStarsPerLine)
    var borderMsg = []string{stars, welcomeMsg, stars}
    return strings.Join(borderMsg, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var trimmed = strings.Trim(oldMsg, "*")
	var trimmedAgain = strings.Trim(trimmed, "\n")
    var trimmPlease = strings.Trim(trimmedAgain, "*")
    var trimmFinal = strings.Trim(trimmPlease, " ")
    return trimmFinal
    }
