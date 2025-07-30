package matx

import "fmt"

var matxExamples = map[string]*Matx{}

func InitExamples() {
	examples := []struct {
		name string
		data []float64
		dims []int
	}{
		{"matx2x2", []float64{1, 2, 3, 4}, []int{2, 2}},
		{"matx3x1", []float64{5, 6, 7}, []int{3, 1}},
		{"matx3x2", []float64{1, 2, 3, 4, 5, 6}, []int{3, 2}},
		{"matx3x3", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3}},
		{"matx1x4", []float64{9, 8, 7, 6}, []int{1, 4}},
		{"matxIdentity2x2", []float64{1, 0, 0, 1}, []int{2, 2}},
		{"matxIdentity3x3", []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}, []int{3, 3}},
		{"matxZero2x2", []float64{0, 0, 0, 0}, []int{2, 2}},
		{"matxZero3x3", make([]float64, 9), []int{3, 3}},
		{"matxDiag3x3", []float64{1, 0, 0, 0, 2, 0, 0, 0, 3}, []int{3, 3}},
		{"matxRow1x5", []float64{1, 2, 3, 4, 5}, []int{1, 5}},
		{"matxCol5x1", []float64{1, 2, 3, 4, 5}, []int{5, 1}},
		{"matxCube2x2x2", []float64{1, 2, 3, 4, 5, 6, 7, 8}, []int{2, 2, 2}},
		{"matxSymmetric3x3", []float64{1, 2, 3, 2, 4, 5, 3, 5, 6}, []int{3, 3}},
		{"matxUpperTri3x3", []float64{1, 2, 3, 0, 4, 5, 0, 0, 6}, []int{3, 3}},
		{"matxLowerTri3x3", []float64{1, 0, 0, 2, 3, 0, 4, 5, 6}, []int{3, 3}},
		{"matxSparse4x4", []float64{0, 0, 0, 1, 0, 2, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0}, []int{4, 4}},
		{"matxMagic3x3", []float64{8, 1, 6, 3, 5, 7, 4, 9, 2}, []int{3, 3}},
		{"matxAntiDiag3x3", []float64{0, 0, 1, 0, 2, 0, 3, 0, 0}, []int{3, 3}},
		{"matxHilbert3x3", []float64{1.0, 0.5, 0.333, 0.5, 0.333, 0.25, 0.333, 0.25, 0.2}, []int{3, 3}},
	}

	for _, ex := range examples {
		m, err := New(ex.data, ex.dims)
		if err != nil {
			panic(fmt.Sprintf("Failed to create %s: %v", ex.name, err))
		}
		matxExamples[ex.name] = m
	}
}

// Give_matx returns a named example matrix
func GiveMatx(name string) (*Matx, error) {
	if m, ok := matxExamples[name]; ok {
		return m, nil
	}
	return nil, fmt.Errorf("unknown matx: %s", name)
}

// List_matx_examples returns the available names
func List_matx_examples() []string {
	keys := make([]string, 0, len(matxExamples))
	for k := range matxExamples {
		keys = append(keys, k)
	}
	return keys
}
