package matx

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
