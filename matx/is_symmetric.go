package matx

import (
	"fmt"
)

func IsSymmetric(m *Matx) (bool, error) {
	if m == nil {
		return false, fmt.Errorf("Nil matrix passed")
	}
	if len(m.Dimensions) != 2 {
		return false, fmt.Errorf("Only 2D matrices supported")
	}
	if m.Dimensions[0] != m.Dimensions[1] {
		return false, nil
	}

	n := m.Dimensions[0]

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			v1, err1 := Get(m, i, j)
			v2, err2 := Get(m, j, i)
			if err1 != nil || err2 != nil {
				return false, fmt.Errorf("Index out of bounds during symmetry check")
			}
			if v1 != v2 {
				return false, nil
			}
		}
	}

	return true, nil
}
