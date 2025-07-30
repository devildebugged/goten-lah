package matx

func mustGet(m *Matx, coords ...int) float64 {
	val, err := Get(m, coords...)
	if err != nil {
		panic(err)
	}
	return val
}

func mustSet(val float64, m *Matx, coords ...int) {
	if err := Set(val, m, coords...); err != nil {
		panic(err)
	}
}
