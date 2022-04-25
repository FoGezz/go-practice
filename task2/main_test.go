package task2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	case1, _ := UnpackString("a4bc2d5e")
	assert.Equal(t, "aaaabccddddde", case1)

	case2, _ := UnpackString("abcd")
	assert.Equal(t, "abcd", case2)

	case3, _ := UnpackString("a4bc2d5e")
	assert.Equal(t, "aaaabccddddde", case3)

	_, err := UnpackString("45")
	assert.NotNil(t, err)

	_, err = UnpackString("g0")
	assert.NotNil(t, err)
}

func TestUnpackStringWithEscape(t *testing.T) {
	case1, _ := UnpackString("qwe\\4\\5")
	assert.Equal(t, "qwe45", case1)

	case2, _ := UnpackString("qwe\\\\5")
	assert.Equal(t, "qwe\\\\\\\\\\", case2)
}
