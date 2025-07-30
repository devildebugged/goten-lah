package matx

import (
	"fmt"
	"math"
)

func LUDecomposeWithPivoting(orig *Matx) (*Matx, *Matx, []int, int, error) {
	n := orig.Dimensions[0]
	if n != orig.Dimensions[1] {
		return nil, nil, nil, 0, fmt.Errorf("Matrix must be square")
	}

	A := &Matx{
		Data:       append([]float64(nil), orig.Data...),
		Dimensions: []int{n, n},
	}

	L, _ := New(make([]float64, n*n), []int{n, n})
	U, _ := New(make([]float64, n*n), []int{n, n})
	pivots := make([]int, n)
	for i := range pivots {
		pivots[i] = i
	}

	swapCount := 0

	for i := 0; i < n; i++ {

		maxIdx := i
		maxVal := math.Abs(mustGet(A, i, i))
		for k := i + 1; k < n; k++ {
			if v := math.Abs(mustGet(A, k, i)); v > maxVal {
				maxVal = v
				maxIdx = k
			}
		}
		if maxVal == 0 {
			return nil, nil, nil, 0, fmt.Errorf("Matrix is singular")
		}

		if maxIdx != i {
			if err := RowSwap(A, i, maxIdx); err != nil {
				return nil, nil, nil, 0, err
			}
			pivots[i], pivots[maxIdx] = pivots[maxIdx], pivots[i]
			swapCount++
		}

		for j := i; j < n; j++ {
			sum := 0.0
			for k := 0; k < i; k++ {
				sum += mustGet(L, i, k) * mustGet(U, k, j)
			}
			mustSet(mustGet(A, i, j)-sum, U, i, j)
		}

		for j := i; j < n; j++ {
			if i == j {
				mustSet(1.0, L, i, i)
			} else {
				sum := 0.0
				for k := 0; k < i; k++ {
					sum += mustGet(L, j, k) * mustGet(U, k, i)
				}
				mustSet((mustGet(A, j, i)-sum)/mustGet(U, i, i), L, j, i)
			}
		}
	}

	return L, U, pivots, swapCount, nil
}
