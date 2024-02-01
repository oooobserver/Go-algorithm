package structure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	h := NewHash()
	for i := 0; i < 100; i++ {
		h.Hash_set(i, i+1)
	}

	for i := 0; i < 100; i++ {
		temp := h.Hash_get(i)
		require.Equal(t, temp, i+1)
	}
}
