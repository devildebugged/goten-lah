package matx

import "fmt"

func Scale(m *Matx, n int) (*Matx, error) {
	if m == nil {
		return nil, fmt.Errorf("Nil matric given")
	}

	for i := 0; i < len(m.Data); i++ {
		m.Data[i] *= float64(n)
	}

	return m, nil
}
