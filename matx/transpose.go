package matx

import "fmt"

func Transpose(m *Matx) (*Matx, error) {
	if m == nil || m.Data == nil || len(m.Dimensions) != 2 {
		return nil, fmt.Errorf("invalid matrix for transpose")
	}

	rows, cols := m.Dimensions[0], m.Dimensions[1]
	result := make([]float64, len(m.Data))

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j*rows+i] = m.Data[i*cols+j]
		}
	}

	newDims := []int{cols, rows}
	mat, err := New(result, newDims)
	if err != nil {
		return nil, fmt.Errorf("transpose error: %w", err)
	}

	return mat, nil
}
