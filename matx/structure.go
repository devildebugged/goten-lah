package matx

import "fmt"

// Matx represents a multi-dimensional matrix.
// - Data contains the flattened elements stored in row-major order.
// - Dimensions defines the shape of the matrix along each axis.
type Matx struct {
	Data       []float64 // Flattened data array representing the matrix contents.
	Dimensions []int     // Size of the matrix along each dimension.
}

// Size computes the total number of elements in the matrix by taking the
// product of all dimensions. Returns an error if the matrix is nil.
// If Dimensions is empty, the function returns 0.
func Size(m *Matx) (int, error) {
	if m == nil {
		return 0, fmt.Errorf("matx is nil")
	}

	if len(m.Dimensions) == 0 {
		return 0, nil
	}

	prod := 1
	for _, v := range m.Dimensions {
		prod *= v
	}

	return prod, nil
}

// CheckDimensionEquality returns true if two slices representing matrix dimensions
// are of equal length and all corresponding dimension sizes match.
func CheckDimensionEquality(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CheckMultiplicationCondition returns true if two matrices with the given shapes
// (represented by dimension slices) can be legally multiplied.
// Assumes both matrices are 2-dimensional.
// Matrix A must have the same number of columns as the number of rows in matrix B.
func CheckMultiplicationCondition(a, b []int) bool {
	return a[1] == b[0]
}
