package matx

import "fmt"

func Negate(m *Matx) (*Matx, error) {
	if m == nil || m.Data == nil {
		return nil, fmt.Errorf("cannot negate: matrix is nil or uninitialized")
	}

	for i := range m.Data {
		m.Data[i] *= -1
	}

	return m, nil
}
