package matx

import "fmt"

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
