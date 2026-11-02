package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := ""
	for i := 0; i < numStarsPerLine; i++ {
		border += "*"
	}
	out := ""
	out += border + "\n"
	out += welcomeMsg + "\n"
	out += border
	return out
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	withoutStars := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(withoutStars)
}
