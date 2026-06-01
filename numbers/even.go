package numbers

// IsEven returns true if the Integer value is divisible by 2.
func IsEven[E Integer](x E) bool {
	return x%2 == 0
}
