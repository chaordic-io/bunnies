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
