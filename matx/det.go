package matx

import (
	"fmt"
	"math"
)

func Det(m *Matx) (float64, error) {
	if m == nil {
		return 0, fmt.Errorf("Nil matrix passed")
	}
	if len(m.Dimensions) != 2 || m.Dimensions[0] != m.Dimensions[1] {
		return 0, fmt.Errorf("Matrix must be square")
	}

	_, U, _, swapCount, err := LUDecomposeWithPivoting(m)
	if err != nil {
		return 0, err
	}

	det := 1.0
	for i := 0; i < m.Dimensions[0]; i++ {
		diag := mustGet(U, i, i)
		if math.Abs(diag) < 1e-12 {
			return 0, nil
		}
		det *= diag
	}

	if swapCount%2 != 0 {
		det = -det
	}

	return det, nil
}
