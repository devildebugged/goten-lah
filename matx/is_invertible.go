package matx

import (
	"fmt"
	"math"
)

func IsInvertible(m *Matx) (bool, error) {
	if m == nil {
		return false, fmt.Errorf("Nil matrix passed")
	}

	det, err := Det(m)
	if err != nil {
		return false, err
	}
	return math.Abs(det) > 1e-12, nil
}
