package either

import (
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

//Exists, Fold, LeftMap, Map2
