package numbers

// Between takes a min and a max value, then returns a new function that takes an int.
// The new functions returns true: e >= min && e < max
func Between[E Number](min E, max E) func(E) bool {
	return func(e E) bool {
		return e >= min && e < max
	}
}
