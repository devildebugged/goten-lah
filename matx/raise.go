package matx

import (
	"fmt"
	"math"
)

func Raise(m *Matx, power float64) (*Matx, error) {
	if m == nil || m.Data == nil {
		return nil, fmt.Errorf("cannot raise: matrix is nil or uninitialized")
	}

	for i := range m.Data {
		m.Data[i] = math.Pow(m.Data[i], power)
	}

	return m, nil
}
