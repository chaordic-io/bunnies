package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	assert.Equal(t, FlatMap([]int{1, 2, 3}, func(i int) []int {
		return []int{i, i + 1}
	}), []int{1, 2, 2, 3, 3, 4})
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map([]int{1, 2, 3}, func(i int) int {
		return i + 1
	}), []int{2, 3, 4})
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Filter([]int{1, 2, 3}, func(i int) bool {
		return i%2 == 0
	}), []int{2})
}

func TestExists(t *testing.T) {
	assert.Equal(t, Exists([]int{1, 2, 3}, func(i int) bool {
		return i%2 == 0
	}), true)

	assert.Equal(t, Exists([]int{1, 2, 3}, func(i int) bool {
		return i%5 == 0
	}), false)
}

func TestFoldLeft(t *testing.T) {
	assert.Equal(t, FoldLeft([]int{1, 2, 3}, 0, func(i, j int) int {
		return i + j
	}), 6)
}

func TestMap2(t *testing.T) {

	assert.Equal(t, Map2([]string{"ha", "heh", "hmm"}, []string{"?", "!", "."}, func(a, b string) string {
		return a + b
	}), []string{"ha?", "ha!", "ha.", "heh?", "heh!", "heh.", "hmm?", "hmm!", "hmm."})

}

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty([]string{}))
	assert.False(t, IsEmpty([]string{"foo"}))
}
