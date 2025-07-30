package matx

import "fmt"

func Dot(m1 *Matx, m2 *Matx) (float64, error) {
	if m1 == nil || m2 == nil {
		return 0, fmt.Errorf("One or both matrices passed are nil")
	}

	if len(m1.Dimensions) != 1 || len(m2.Dimensions) != 1 {
		return 0, fmt.Errorf("Dot product is only for 1-D matrices")
	}

	if len(m1.Data) != len(m2.Data) {
		return 0, fmt.Errorf("Vectors must be of equal length")
	}

	sum := 0.0
	for i := 0; i < len(m1.Data); i++ {
		sum += m1.Data[i] * m2.Data[i]
	}

	return sum, nil
}
