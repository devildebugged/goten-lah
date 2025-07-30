package matx

import "fmt"

func Clone(m *Matx) (*Matx, error) {

	cloneData := make([]float64, len(m.Data))
	copy(cloneData, m.Data)

	cloneDimensions := make([]int, len(m.Dimensions))
	copy(cloneDimensions, m.Dimensions)

	cloneMatx, err := New(cloneData, cloneDimensions)
	if err != nil {
		return nil, fmt.Errorf("failed to clone matrix: %w", err)
	}

	return cloneMatx, nil
}
