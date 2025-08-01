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

func PrintMatx(m *Matx, format ...string) {
	if m == nil || m.Data == nil || len(m.Dimensions) == 0 {
		fmt.Println("Invalid or empty matrix")
		return
	}

	// Default format
	f := "%.4f"

	// Optional alias map
	formatAliases := map[string]string{
		"int":   "%.0f", // no decimal
		"float": "%.4f", // default
		"f":     "%.4f",
		"short": "%.2f",
		"sci":   "%e", // scientific notation
		"g":     "%g", // shortest
	}

	if len(format) > 0 {
		if alias, ok := formatAliases[format[0]]; ok {
			f = alias
		} else {
			f = format[0] // use raw fmt string if not an alias
		}
	}

	data := m.Data
	shape := m.Dimensions

	var printRecursive func(offset, dim, depth int)

	printRecursive = func(offset, dim, depth int) {
		indent := func(d int) {
			for i := 0; i < d; i++ {
				fmt.Print("  ")
			}
		}

		if dim == len(shape)-1 {
			indent(depth)
			fmt.Print("{")
			for i := 0; i < shape[dim]; i++ {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf(f, data[offset+i])
			}
			fmt.Print("}")
		} else {
			indent(depth)
			fmt.Print("{\n")
			stride := 1
			for _, s := range shape[dim+1:] {
				stride *= s
			}
			for i := 0; i < shape[dim]; i++ {
				if i > 0 {
					fmt.Print(",\n")
				}
				printRecursive(offset+i*stride, dim+1, depth+1)
			}
			fmt.Print("\n")
			indent(depth)
			fmt.Print("}")
		}
	}

	printRecursive(0, 0, 0)
	fmt.Println()
}

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
