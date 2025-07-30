package matx

import "fmt"

func GetCol(m *Matx, col int) ([]float64, error) {
	if col < 0 || col >= m.Dimensions[1] {
		return nil, fmt.Errorf("Column out of range")
	}

	result := make([]float64, m.Dimensions[0])

	for i := 0; i < m.Dimensions[0]; i++ {
		result[i] = m.Data[i*m.Dimensions[1]+col]
	}

	return result, nil

}
