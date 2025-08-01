package matx

import (
	"fmt"
	"math"
)

func Det(m *Matx) (float64, error) {
	if m == nil {
		return 0, fmt.Errorf("Nil matrix passed")
	}
	if len(m.Dimensions) != 2 || m.Dimensions[0] != m.Dimensions[1] {
		return 0, fmt.Errorf("Matrix must be square")
	}

	_, U, _, swapCount, err := LUDecomposeWithPivoting(m)
	if err != nil {
		return 0, err
	}

	det := 1.0
	for i := 0; i < m.Dimensions[0]; i++ {
		diag := mustGet(U, i, i)
		if math.Abs(diag) < 1e-12 {
			return 0, nil
		}
		det *= diag
	}

	if swapCount%2 != 0 {
		det = -det
	}

	return det, nil
}

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

func IsInvertible(m *Matx) (bool, error) {
	if m == nil {
		return false, fmt.Errorf("Nil matrix passed")
	}

	det, err := Det(m)
	if err != nil {
		return false, err
	}
	return math.Abs(det) > 1e-12, nil
}

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

func IsSymmetric(m *Matx) (bool, error) {
	if m == nil {
		return false, fmt.Errorf("Nil matrix passed")
	}
	if len(m.Dimensions) != 2 {
		return false, fmt.Errorf("Only 2D matrices supported")
	}
	if m.Dimensions[0] != m.Dimensions[1] {
		return false, nil
	}

	n := m.Dimensions[0]

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v1, err1 := Get(m, i, j)
			v2, err2 := Get(m, j, i)
			if err1 != nil || err2 != nil {
				return false, fmt.Errorf("Index out of bounds during symmetry check")
			}
			if v1 != v2 {
				return false, nil
			}
		}
	}

	return true, nil
}

func Dot(m1 *Matx, m2 *Matx) (float64, error) {
	if m1 == nil || m2 == nil {
		return 0, fmt.Errorf("One or both matrices passed are nil")
	}

	if len(m1.Dimensions) != 1 || len(m2.Dimensions) != 1 {
		return 0, fmt.Errorf("Dot product is only for 1-D matrices")
	}

	if len(m1.Data) != len(m2.Data) {
		return 0, fmt.Errorf("Vectors must be of equal length")
	}

	sum := 0.0
	for i := 0; i < len(m1.Data); i++ {
		sum += m1.Data[i] * m2.Data[i]
	}

	return sum, nil
}

func Transpose(m *Matx) (*Matx, error) {
	if m == nil || m.Data == nil || len(m.Dimensions) != 2 {
		return nil, fmt.Errorf("invalid matrix for transpose")
	}

	rows, cols := m.Dimensions[0], m.Dimensions[1]
	result := make([]float64, len(m.Data))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j*rows+i] = m.Data[i*cols+j]
		}
	}

	newDims := []int{cols, rows}
	mat, err := New(result, newDims)
	if err != nil {
		return nil, fmt.Errorf("transpose error: %w", err)
	}

	return mat, nil
}

func RowSwap(m *Matx, row1, row2 int) error {
	if m == nil {
		return fmt.Errorf("Matrix is nil")
	}
	if len(m.Dimensions) != 2 {
		return fmt.Errorf("Matrix must be 2D")
	}

	rows, cols := m.Dimensions[0], m.Dimensions[1]

	if row1 < 0 || row1 >= rows || row2 < 0 || row2 >= rows {
		return fmt.Errorf("Row indices out of bounds")
	}

	for j := 0; j < cols; j++ {
		i1 := row1*cols + j
		i2 := row2*cols + j
		m.Data[i1], m.Data[i2] = m.Data[i2], m.Data[i1]
	}

	return nil
}

func Multiply(m1, m2 *Matx) (*Matx, error) {
	if m1 == nil || m2 == nil {
		return nil, fmt.Errorf("one or both input matrices are nil")
	}

	if !CheckMultiplicationCondition(m1.Dimensions, m2.Dimensions) {
		return nil, fmt.Errorf(
			"multiplication not possible: m1 columns (%d) != m2 rows (%d)",
			m1.Dimensions[1], m2.Dimensions[0],
		)
	}

	resultRows := m1.Dimensions[0]
	resultCols := m2.Dimensions[1]
	resultData := make([]float64, resultRows*resultCols)

	result, err := New(resultData, []int{resultRows, resultCols})
	if err != nil {
		return nil, fmt.Errorf("failed to create result matrix: %w", err)
	}

	for i := 0; i < resultRows; i++ {
		for j := 0; j < resultCols; j++ {
			sum := 0.0
			for k := 0; k < m1.Dimensions[1]; k++ {

				a := m1.Data[i*m1.Dimensions[1]+k]
				b := m2.Data[k*m2.Dimensions[1]+j]
				sum += a * b
			}
			result.Data[i*resultCols+j] = sum
		}
	}

	return result, nil
}
