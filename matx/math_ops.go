package matx

import (
	"fmt"
	"math"
)

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

func Negate(m *Matx) (*Matx, error) {
	if m == nil || m.Data == nil {
		return nil, fmt.Errorf("cannot negate: matrix is nil or uninitialized")
	}

	for i := range m.Data {
		m.Data[i] *= -1
	}

	return m, nil
}

func Scale(m *Matx, n int) (*Matx, error) {
	if m == nil {
		return nil, fmt.Errorf("Nil matric given")
	}

	for i := 0; i < len(m.Data); i++ {
		m.Data[i] *= float64(n)
	}

	return m, nil
}

func Raise(m *Matx, power float64) (*Matx, error) {
	if m == nil || m.Data == nil {
		return nil, fmt.Errorf("cannot raise: matrix is nil or uninitialized")
	}

	for i := range m.Data {
		m.Data[i] = math.Pow(m.Data[i], power)
	}

	return m, nil
}
