package slt

func Mode[S ~[]E, E comparable](s S) (E, int) {
	var result E

	return result, 0
}

func ModeFunc[S ~[]E, E any, C comparable](s S, f func(E) C) (E, int) {
	var result E

	return result, 0
}
