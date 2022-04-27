package option

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	assert.Equal(t, FlatMap(Some(1), func(i int) Option[string] {
		return Some(fmt.Sprintf("no-%v", i))
	}), Some("no-1"))

	assert.Equal(t, FlatMap(Some(1), func(i int) Option[string] {
		return None[string]()
	}), None[string]())

	assert.Equal(t, FlatMap(None[int](), func(i int) Option[string] {
		return Some(fmt.Sprintf("no-%v", i))
	}), None[string]())
}

func TestMap(t *testing.T) {
	mapFn := func(i int) string {
		return fmt.Sprintf("no-%v", i)
	}
	assert.Equal(t, Map(Some(1), mapFn), Some("no-1"))

	assert.Equal(t, Map(None[int](), mapFn), None[string]())
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Filter(Some(1), func(i int) bool {
		return i == 1
	}), Some(1))

	assert.Equal(t, Filter(Some(2), func(i int) bool {
		return i == 1
	}), None[int]())
}

func TestExists(t *testing.T) {
	assert.True(t, Exists(Some(1), func(i int) bool {
		return i == 1
	}))

	assert.False(t, Exists(Some(2), func(i int) bool {
		return i == 1
	}))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, None[int]().IsEmpty())
	assert.False(t, Some(1).IsEmpty())
}

func TestGetOrElse(t *testing.T) {
	assert.Equal(t, None[int]().GetOrElse(2), 2)
	assert.Equal(t, Some(1).GetOrElse(2), 1)
}
