package linkedlist

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"shenye.com/util"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestLinkedListAppend(t *testing.T) {
	l := New()
	order_l := New()

	length := util.RandomInt(0, 100)
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

func TestReverse(t *testing.T) {
	l := New()
	l.Append(114)
	l.Append(514)
	l.Append(12)
	l.Append(312313)
	fmt.Println(l)

	l.Reverse(4)

	fmt.Println(l)
}
