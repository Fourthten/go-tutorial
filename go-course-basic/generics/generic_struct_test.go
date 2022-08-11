package generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First: "Setia",
		Second: "Ajung",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First: "Setia",
		Second: "Ajung",
	}

	assert.Equal(t, "Setia", data.ChangeFirst("Setia"))
	assert.Equal(t, "Hello Setia", data.SayHello("Setia"))
}