package matx

import "fmt"

func Ones(dimensions []int) (*Matx, error) {
	if len(dimensions) != 2 {
		return nil, fmt.Errorf("ones matrix must be 2D, got %dD", len(dimensions))
	}

	dataSize := dimensions[0] * dimensions[1]
	data := make([]float64, dataSize)
	for i := range data {
		data[i] = 1
	}

	mat, err := New(data, dimensions)
	if err != nil {
		return nil, fmt.Errorf("failed to create ones matrix: %w", err)
	}

	return mat, nil
}
