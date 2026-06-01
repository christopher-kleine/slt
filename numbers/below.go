package numbers

// Below takes a max value to compare to and returns a function than
// compares a number against this value.
//
// The result is e < max
func Below[E Number](max E, includeEqual bool) func(E) bool {
	return func(e E) bool {
		return e < max
	}
}

// BelowOrEqual takes a max value to compare to and returns a function than
// compares a number against this value.
//
// The result is e <= max
func BelowOrEqual[E Number](max E) func(E) bool {
	return func(e E) bool {
		return e <= max
	}
}
