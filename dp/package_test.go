package dp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZeroOnePackage(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	r := ZeroOnePackage(weight, value, 4)
	require.Equal(t, r, 35)

	r = ZeroOnePackagePro(weight, value, 4)
	require.Equal(t, r, 35)
}
