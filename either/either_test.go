package either

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftRight(t *testing.T) {
	assert.Equal(t, GetLeftOrElse(Left[int](true), false), true)
	assert.Equal(t, GetOrElse(Left[int](true), 1), 1)
	assert.Equal(t, GetOrElse(Right[bool](2), 1), 2)
	assert.Equal(t, GetLeftOrElse(Right[bool](2), false), false)

}

func TestMap(t *testing.T) {
	fn := func(i int) int {
		return i + 1
	}
	assert.Equal(t, Map(Left[int](false), fn), Left[int](false))
	assert.Equal(t, Map(Right[bool](2), fn), Right[bool](3))
}

func TestToEither(t *testing.T) {
	res := ToEither(func() (int, error) {
		return -1, fmt.Errorf("error!")
	})
	assert.Equal(t, Left[int](fmt.Errorf("error!")), res)

	res = ToEither(func() (int, error) {
		return 1, nil
	})
	assert.Equal(t, Right[error](1), res)
}

func TestExists(t *testing.T) {
	fn := func(i int) bool {
		return i == 4
	}
	assert.False(t, Exists(Left[int](4), fn))
	assert.False(t, Exists(Right[int](3), fn))
	assert.True(t, Exists(Right[int](4), fn))
}

func TestFold(t *testing.T) {
	fns := func(s string) int {
		return 0
	}
	fni := func(i int) int {
		return i + 1
	}

	assert.Equal(t, 0, Fold(Left[int]("hello"), fns, fni))
	assert.Equal(t, 2, Fold(Right[string](1), fns, fni))
}

func TestLeftMap(t *testing.T) {
	fn := func(i int) int {
		return i + 2
	}

	assert.Equal(t, Left[string](5), LeftMap(Left[string](3), fn))

	assert.Equal(t, Right[int]("hello"), LeftMap(Right[int]("hello"), fn))

}

// Map2
