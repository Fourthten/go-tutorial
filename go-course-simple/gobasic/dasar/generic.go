package main

import "fmt"

type Number interface {
	int64 | float64
}

type UserModel[T int | float64] struct {
	Name string
	Scores []T
}

func Sum[V int](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

func Sum2(numbers []int) int {
	var total int
	for _, e := range numbers {
		total += e
	}
	return total
}

func Sum3[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total: ", total1)

	total2 := Sum2([]int{1, 2, 3, 4, 5})
	fmt.Println("total: ", total2)

	total3 := Sum3([]float32{2.5, 5.1})
	fmt.Println("total: ", total3)

	total4 := Sum3([]float64{1.32, 2.33, 11.3})
	fmt.Println("total: ", total4)

	total5 := Sum3[int]([]int{1, 2, 3, 4, 5})
	fmt.Println("total: ", total5)

	total6 := Sum3[float32]([]float32{2.5, 5.1})
	fmt.Println("total: ", total6)

	total7 := Sum3[float64]([]float64{1.32, 2.33, 11.3})
	fmt.Println("total: ", total7)

	ints := map[string]int64{ "first": 34, "second": 12 }
	floats := map[string]float64{ "first": 30.2, "second": 21.3 }
	fmt.Printf("Non generic: %v\n",
		SumNumbers1(ints))

	fmt.Printf("Generic sums with constraint: %v and %v\n",
		SumNumbers2(ints),
		SumNumbers2(floats))

	var m1 UserModel[int]
	m1.Name = "Setia"
	m1.Scores = []int{1, 2, 3}
	fmt.Println("scores: ", m1.Scores)

	var m2 UserModel[float64]
	m2.Name = "Ajung"
	m2.SetScoresB([]float64{10, 11, 12})
	fmt.Println("scores: ", m2.Scores)
}

func SumNumbers1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func (m *UserModel[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

// generic ready on 1.18 beta
// go install golang.org/dl/go1.18beta2@latest
// go1.18beta2 run generic.go

