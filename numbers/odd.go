package numbers

// Odd returns true if the Integer value is not divisible by 2.
func Odd[E Integer](x E) bool {
	return x%2 != 0
}
