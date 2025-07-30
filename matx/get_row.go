package matx

import "fmt"

func GetRow(m *Matx, row int) ([]float64, error) {
	if row < 0 || row >= m.Dimensions[0] {
		return nil, fmt.Errorf("Row out of range")
	}

	start := row * m.Dimensions[1]
	end := start + m.Dimensions[1]

	return m.Data[start:end], nil
}
