package matx

import "fmt"

func Size(m *Matx) (int, error) {
	if m == nil {
		return 0, fmt.Errorf("matx is nil")
	}

	if len(m.Dimensions) == 0 {
		return 0, nil
	}

	prod := 1
	for _, v := range m.Dimensions {
		prod *= v
	}

	return prod, nil
}
