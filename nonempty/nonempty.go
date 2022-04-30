package nonempty

import "github.com/chaordic-io/bunnies/list"

type OneAnd[A any, B any] struct {
	Head A
	Tail B
}

type NonEmptyList[A any] OneAnd[A, []A]

func One[A any](a A) NonEmptyList[A] {
	return NonEmptyList[A]{Head: a, Tail: []A{}}
}

func Nel[A any](a A, t []A) NonEmptyList[A] {
	return NonEmptyList[A]{Head: a, Tail: t}
}

func Append[A any](first NonEmptyList[A], snd NonEmptyList[A]) NonEmptyList[A] {
	nel := One(first.Head)
	nel.Tail = append(nel.Tail, first.Tail...)
	nel.Tail = append(nel.Tail, snd.Head)
	nel.Tail = append(nel.Tail, snd.Tail...)

	return nel
}

func ToSlice[A any](nel NonEmptyList[A]) []A {
	result := []A{nel.Head}
	return append(result, nel.Tail...)
}

func AppendOne[A any](first NonEmptyList[A], elem A) NonEmptyList[A] {
	nel := One(first.Head)
	nel.Tail = append(first.Tail, elem)
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

func Filter[A any](nel NonEmptyList[A], f func(A) bool) []A {
	fn := func(a A) []A {
		if f(a) {
			return []A{a}
		}
		return []A{}
	}
	in := []A{nel.Head}
	return list.FlatMap(append(in, nel.Tail...), fn)
}

func Exists[A any](nel NonEmptyList[A], f func(A) bool) bool {
	if f(nel.Head) {
		return true
	}
	for _, v := range nel.Tail {
		if f(v) {
			return true
		}
	}
	return false
}

func FoldLeft[A any, B any](nel NonEmptyList[A], initVal B, accumulator func(B, A) B) B {
	result := accumulator(initVal, nel.Head)
	for _, v := range nel.Tail {
		result = accumulator(result, v)
	}
	return result
}

func Map2[A any, B any, C any](slice1 NonEmptyList[A], slice2 NonEmptyList[B], f func(A, B) C) NonEmptyList[C] {
	return FlatMap2(slice1, slice2, func(a A, b B) NonEmptyList[C] {
		return One(f(a, b))
	})
}

func FlatMap2[A any, B any, C any](slice1 NonEmptyList[A], slice2 NonEmptyList[B], f func(A, B) NonEmptyList[C]) NonEmptyList[C] {
	nel := f(slice1.Head, slice2.Head)
	for _, v := range slice2.Tail {
		nel = Append(nel, f(slice1.Head, v))
	}

	for _, v1 := range slice1.Tail {
		nel = Append(nel, f(v1, slice2.Head))
		for _, v2 := range slice2.Tail {
			nel = Append(nel, f(v1, v2))
		}
	}

	return nel
}
