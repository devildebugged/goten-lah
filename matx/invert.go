package matx

import "fmt"

func Invert(m *Matx) (*Matx, error) {
	if m == nil {
		return nil, fmt.Errorf("Nil matrix")
	}

	n := m.Dimensions[0]
	if len(m.Dimensions) != 2 || n != m.Dimensions[1] {
		return nil, fmt.Errorf("Matrix must be square")
	}

	L, U, pivots, _, err := LUDecomposeWithPivoting(m)
	if err != nil {
		return nil, err
	}

	inv, _ := New(make([]float64, n*n), []int{n, n})
	e := make([]float64, n)
	y := make([]float64, n)
	x := make([]float64, n)

	for col := 0; col < n; col++ {
		// zero out slices
		for i := 0; i < n; i++ {
			e[i], y[i], x[i] = 0, 0, 0
		}
		e[pivots[col]] = 1.0

		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < i; j++ {
				sum += mustGet(L, i, j) * y[j]
			}
			y[i] = e[i] - sum
		}

		for i := n - 1; i >= 0; i-- {
			sum := 0.0
			for j := i + 1; j < n; j++ {
				sum += mustGet(U, i, j) * x[j]
			}
			x[i] = (y[i] - sum) / mustGet(U, i, i)
		}

		for row := 0; row < n; row++ {
			mustSet(x[row], inv, row, col)
		}
	}

	return inv, nil
}
