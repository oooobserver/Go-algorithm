package structure

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestLinkedListAppend(t *testing.T) {
	l := LinkedList{}
	order_l := LinkedList{}

	length := rng.Intn(100)
	record := make([]int, 0)

	for i := 0; i < length; i++ {
		tmp := rng.Int()
		l.Append(tmp)
		order_l.OrderAppend(tmp)
		record = append(record, tmp)
	}

	array := order_l.LL2Array()
	require.Equal(t, record, array)
	for i, j := 0, len(record)-1; i < j; i, j = i+1, j-1 {
		record[i], record[j] = record[j], record[i]
	}
	array = l.LL2Array()
	require.Equal(t, record, array)
}
