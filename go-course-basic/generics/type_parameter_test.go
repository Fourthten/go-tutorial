package generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Setia")
	assert.Equal(t, "Setia", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}