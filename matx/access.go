package matx

import "fmt"

// Get retrieves the value at the specified multi-dimensional `coordinates` in the matrix `m`.
// Returns an error if the coordinates are out of bounds or the matrix is nil.
func Get(m *Matx, coordinates ...int) (float64, error) {
	if m == nil {
		return 0, fmt.Errorf("given matrix is nil")
	}

	if len(coordinates) != len(m.Dimensions) {
		return 0, fmt.Errorf("matrix dimensions: %v, given coordinates: %v", m.Dimensions, coordinates)
	}

	// Bounds check for each dimension
	for i, value := range coordinates {
		if value < 0 {
			return 0, fmt.Errorf("negative coordinates aren't allowed")
		}
		if value >= m.Dimensions[i] {
			return 0, fmt.Errorf("coordinate %d out of bounds for dimension size %d", value, m.Dimensions[i])
		}
	}

	// Compute flattened index using row-major order
	index := 0
	stride := 1
	for i := len(m.Dimensions) - 1; i >= 0; i-- {
		stride *= m.Dimensions[i]
	}
	stride /= m.Dimensions[0]

	for i := 0; i < len(m.Dimensions); i++ {
		index += coordinates[i] * stride
		stride /= m.Dimensions[i]
	}

	return m.Data[index], nil
}

// Set assigns the value `a` at the specified `coordinates` in matrix `m`.
// Returns an error if the coordinates are invalid or matrix is nil.
func Set(a float64, m *Matx, coordinates ...int) error {
	if m == nil {
		return fmt.Errorf("given matrix is nil")
	}

	if len(coordinates) != len(m.Dimensions) {
		return fmt.Errorf("matrix dimensions: %v, given coordinates: %v", m.Dimensions, coordinates)
	}

	// Bounds check for each dimension
	for i, value := range coordinates {
		if value < 0 {
			return fmt.Errorf("negative coordinates aren't allowed")
		}
		if value >= m.Dimensions[i] {
			return fmt.Errorf("coordinate %d out of bounds for dimension size %d", value, m.Dimensions[i])
		}
	}

	// Compute flattened index (row-major order)
	index := 0
	stride := 1
	for i := len(m.Dimensions) - 1; i >= 0; i-- {
		index += coordinates[i] * stride
		stride *= m.Dimensions[i]
	}

	m.Data[index] = a
	return nil
}

// GetRow returns a slice representing the `row`th row of a 2D matrix `m`.
// Returns an error if `row` is out of bounds.
// Assumes matrix has exactly two dimensions.
func GetRow(m *Matx, row int) ([]float64, error) {
	if len(m.Dimensions) != 2 {
		return nil, fmt.Errorf("GetRow only supports 2D matrices")
	}
	if row < 0 || row >= m.Dimensions[0] {
		return nil, fmt.Errorf("row out of range")
	}

	start := row * m.Dimensions[1]
	end := start + m.Dimensions[1]

	return m.Data[start:end], nil
}

// GetCol returns a slice representing the `col`th column of a 2D matrix `m`.
// Returns an error if `col` is out of bounds.
// Assumes matrix has exactly two dimensions.
func GetCol(m *Matx, col int) ([]float64, error) {
	if len(m.Dimensions) != 2 {
		return nil, fmt.Errorf("GetCol only supports 2D matrices")
	}
	if col < 0 || col >= m.Dimensions[1] {
		return nil, fmt.Errorf("column out of range")
	}

	result := make([]float64, m.Dimensions[0])
	for i := 0; i < m.Dimensions[0]; i++ {
		result[i] = m.Data[i*m.Dimensions[1]+col]
	}

	return result, nil
}
