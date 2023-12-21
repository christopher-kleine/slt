package numbers

// Between takes a min and a max value, then returns a new function that takes an int.
//
// If lessOrEqual and moreOrEqual are true, the function return true if:
// e >= min && e <= max
//
// If lessOrEqual is true, but moreOrEqual is false, the function return true if:
// e >= min && e < max
//
// If lessOrEqual is false, but moreOrEqual is true, the function return true if:
// e > min && e <= max
//
// If lessOrEqual and moreOrEqual are false, the function return true if:
// e > min && e < max
func Between[E Number](min E, max E, lessOrEqual, moreOrEqual bool) func(E) bool {
	if lessOrEqual && moreOrEqual {
		return func(e E) bool {
			return e >= min && e <= max
		}
	}

	if lessOrEqual {
		return func(e E) bool {
			return e >= min && e < max
		}
	}

	if moreOrEqual {
		return func(e E) bool {
			return e > min && e <= max
		}
	}

	return func(e E) bool {
		return e > min && e < max
	}
}
