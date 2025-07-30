package matx

import "fmt"

func Multiply(m1, m2 *Matx) (*Matx, error) {
	if m1 == nil || m2 == nil {
		return nil, fmt.Errorf("one or both input matrices are nil")
	}

	if !CheckMultiplicationCondition(m1.Dimensions, m2.Dimensions) {
		return nil, fmt.Errorf(
			"multiplication not possible: m1 columns (%d) != m2 rows (%d)",
			m1.Dimensions[1], m2.Dimensions[0],
		)
	}

	resultRows := m1.Dimensions[0]
	resultCols := m2.Dimensions[1]
	resultData := make([]float64, resultRows*resultCols)

	result, err := New(resultData, []int{resultRows, resultCols})
	if err != nil {
		return nil, fmt.Errorf("failed to create result matrix: %w", err)
	}

	for i := 0; i < resultRows; i++ {
		for j := 0; j < resultCols; j++ {
			sum := 0.0
			for k := 0; k < m1.Dimensions[1]; k++ {

				a := m1.Data[i*m1.Dimensions[1]+k]
				b := m2.Data[k*m2.Dimensions[1]+j]
				sum += a * b
			}
			result.Data[i*resultCols+j] = sum
		}
	}

	return result, nil
}
