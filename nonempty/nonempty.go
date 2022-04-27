package nonempty

type NonEmptyList[A any] struct {
	Head A
	Tail []A
}

func One[A any](a A) NonEmptyList[A] {
	return NonEmptyList[A]{Head: a, Tail: []A{}}
}

func Append[A any](first NonEmptyList[A], snd NonEmptyList[A]) NonEmptyList[A] {
	nel := One(first.Head)
	nel.Tail = append(nel.Tail, first.Tail...)
	nel.Tail = append(nel.Tail, snd.Head)
	nel.Tail = append(nel.Tail, snd.Tail...)

	return nel
}

func FlatMap[A any, B any](nel NonEmptyList[A], f func(A) NonEmptyList[B]) NonEmptyList[B] {
	result := f(nel.Head)
	for _, v := range nel.Tail {
		result = Append(result, f(v))
	}
	return result
}

func Map[A any, B any](nel NonEmptyList[A], f func(A) B) NonEmptyList[B] {
	fn := func(a A) NonEmptyList[B] {
		return One(f(a))
	}
	return FlatMap(nel, fn)
}

func Sort[A any](nel NonEmptyList[A], f func(A, A) (A, A)) NonEmptyList[A] {

	return nel
}

// func Filter[A any](slice []A, f func(A) bool) []A {
// 	fn := func(a A) []A {
// 		if f(a) {
// 			return []A{a}
// 		}
// 		return []A{}
// 	}
// 	return FlatMap(slice, fn)
// }

// func Exists[A any](slice []A, f func(A) bool) bool {
// 	for _, v := range slice {
// 		if f(v) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func FoldLeft[A any, B any](slice []A, initVal B, accumulator func(B, A) B) B {
// 	result := initVal
// 	for _, v := range slice {
// 		result = accumulator(result, v)
// 	}
// 	return result
// }
