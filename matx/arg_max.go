package matx

import "fmt"

func ArgMax(m *Matx, axis int) ([]int, error) {
	if m == nil || axis < 0 || axis >= len(m.Dimensions) {
		return nil, fmt.Errorf("Invalid input or axis")
	}

	stride := 1
	for i := axis + 1; i < len(m.Dimensions); i++ {
		stride *= m.Dimensions[i]
	}
	block := m.Dimensions[axis] * stride
	numBlocks := len(m.Data) / block

	result := make([]int, numBlocks*stride)

	for b := 0; b < numBlocks; b++ {
		for i := 0; i < stride; i++ {
			maxVal := m.Data[b*block+i]
			maxIdx := 0
			for j := 1; j < m.Dimensions[axis]; j++ {
				idx := b*block + j*stride + i
				if m.Data[idx] > maxVal {
					maxVal = m.Data[idx]
					maxIdx = j
				}
			}
			result[b*stride+i] = maxIdx
		}
	}
	return result, nil
}
