package benchmarks

import "testing"

var data []int
var size = 100_000_000

func init() {
	for i := 0; i < size; i++ {
		data = append(data, i)
	}
}

func BenchmarkCommonWrite(t *testing.B) {
	res := []int{}
	res = append(res, data...)
}

func BenchmarkReservedWrite(t *testing.B) {
	res := make([]int, size)
	copy(res, data)
}

func BenchmarkBufferWrite(t *testing.B) {
	res := make([]int, size)
	buf_size := 10000
	buf := make([]int, buf_size)

	for i, d := range data {
		buf = append(buf, d)
		if i%buf_size == 0 {
			res = append(res, buf...)
			buf = make([]int, buf_size)
		}
	}
}
