package numbers

// Even returns true if the Integer value is divisible by 2.
func Even[E Integer](x E) bool {
	return x%2 == 0
}
