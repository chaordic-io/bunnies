package option

// option is the package private struct that implements Optional
type option[A any] struct {
	theValue *A
}

// Optional is a type that may or may not contain a value and allows the user to safely operate on this assumption.
type Optional[A any] interface {
	GetOrElse(defaultVal A) A
	//IsEmpty tells if this optional is empty or contains a value
	IsEmpty() bool
	value() A
}

//IsEmpty tells if this optional is empty or contains a value
func (opt *option[A]) IsEmpty() bool {
	return opt.theValue == nil
}

func (opt *option[A]) GetOrElse(defaultVal A) A {
	if opt.IsEmpty() {
		return defaultVal
	}
	return opt.value()
}

//value is a package private method that takes the value out of the optional unsafely. Used for internal implementation.
func (opt *option[A]) value() A {
	return *opt.theValue
}

//Some is the Optional implementation of `Pure`, it constructs an Optional containing a value.
func Some[A any](val A) Optional[A] {
	return &option[A]{&val}
}

//Empty constructs an Optional that contains no value
func Empty[A any]() Optional[A] {
	return &option[A]{nil}
}

func Map[A any, B any](opt Optional[A], fn func(A) B) Optional[B] {
	if opt.IsEmpty() {
		return Empty[B]()
	}
	bVal := fn(opt.value())
	return Some(bVal)
}

func FlatMap[A any, B any](opt Optional[A], fn func(A) Optional[B]) Optional[B] {
	if opt.IsEmpty() {
		return Empty[B]()
	}
	return fn(opt.value())
}

func Filter[A any](opt Optional[A], fn func(A) bool) Optional[A] {
	if opt.IsEmpty() || !fn(opt.value()) {
		return Empty[A]()
	}
	return opt
}

func Exists[A any](opt Optional[A], fn func(A) bool) bool {
	return !Filter(opt, fn).IsEmpty()
}
