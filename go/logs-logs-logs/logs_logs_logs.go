package logs

var chToApp = map[rune]string{
	'\u2757':     "recommendation",
	'\U0001F50D': "search",
	'\u2600':     "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		appName, ok := chToApp[r]
		if ok {
			return appName
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	b := []rune(log)
	for i, r := range b {
		if r == oldRune {
			b[i] = newRune
		}
	}
	return string(b)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}
