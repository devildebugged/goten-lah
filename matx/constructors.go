package matx

import (
	"fmt"
	"math/rand"
	"time"
)

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

func Rand(rows, cols int, bounds ...float64) (*Matx, error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("invalid matrix size: %dx%d", rows, cols)
	}

	min, max := 0.0, 1.0
	if len(bounds) == 2 {
		min, max = bounds[0], bounds[1]
		if min >= max {
			return nil, fmt.Errorf("invalid range: min (%v) must be less than max (%v)", min, max)
		}
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	data := make([]float64, rows*cols)
	for i := range data {
		data[i] = rng.Float64()*(max-min) + min
	}

	return &Matx{
		Data:       data,
		Dimensions: []int{rows, cols},
	}, nil
}

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
