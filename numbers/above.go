package numbers

func Above[E Number](min E, includeEqual bool) func(E) bool {
	if includeEqual {
		return func(e E) bool {
			return e >= min
		}
	} else {
		return func(e E) bool {
			return e > min
		}
	}
}
