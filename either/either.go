package either

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

func Left[A any, B any](a A) Either[A, B] {
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

// MapRight, MapLeft, GetOrElse, GetLeftOrElse
