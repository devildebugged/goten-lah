package matx

import "fmt"

func Identity(dimensions ...int) (*Matx, error) {
	if dimensions == nil || len(dimensions) != 2 || dimensions[0] != dimensions[1] {
		return nil, fmt.Errorf("identity matrix must be square (got: %v)", dimensions)
	}

	size := dimensions[0]
	data := make([]float64, size*size)

	for i := 0; i < size; i++ {
		data[i*size+i] = 1
	}

	mat, err := New(data, dimensions)
	if err != nil {
		return nil, fmt.Errorf("failed to create identity matrix: %w", err)
	}

	return mat, nil
}
