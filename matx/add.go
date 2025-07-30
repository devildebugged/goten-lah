package matx

import (
	"fmt"
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
