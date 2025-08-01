package matx

import "testing"

func BenchmarkMultiply2x2(b *testing.B) {
	a, _ := New([]float64{1, 2, 3, 4}, []int{2, 2})
	c, _ := New([]float64{5, 6, 7, 8}, []int{2, 2})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Multiply(a, c)
	}
}

func BenchmarkMultiply100x100(b *testing.B) {
	size := 100
	data := make([]float64, size*size)
	for i := range data {
		data[i] = 1
	}
	a, _ := New(data, []int{size, size})
	c, _ := New(data, []int{size, size})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Multiply(a, c)
	}
}

func BenchmarkInvert10x10(b *testing.B) {
	size := 10
	data := make([]float64, size*size)
	for i := range data {
		data[i] = float64((i % size) + 1)
	}
	mat, _ := New(data, []int{size, size})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Invert(mat)
	}
}

func BenchmarkDot1000(b *testing.B) {
	size := 1000
	data := make([]float64, size)
	for i := range data {
		data[i] = float64(i)
	}
	a, _ := New(data, []int{size})
	bb, _ := New(data, []int{size})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Dot(a, bb)
	}
}
