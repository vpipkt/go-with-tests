package repeater

// Return a given string repeated so many times
func Repeat(s string, n int) string {
	var repeated string
	for range n {
		repeated += s
	}
	return repeated
}
