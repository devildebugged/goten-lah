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

func GetRow(m *Matx, row int) ([]float64, error) {
	if row < 0 || row >= m.Dimensions[0] {
		return nil, fmt.Errorf("Row out of range")
	}

	start := row * m.Dimensions[1]
	end := start + m.Dimensions[1]

	return m.Data[start:end], nil
}

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
