package option

type optional[A any] struct {
	theValue *A
}

type Option[A any] interface {
	GetOrElse(defaultVal A) A
	IsEmpty() bool
	value() A
}

func (opt *optional[A]) IsEmpty() bool {
	return opt.theValue == nil
}

func (opt *optional[A]) GetOrElse(defaultVal A) A {
	if opt.IsEmpty() {
		return defaultVal
	}
	return opt.value()
}
func (opt *optional[A]) value() A {
	return *opt.theValue
}

func Some[A any](val A) Option[A] {
	return &optional[A]{&val}
}

func None[A any]() Option[A] {
	return &optional[A]{nil}
}

func Map[A any, B any](opt Option[A], fn func(A) B) Option[B] {
	if opt.IsEmpty() {
		return None[B]()
	}
	bVal := fn(opt.value())
	return Some(bVal)
}

func FlatMap[A any, B any](opt Option[A], fn func(A) Option[B]) Option[B] {
	if opt.IsEmpty() {
		return None[B]()
	}
	return fn(opt.value())
}

func Filter[A any](opt Option[A], fn func(A) bool) Option[A] {
	if opt.IsEmpty() || !fn(opt.value()) {
		return None[A]()
	}
	return opt
}

func Exists[A any](opt Option[A], fn func(A) bool) bool {
	return !Filter(opt, fn).IsEmpty()
}
