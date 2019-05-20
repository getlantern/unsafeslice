package unsafeslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftHeadLeftInt(t *testing.T) {
	orig := []int{1, 2, 3, 4}
	s := orig[1:]
	ShiftHeadLeft(&s, 1)
	assert.EqualValues(t, s, orig)
	assert.Equal(t, cap(orig), cap(s))
	assert.Equal(t, len(orig), len(s))
}

func TestShiftHeadLeftInterface(t *testing.T) {
	orig := []interface{}{"1", "2", "3", "4"}
	orig = orig[1:3]
	s := orig[1:]
	ShiftHeadLeft(&s, 1)
	assert.EqualValues(t, s, orig)
	assert.Equal(t, cap(orig), cap(s))
	assert.Equal(t, len(orig), len(s))
}
