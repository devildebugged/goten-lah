package matx

import "fmt"

func RowSwap(m *Matx, row1, row2 int) error {
	if m == nil {
		return fmt.Errorf("Matrix is nil")
	}
	if len(m.Dimensions) != 2 {
		return fmt.Errorf("Matrix must be 2D")
	}

	rows, cols := m.Dimensions[0], m.Dimensions[1]

	if row1 < 0 || row1 >= rows || row2 < 0 || row2 >= rows {
		return fmt.Errorf("Row indices out of bounds")
	}

	for j := 0; j < cols; j++ {
		i1 := row1*cols + j
		i2 := row2*cols + j
		m.Data[i1], m.Data[i2] = m.Data[i2], m.Data[i1]
	}

	return nil
}
