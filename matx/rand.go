package matx

import (
	"fmt"
	"math/rand"
	"time"
)

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
