package rope

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestBasic(t *testing.T) {
	s1 := "The quick brown fox jumped over"
	r1 := FromString(s1)
	assert.Equal(t, s1, r1.ToString())

}
