package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	borderChars := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", borderChars, welcomeMsg, borderChars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Trim(strings.Trim(strings.Trim(oldMsg, "*"), "\n"), "*"))
}
