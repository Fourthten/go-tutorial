package simple

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Setia",
			request: "Setia",
		},
		{
			name:    "Ajung",
			request: "Ajung",
		},
		{
			name:    "Agung",
			request: "Agung",
		},
		{
			name:    "AgungSetia",
			request: "Agung Setia",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Ajung", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ajung")
		}
	})
	b.Run("Setia", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Setia")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ajung")
	}
}

func BenchmarkHelloWorldSetia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Setia")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ajung",
			request:  "Ajung",
			expected: "Hello Ajung",
		},
		{
			name:     "Setia",
			request:  "Setia",
			expected: "Hello Setia",
		},
		{
			name:     "Agung",
			request:  "Agung",
			expected: "Hello Agung",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := HelloWorld("Ajung")
	require.Equal(t, "Hello Ajung", result, "Result must be 'Hello Ajung'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ajung")
	require.Equal(t, "Hello Ajung", result, "Result must be 'Hello Ajung'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ajung")
	assert.Equal(t, "Hello Ajung", result, "Result must be 'Hello Ajung'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAjung(t *testing.T) {
	result := HelloWorld("Ajung")

	if result != "Hello Ajung" {
		// error
		t.Error("Result must be 'Hello Ajung'")
	}
	fmt.Println("TestHelloWorldAjung Done")
}

func TestHelloWorldSetia(t *testing.T) {
	result := HelloWorld("Setia")

	if result != "Hello Setia" {
		// error
		t.Fatal("Result must be 'Hello Setia'")
	}
	fmt.Println("TestHelloWorldSetia Done")
}
