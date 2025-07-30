package matx

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
)

type testMeta struct {
	num   int
	name  string
	start time.Time
	t     *testing.T
}

func begin(t *testing.T, num int, name string) testMeta {
	meta := testMeta{num, name, time.Now(), t}
	fmt.Printf("Test %02d: %-30s", num, name)
	return meta
}

func (meta testMeta) end(ok bool) {
	duration := time.Since(meta.start)
	status := "[FAIL]"
	if ok {
		status = "[PASS]"
	}
	fmt.Printf("%s\t(%s)\n", status, duration)
	if !ok {
		meta.t.Fail()
	}
}

func TestMatx(t *testing.T) {
	n := 1

	{ // New()
		m := begin(t, n, "New() with valid input")
		n++
		mat, err := New([]float64{1, 2, 3, 4}, []int{2, 2})
		ok := err == nil && reflect.DeepEqual(mat.Data, []float64{1, 2, 3, 4})
		m.end(ok)
	}

	{ // New with bad dims
		m := begin(t, n, "New() with size mismatch")
		n++
		_, err := New([]float64{1, 2, 3}, []int{2, 2})
		m.end(err != nil)
	}

	{ // Multiply
		m := begin(t, n, "Multiply() 2x2 matrices")
		n++
		a, _ := New([]float64{1, 2, 3, 4}, []int{2, 2})
		b, _ := New([]float64{5, 6, 7, 8}, []int{2, 2})
		c, err := Multiply(a, b)
		exp := []float64{19, 22, 43, 50}
		m.end(err == nil && reflect.DeepEqual(c.Data, exp))
	}

	{ // Ones
		m := begin(t, n, "Ones() 2x3 matrix")
		n++
		mat, err := Ones([]int{2, 3})
		allOnes := true
		for _, v := range mat.Data {
			if v != 1 {
				allOnes = false
				break
			}
		}
		m.end(err == nil && allOnes)
	}

	{ // Det
		m := begin(t, n, "Det() of 2x2 matrix")
		n++
		mat, _ := New([]float64{4, 3, 6, 3}, []int{2, 2})
		d, err := Det(mat)
		m.end(err == nil && math.Abs(d+6) < 1e-9)
	}

	{ // IsInvertible
		m := begin(t, n, "IsInvertible() true")
		n++
		mat, _ := New([]float64{4, 3, 6, 3}, []int{2, 2})
		res, err := IsInvertible(mat)
		m.end(err == nil && res)
	}

	{ // Invert
		m := begin(t, n, "Invert() and validate I")
		n++
		mat, _ := New([]float64{4, 3, 6, 3}, []int{2, 2})
		inv, err := Invert(mat)
		if err != nil {
			m.end(false)
			return
		}
		id, err := Multiply(mat, inv)
		if err != nil {
			m.end(false)
			return
		}
		ok := true
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				v := id.Data[i*2+j]
				if i == j && math.Abs(v-1) > 1e-6 {
					ok = false
				} else if i != j && math.Abs(v) > 1e-6 {
					ok = false
				}
			}
		}
		m.end(ok)
	}

	{ // Dot
		m := begin(t, n, "Dot() of 1D vectors")
		n++
		a, _ := New([]float64{1, 2, 3}, []int{3})
		b, _ := New([]float64{4, 5, 6}, []int{3})
		res, err := Dot(a, b)
		m.end(err == nil && res == 32)
	}
}
