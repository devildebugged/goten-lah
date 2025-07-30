package matx

import "fmt"

func Sum(m *Matx, axis int) ([]float64, error) {
	if m == nil {
		return nil, fmt.Errorf("Matrix is nil")
	}
	if axis < 0 || axis >= len(m.Dimensions) {
		return nil, fmt.Errorf("Invalid axis")
	}

	stride := 1
	for i := axis + 1; i < len(m.Dimensions); i++ {
		stride *= m.Dimensions[i]
	}

	block := m.Dimensions[axis] * stride
	numBlocks := len(m.Data) / block

	outShape := append([]int{}, m.Dimensions[:axis]...)
	outShape = append(outShape, m.Dimensions[axis+1:]...)
	outSize := 1
	for _, s := range outShape {
		outSize *= s
	}

	result := make([]float64, outSize)

	for b := 0; b < numBlocks; b++ {
		for i := 0; i < stride; i++ {
			sum := 0.0
			for j := 0; j < m.Dimensions[axis]; j++ {
				idx := b*block + j*stride + i
				sum += m.Data[idx]
			}
			result[b*stride+i] = sum
		}
	}

	return result, nil
}
