package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcome := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return welcome
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var fancy string
	fancy = strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
	return fancy
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	fixed := strings.ReplaceAll(oldMsg, "*", "")
	fixed = strings.TrimSpace(fixed)
	return fixed
}
