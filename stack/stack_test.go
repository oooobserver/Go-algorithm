package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
	"shenye.com/util"
)

func TestStack(t *testing.T) {
	s := New()

	items := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		if i < 50 {
			items[i] = util.JustRandomInt()
		} else {
			items[i] = util.RandomString(10)
		}
		s.Push_back(items[i])
	}

	l := s.Len()
	require.Equal(t, 100, l)

	for i := 99; i >= 0; i-- {
		require.Equal(t, items[i], s.Pop_off())
	}
}
