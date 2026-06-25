package slt

type Slice[T any] []T

func NewSlice[S ~[]T, T any](s S) Slice[T] {
	return Slice[T](s)
}

func (s Slice[T]) Filter(f func(T) bool) Slice[T] {
	result := make(Slice[T], 0, len(s))

	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s Slice[T]) Reject(f func(T) bool) Slice[T] {
	result := make(Slice[T], 0, len(s))

	for _, v := range s {
		if !f(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s Slice[T]) Map[R any](f func(T) R) Slice[R] {
	result := make(Slice[R], len(s))

	for k, v := range s {
		result[k] = f(v)
	}

	return result
}

func (s Slice[T]) Add(v ...T) Slice[T] {
	return append(s, v...)
}

func (s Slice[T]) Unique[R comparable](f func(T) R) Slice[T] {
	result := make(Slice[T], 0, len(s))

	lookup := make(map[R]struct{})
	for _, v := range s {
		k := f(v)
		if _, ok := lookup[k]; !ok {
			lookup[k] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
