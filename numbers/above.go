package numbers

// Above takes a min value to compare to and returns a function than compares
// a number against this value.
//
// The result is e > min
func Above[E Number](min E) func(E) bool {
	return func(e E) bool {
		return e > min
	}
}

// AboveOrEqual takes a min value to compare to and returns a function than
// compares a number against this value.
//
// The result is e >= min
func AboveOrEqual[E Number](min E) func(E) bool {
	return func(e E) bool {
		return e >= min
	}
}
