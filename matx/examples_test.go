package matx_test

import (
	"fmt"
	"goten/matx"
	"goten/utils"
)

func ExampleNew() {
	data := []float64{1, 2, 3, 4}
	dimensions := []int{2, 2}
	m, err := matx.New(data, dimensions)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matx created")
	fmt.Println("Matx data: ", m.Data)
	fmt.Println("Matx dimensions", m.Dimensions)
}

func ExampleOnes() {
	m, err := matx.Ones([]int{2, 2})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ones Matx created")
	fmt.Println("Matx data: ", m.Data)
	fmt.Println("Matx dimensions: ", m.Dimensions)
}

func ExampleZeros() {
	m, err := matx.Zeros([]int{2, 2})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Zeros Matx created")
	fmt.Println("Matx data: ", m.Data)
	fmt.Println("Matx dimensions: ", m.Dimensions)
}

func ExampleDot() {
	m, err := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1})
	n, err2 := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1})

	if err != nil || err2 != nil {
		fmt.Println(err, err2)
		return
	}

	fmt.Println("Matrix 1: ")
	utils.PrintMatx(m)

	fmt.Println("Matrix 2: ")
	utils.PrintMatx(n)

	result, errr := matx.Dot(m, n)
	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Result: ")
	fmt.Println(result)
}

func ExampleAdd() {
	m, err := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})
	n, err2 := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})

	if err != nil || err2 != nil {
		fmt.Println(err, err2)
		return
	}

	fmt.Println("Matrix 1: ")
	utils.PrintMatx(m)

	fmt.Println("Matrix 2: ")
	utils.PrintMatx(n)

	result, errr := matx.Add(m, n)
	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Result: ")
	utils.PrintMatx(result)
}

func ExampleMultiply() {
	m, err := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})
	n, err2 := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})

	if err != nil || err2 != nil {
		fmt.Println(err, err2)
		return
	}

	fmt.Println("Matrix 1: ")
	utils.PrintMatx(m)

	fmt.Println("Matrix 2: ")
	utils.PrintMatx(n)

	result, errr := matx.Multiply(m, n)
	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Result: ")
	utils.PrintMatx(result)
}

func ExampleGet() {
	m, err := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix 1: ")
	utils.PrintMatx(m)

	element, errr := matx.Get(m, 1, 2)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println(element)
}

func ExampleSet() {
	m, err := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix (before): ")
	utils.PrintMatx(m)

	errr := matx.Set(0, m, 1, 1)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Matrix (after): ")
	utils.PrintMatx(m)
}

func ExampleTranspose() {
	m, err := matx.New([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{3, 3})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix (before): ")
	utils.PrintMatx(m)

	tm, errr := matx.Transpose(m)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Matrix (after): ")
	utils.PrintMatx(tm)
}

func ExampleRand() {
	m, err := matx.Rand(3, 3, 1, 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix: ")
	utils.PrintMatx(m)
}

func ExampleInvert() {
	m, err := matx.Rand(3, 3, 1, 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix: ")
	utils.PrintMatx(m)

	im, errr := matx.Invert(m)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Inverted matrix: ")
	utils.PrintMatx(im)
}

func ExampleDet() {
	m, err := matx.Rand(3, 3, 1, 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix: ")
	utils.PrintMatx(m)

	det, errr := matx.Det(m)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Determinant: ")
	fmt.Println(det)
}

func ExampleIdentity() {
	m, err := matx.Identity(3, 3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix: ")
	utils.PrintMatx(m)
}

func ExampleClone() {
	m, err := matx.Rand(3, 3, 1, 100)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Matrix: ")
	utils.PrintMatx(m)

	cm, errr := matx.Clone(m)

	if errr != nil {
		fmt.Println(errr)
		return
	}

	fmt.Println("Cloned matrix: ")
	utils.PrintMatx(cm)
}

func ExampleLUDecomposeWithPivoting() {
	m, err := matx.New([]float64{
		2, 1, 1,
		4, -6, 0,
		-2, 7, 2,
	}, []int{3, 3})

	if err != nil {
		fmt.Println("Error creating matrix:", err)
		return
	}

	L, U, pivots, swaps, err := matx.LUDecomposeWithPivoting(m)
	if err != nil {
		fmt.Println("LU decomposition error:", err)
		return
	}

	fmt.Println("L:")
	utils.PrintMatx(L)

	fmt.Println("U:")
	utils.PrintMatx(U)

	fmt.Println("Pivot indices:", pivots)
	fmt.Println("Number of row swaps:", swaps)
}

func CallAll() {
	ExampleAdd()
	ExampleClone()
	ExampleDet()
	ExampleDot()
	ExampleGet()
	ExampleIdentity()
	ExampleInvert()
	ExampleLUDecomposeWithPivoting()
	ExampleMultiply()
	ExampleNew()
	ExampleOnes()
	ExampleRand()
	ExampleSet()
	ExampleTranspose()
	ExampleZeros()
}
