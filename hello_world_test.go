package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run() // eksekusi semua unit test

	// after
	fmt.Println("After Unit Test")
}

func BenchmarkHelloWorldFerrian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ferrian")
	}
}

func BenchmarkHelloWorldTitus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Titus")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ferrian",
			request:  "Ferrian",
			expected: "Hello Ferrian",
		},
		{
			name:     "Titus",
			request:  "Titus",
			expected: "Hello Titus",
		},
		{
			name:     "Pradana",
			request:  "Pradana",
			expected: "Hello Pradana",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorldUser(t *testing.T) {
	result := HelloWorld("User")
	if result != "Hello User" {
		// unit test failed
		t.Error("Result must be 'Hello User'")
	}

	fmt.Println("TestHelloWorldUser Done")
}

func TestHelloWorldFerri(t *testing.T) {
	result := HelloWorld("Ferri")
	if result != "Hello Ferri" {
		// unit test failed
		t.Fatal("Result must be 'Hello Ferri'")
	}

	fmt.Println("TestHelloWorldFerri Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ferr")
	// assert with Fail()
	assert.Equal(t, "Hello Ferr", result, "Result must be 'Hello Ferr'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Titus")
	// require with FailNow()
	require.Equal(t, "Hello Titus", result, "Result must be 'Hello Titus'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Mac OS")
	}

	result := HelloWorld("Pradana")
	require.Equal(t, "Hello Pradana", result, "Result must be 'Hello Pradana'")
}

func TestSubTest(t *testing.T) {
	t.Run("Ferri", func(t *testing.T) {
		result := HelloWorld("Ferri")
		require.Equal(t, "Hello Ferri", result, "Result must be 'Hello Ferri'")

	})
	t.Run("Titus", func(t *testing.T) {
		result := HelloWorld("Titus")
		require.Equal(t, "Hello Titus", result, "Result must be 'Hello Titus'")

	})
}
