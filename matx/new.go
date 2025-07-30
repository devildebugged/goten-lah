package matx

import "fmt"

func New(data []float64, dims []int) (*Matx, error) {
	if data == nil || dims == nil {
		return nil, fmt.Errorf("data or dimensions cannot be nil")
	}

	size := 1
	for _, d := range dims {
		size *= d
	}

	if size != len(data) {
		return nil, fmt.Errorf(
			"data size mismatch: expected %d elements for shape %v, but got %d",
			size, dims, len(data),
		)
	}

	return &Matx{
		Data:       data,
		Dimensions: dims,
	}, nil
}
