package list

func FlatMap[A any, B any](slice []A, f func(A) []B) []B {
	result := []B{}
	for _, v := range slice {
		result = append(result, f(v)...)
	}
	return result
}

func Map[A any, B any](slice []A, f func(A) B) []B {
	fn := func(a A) []B {
		return []B{f(a)}
	}
	return FlatMap(slice, fn)
}

func Filter[A any](slice []A, f func(A) bool) []A {
	fn := func(a A) []A {
		if f(a) {
			return []A{a}
		}
		return []A{}
	}
	return FlatMap(slice, fn)
}

func Exists[A any](slice []A, f func(A) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

func FoldLeft[A any, B any](slice []A, initVal B, accumulator func(B, A) B) B {
	result := initVal
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}

func Sort[A any](slice []A, f func(A, A) (A, A)) []A {

	return slice
}
