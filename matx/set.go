package matx

import "fmt"

func Set(a float64, m *Matx, coordinates ...int) error {
	if m == nil {
		return fmt.Errorf("given matrix is nil")
	}

	if len(coordinates) != len(m.Dimensions) {
		return fmt.Errorf("matrix dimensions: %v, given coordinates: %v", m.Dimensions, coordinates)
	}

	for i, value := range coordinates {
		if value < 0 {
			return fmt.Errorf("negative coordinates aren't allowed")
		}
		if value >= m.Dimensions[i] {
			return fmt.Errorf("coordinate %d out of bounds for dimension size %d", value, m.Dimensions[i])
		}
	}

	index := 0
	stride := 1
	for i := len(m.Dimensions) - 1; i >= 0; i-- {
		index += coordinates[i] * stride
		stride *= m.Dimensions[i]
	}

	m.Data[index] = a
	return nil
}
