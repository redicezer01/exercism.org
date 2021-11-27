package techpalace

import "fmt"
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := make([]rune, numStarsPerLine, numStarsPerLine)
	for i := 0; i < numStarsPerLine; i++ {
		stars[i] = '*'
	}
	return strings.Join([]string{string(stars), "\n", welcomeMsg, "\n", string(stars)}, "")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	res := strings.Split(oldMsg, "\n")
	msg := strings.Join(res, "")
	return strings.Trim(strings.Trim(msg, "*"), " ")
}
