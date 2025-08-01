package matx

import "fmt"

// Mean computes the mean (average) of the matrix values along the specified axis.
// Parameters:
// - m: input matrix
// - axis: the axis along which the mean is computed
// Returns:
// - A slice containing the mean values along the axis
// - An error if computation fails (e.g., invalid axis)
func Mean(m *Matx, axis int) ([]float64, error) {
	sum, err := Sum(m, axis)
	if err != nil {
		return nil, err
	}

	count := float64(m.Dimensions[axis])

	// Divide each summed element by count to get the mean
	for i := range sum {
		sum[i] /= count
	}

	return sum, nil
}

// Sum computes the sum of matrix values along the specified axis.
// Parameters:
// - m: input matrix
// - axis: the axis along which the summation is performed
// Returns:
// - A slice containing the summed values along the axis
// - An error if input is nil or axis is out of bounds
func Sum(m *Matx, axis int) ([]float64, error) {
	if m == nil {
		return nil, fmt.Errorf("Matrix is nil")
	}
	if axis < 0 || axis >= len(m.Dimensions) {
		return nil, fmt.Errorf("Invalid axis")
	}

	// Calculate stride: product of dimensions after the target axis
	stride := 1
	for i := axis + 1; i < len(m.Dimensions); i++ {
		stride *= m.Dimensions[i]
	}

	block := m.Dimensions[axis] * stride // Total elements in a single axis block
	numBlocks := len(m.Data) / block     // Number of such blocks in the flat data slice

	// Calculate output shape (excluding the specified axis)
	outShape := append([]int{}, m.Dimensions[:axis]...)
	outShape = append(outShape, m.Dimensions[axis+1:]...)

	// Determine size of output
	outSize := 1
	for _, s := range outShape {
		outSize *= s
	}

	result := make([]float64, outSize)

	// Perform summation along the axis
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

// Min returns the minimum value along the specified axis.
// Parameters:
// - m: input matrix
// - axis: axis along which the minimum is calculated
// Returns:
// - A slice of minimum values along the axis
// - An error if input is nil or axis is invalid
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

	// Initialize result with the first element of each slice
	for i := range result {
		result[i] = m.Data[i]
	}

	// Compute minimum across axis
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

// Max returns the maximum value along the specified axis.
// Parameters:
// - m: input matrix
// - axis: axis along which the maximum is calculated
// Returns:
// - A slice of maximum values along the axis
// - An error if input is nil or axis is invalid
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

	// Initialize result with the first element of each slice
	for i := range result {
		result[i] = m.Data[i]
	}

	// Compute maximum across axis
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

// ArgMax returns the indices of the maximum values along the specified axis.
// Parameters:
// - m: input matrix
// - axis: axis along which the argmax is calculated
// Returns:
// - A slice of indices corresponding to the maximum value in each slice
// - An error if input is nil or axis is invalid
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

	// Compute index of maximum value across axis
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

// ArgMin returns the indices of the minimum values along the specified axis.
// Parameters:
// - m: input matrix
// - axis: axis along which the argmin is calculated
// Returns:
// - A slice of indices corresponding to the minimum value in each slice
// - An error if input is nil or axis is invalid
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

	// Compute index of minimum value across axis
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
