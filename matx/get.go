package matx

import "fmt"

func Get(m *Matx, coordinates ...int) (float64, error) {
	if m == nil {
		return 0, fmt.Errorf("given matrix is nil")
	}

	if len(coordinates) != len(m.Dimensions) {
		return 0, fmt.Errorf("matrix dimensions: %v, given coordinates: %v", m.Dimensions, coordinates)
	}

	for i, value := range coordinates {
		if value < 0 {
			return 0, fmt.Errorf("negative coordinates aren't allowed")
		}
		if value >= m.Dimensions[i] {
			return 0, fmt.Errorf("coordinate %d out of bounds for dimension size %d", value, m.Dimensions[i])
		}
	}

	index := 0
	stride := 1
	for i := len(m.Dimensions) - 1; i >= 0; i-- {
		stride *= m.Dimensions[i]
	}
	stride = stride / m.Dimensions[0]

	for i := 0; i < len(m.Dimensions); i++ {
		index += coordinates[i] * stride
		stride = stride / m.Dimensions[i]
	}

	return m.Data[index], nil
}
