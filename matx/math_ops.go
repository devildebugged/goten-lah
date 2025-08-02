package matx

import (
	"fmt"
	"math"
)

// Add returns a new matrix that is the element-wise sum of matrices `m1` and `m2`.
// Both input matrices must be non-nil and have identical dimensions and sizes.
// Returns an error if validation fails.
func Add(m1, m2 *Matx) (*Matx, error) {
	if m1 == nil || m2 == nil {
		return nil, fmt.Errorf("one or both the matrices are nil")
	}
	if !CheckDimensionEquality(m1.Dimensions, m2.Dimensions) {
		return nil, fmt.Errorf("dimension mismatch: %v vs %v", m1.Dimensions, m2.Dimensions)
	}

	m1Size, err := Size(m1)
	if err != nil {
		return nil, fmt.Errorf("failed to compute size of m1: %w", err)
	}
	m2Size, err := Size(m2)
	if err != nil {
		return nil, fmt.Errorf("failed to compute size of m2: %w", err)
	}
	if m1Size != m2Size {
		return nil, fmt.Errorf("data size mismatch: %d vs %d", m1Size, m2Size)
	}

	resultData := make([]float64, m1Size)
	for i := 0; i < m1Size; i++ {
		resultData[i] = m1.Data[i] + m2.Data[i]
	}

	resultMatx, err := New(resultData, m1.Dimensions)
	if err != nil {
		return nil, fmt.Errorf("failed to create result matrix: %w", err)
	}
	return resultMatx, nil
}

// Negate performs an in-place negation of all elements in the matrix.
func (m *Matx) Negate() error {
	if m == nil || m.Data == nil {
		return fmt.Errorf("cannot negate: matrix is nil or uninitialized")
	}

	for i := range m.Data {
		m.Data[i] *= -1
	}
	return nil
}

// Scale multiplies all elements of the matrix by scalar `n`.
func (m *Matx) Scale(n int) error {
	if m == nil || m.Data == nil {
		return fmt.Errorf("nil matrix given")
	}

	for i := range m.Data {
		m.Data[i] *= float64(n)
	}
	return nil
}

// Raise raises each element of the matrix to the specified `power`.
func (m *Matx) Raise(power float64) error {
	if m == nil || m.Data == nil {
		return fmt.Errorf("cannot raise: matrix is nil or uninitialized")
	}

	for i := range m.Data {
		m.Data[i] = math.Pow(m.Data[i], power)
	}
	return nil
}

// Reciprocal transforms each element of the matrix to its multiplicative inverse (1/x).
func (m *Matx) Reciprocal() error {
	if m == nil || m.Data == nil {
		return fmt.Errorf("cannot reciprocate: matrix or matrix data is nil")
	}

	for i := range m.Data {
		if m.Data[i] == 0 {
			return fmt.Errorf("cannot reciprocate: division by zero at index %d", i)
		}
		m.Data[i] = 1 / m.Data[i]
	}
	return nil
}
