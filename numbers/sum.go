package numbers

func Sum[S ~[]E, E Number](s S) E {
	var result E

	for _, v := range s {
		result += v
	}

	return result
}
