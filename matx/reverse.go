package matx

import "fmt"

func Reverse(m *Matx, axis int) (*Matx, error) {
	if m == nil {
		return nil, fmt.Errorf("Matrix is nil")
	}
	if axis < 0 || axis >= len(m.Dimensions) {
		return nil, fmt.Errorf("Invalid axis")
	}

	out := make([]float64, len(m.Data))
	copy(out, m.Data)

	stride := 1
	for i := axis + 1; i < len(m.Dimensions); i++ {
		stride *= m.Dimensions[i]
	}

	block := m.Dimensions[axis] * stride
	numBlocks := len(m.Data) / block

	for b := 0; b < numBlocks; b++ {
		for i := 0; i < stride; i++ {
			for j := 0; j < m.Dimensions[axis]; j++ {

				srcIdx := b*block + j*stride + i
				dstIdx := b*block + (m.Dimensions[axis]-1-j)*stride + i
				out[dstIdx] = m.Data[srcIdx]
			}
		}
	}

	return &Matx{
		Data:       out,
		Dimensions: append([]int{}, m.Dimensions...),
	}, nil
}
