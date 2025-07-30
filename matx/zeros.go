package matx

import "fmt"

func Zeros(dimensions []int) (*Matx, error) {
	if len(dimensions) == 0 {
		return nil, fmt.Errorf("zeros matrix creation failed: dimensions cannot be empty")
	}

	dataSize := 1
	for _, dim := range dimensions {
		dataSize *= dim
	}

	mat, err := New(make([]float64, dataSize), dimensions)
	if err != nil {
		return nil, fmt.Errorf("failed to create zeros matrix: %w", err)
	}

	return mat, nil
}
