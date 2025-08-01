package matx

import "fmt"

func Mean(m *Matx, axis int) ([]float64, error) {
	sum, err := Sum(m, axis)
	if err != nil {
		return nil, err
	}

	count := float64(m.Dimensions[axis])

	for i := range sum {
		sum[i] /= count
	}

	return sum, nil
}

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

func Max(m *Matx, axis int) ([]float64, error) {
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
				if m.Data[idx] > result[outIdx] {
					result[outIdx] = m.Data[idx]
				}
			}
		}
	}
	return result, nil
}

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

func ArgMin(m *Matx, axis int) ([]int, error) {
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
			minVal := m.Data[b*block+i]
			minIdx := 0
			for j := 1; j < m.Dimensions[axis]; j++ {
				idx := b*block + j*stride + i
				if m.Data[idx] < minVal {
					minVal = m.Data[idx]
					minIdx = j
				}
			}
			result[b*stride+i] = minIdx
		}
	}
	return result, nil
}
