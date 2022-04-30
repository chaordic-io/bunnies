package nonempty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	fn := func(i int) NonEmptyList[int] {
		return Nel(i, []int{i + 1})
	}
	result := FlatMap(buildNel(), fn)
	assert.Equal(t, result, Nel(1, []int{2, 2, 3, 3, 4}))

}

func TestMap(t *testing.T) {
	fn := func(i int) int {
		return i + 1
	}
	result := Map(buildNel(), fn)
	assert.Equal(t, result, Nel(2, []int{3, 4}))

}

func TestFilter(t *testing.T) {
	fn := func(i int) bool {
		return i%2 == 0
	}
	result := Filter(buildNel(), fn)
	assert.Equal(t, result, []int{2})
}

func TestExists(t *testing.T) {
	fn := func(i int) bool {
		return i%2 == 0
	}
	assert.True(t, Exists(buildNel(), fn))

	fn2 := func(i int) bool {
		return i%5 == 0
	}
	assert.False(t, Exists(buildNel(), fn2))

	fn3 := func(i int) bool {
		return i == 1
	}
	assert.True(t, Exists(buildNel(), fn3))
}

func TestFoldLeft(t *testing.T) {
	fn := func(i, j int) int {
		return i + j
	}
	assert.Equal(t, FoldLeft(buildNel(), 0, fn), 6)
}

func TestMap2(t *testing.T) {
	assert.Equal(t, ToSlice(Map2(Nel("ha", []string{"heh", "hmm"}), Nel("?", []string{"!", "."}), func(a, b string) string {
		return a + b
	})), []string{"ha?", "ha!", "ha.", "heh?", "heh!", "heh.", "hmm?", "hmm!", "hmm."})
}

func buildNel() NonEmptyList[int] {
	return Nel(1, []int{2, 3})
}

func TestAppendOne(t *testing.T) {
	assert.Equal(t, ToSlice(AppendOne(AppendOne(One(1), 2), 3)), []int{1, 2, 3})
}
