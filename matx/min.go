package matx

import "fmt"

func Min(m *Matx, axis int) ([]float64, error) {
	if m == nil || axis < 0 || axis >= len(m.Dimensions) {
		return nil, fmt.Errorf("Invalid input or axis")
	}

	stride := 1
	for i := axis + 1; i < len(m.Dimensions); i++ {
		stride *= m.Dimensions[i]
	}
	block := m.Dimensions[axis] * stride
	numBlocks := len(m.Data) / block

	result := make([]float64, numBlocks*stride)
	for i := range result {
		result[i] = m.Data[i]
	}

	for b := 0; b < numBlocks; b++ {
		for i := 0; i < stride; i++ {
			for j := 1; j < m.Dimensions[axis]; j++ {
				idx := b*block + j*stride + i
				outIdx := b*stride + i
				if m.Data[idx] < result[outIdx] {
					result[outIdx] = m.Data[idx]
				}
			}
		}
	}
	return result, nil
}
