package either

// Should either be (A, B) aliased

type Either[A any, B any] interface {
	IsLeft() bool
	IsRight() bool
	left() A
	right() B
}

type leftOrRight[A any, B any] struct {
	l *A
	r *B
}

func ToEither[A any](fn func() (A, error)) Either[error, A] {
	res, err := fn()
	if err != nil {
		return Left[A](err)
	}
	return Right[error](res)
}

func (e *leftOrRight[A, B]) left() A {
	return *e.l
}

func (e *leftOrRight[A, B]) right() B {
	return *e.r
}

func (e *leftOrRight[A, B]) IsLeft() bool {
	return e.l != nil
}

func (e *leftOrRight[A, B]) IsRight() bool {
	return e.r != nil
}

func Left[B any, A any](a A) Either[A, B] {
	return &leftOrRight[A, B]{l: &a, r: nil}
}

func Right[A any, B any](b B) Either[A, B] {
	return &leftOrRight[A, B]{l: nil, r: &b}
}

func Fold[A any, B any, C any](either Either[A, B], left func(A) C, right func(B) C) C {
	if either.IsLeft() {
		return left(either.left())
	}
	return right(either.right())
}

func FlatMap[A any, B any, C any](e Either[A, B], f func(B) Either[A, C]) Either[A, C] {
	if e.IsLeft() {
		return Left[C](e.left())
	}
	return f(e.right())
}

func Map[A any, B any, C any](e Either[A, B], f func(B) C) Either[A, C] {
	fn := func(v B) Either[A, C] {
		return Right[A](f(v))
	}
	return FlatMap(e, fn)
}

func LeftMap[A any, B any, C any](e Either[A, B], f func(A) C) Either[C, B] {
	if e.IsRight() {
		return Right[C](e.right())
	}
	return Left[B](f(e.left()))
}

func GetOrElse[A any, B any](e Either[A, B], defaultVal B) B {
	if e.IsLeft() {
		return defaultVal
	}
	return e.right()
}

func GetLeftOrElse[A any, B any](e Either[A, B], defaultVal A) A {
	if e.IsLeft() {
		return e.left()
	}
	return defaultVal
}

func Exists[A any, B any](e Either[A, B], fn func(B) bool) bool {
	return e.IsRight() && fn(e.right())
}

// Map2, FlatMap2
