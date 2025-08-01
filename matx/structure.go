package matx

import "fmt"

type Matx struct {
	Data       []float64
	Dimensions []int
}

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

func CheckDimensionEquality(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func CheckMultiplicationCondition(a, b []int) bool {
	return a[1] == b[0]
}
