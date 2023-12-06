package numbers

func Below[E Number](max E, includeEqual bool) func(E) bool {
	if includeEqual {
		return func(e E) bool {
			return e <= max
		}
	} else {
		return func(e E) bool {
			return e < max
		}
	}
}
