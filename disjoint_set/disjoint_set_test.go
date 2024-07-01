package disjointset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 1-0-2-3 4 5
func TestBasic(t *testing.T) {
	ds := New(5)

	ds.UnionByRank(0, 1)
	ds.UnionByRank(2, 3)
	ds.UnionByRank(0, 2)

	require.Equal(t, ds.Find(0), 0)
	require.Equal(t, ds.Find(3), 0)

	require.Equal(t, ds.rank, []int{2, 0, 1, 0, 0})
}

// 1-0-2-3 4 5
func TestBasicSize(t *testing.T) {
	ds := New(5)

	ds.UnionBySize(0, 1)
	ds.UnionBySize(2, 3)
	ds.UnionBySize(0, 2)

	require.Equal(t, ds.Find(0), 0)
	require.Equal(t, ds.Find(3), 0)

	require.Equal(t, ds.size, []int{4, 1, 2, 1, 1})
}
