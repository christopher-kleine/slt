package numbers

func IsDividableBy[E Integer](value E) func(E) bool {
	return func(e E) bool {
		return e%value == 0
	}
}
