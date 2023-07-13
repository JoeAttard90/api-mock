package utils

import (
	"reflect"
	"testing"
)

func TestToCamelCase(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"HELLO WORLD", "helloWorld"},
		{"HelloWorld", "helloWorld"},
		{"hello-world", "helloWorld"},
		{"HELLO_WORLD", "helloWorld"},
		{"_hello_world_", "helloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		testName := tt.input
		t.Run(testName, func(t *testing.T) {
			ans := ToCamelCase(tt.input)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"HELLO WORLD", "HelloWorld"},
		{"helloWorld", "HelloWorld"},
		{"hello-world", "HelloWorld"},
		{"HELLO_WORLD", "HelloWorld"},
		{"_hello_world_", "HelloWorld"},
		{"", ""},
	}

	for _, tt := range tests {
		testName := tt.input
		t.Run(testName, func(t *testing.T) {
			ans := ToPascalCase(tt.input)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestPathToTitle(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"/users/{userId}/posts/{postId}", "UsersByUserIdPostsByPostId"},
		{"/home", "Home"},
		{"/shoppingCart/items/{itemId}", "ShoppingCartItemsByItemId"},
		{"/", ""},
		{"/{entity}", "ByEntity"},
		{"", ""},
	}

	for _, tt := range tests {
		testName := tt.input
		t.Run(testName, func(t *testing.T) {
			ans := PathToTitle(tt.input)
			if ans != tt.expected {
				t.Errorf("got %s, want %s", ans, tt.expected)
			}
		})
	}
}

func TestExtractSlugs(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"/users/{userId}/posts/{postId}", []string{"userId", "postId"}},
		{"/home", nil},
		{"/shoppingCart/items/{itemId}", []string{"itemId"}},
		{"/", nil},
		{"/{entity}", []string{"entity"}},
		{"", nil},
		{"/users/{userId}/posts/{postId}/comments/{commentId}", []string{"userId", "postId", "commentId"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ans := ExtractSlugs(tt.input)
			if !reflect.DeepEqual(ans, tt.expected) {
				t.Errorf("ExtractSlugs(%s): expected %v, got %v. Types: %T, %T", tt.input, tt.expected, ans, ans, tt.expected)
			}
		})
	}
}

func BenchmarkToCamelCase(b *testing.B) {
	benchmarks := []struct {
		name string
		str  string
	}{
		{"Single word", "Test"},
		{"Multiple words", "test string"},
		{"Pascal case", "TestCase"},
		{"Punctuation and spaces", "test-string, for example"},
		{"Empty string", ""},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ToCamelCase(benchmark.str)
			}
		})
	}
}

func BenchmarkToPascalCase(b *testing.B) {
	benchmarks := []struct {
		name string
		str  string
	}{
		{"Short String", "hello"},
		{"Long String", "hello_thisIsAString_withDifferent_cases"},
		{"String With Punctuation", "hello, world!"},
		{"Mixed Case String", "HelloWorld"},
		{"Empty String", ""},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ToPascalCase(benchmark.str)
			}
		})
	}
}

func BenchmarkPathToTitle(b *testing.B) {
	benchmarks := []struct {
		name string
		path string
	}{
		{"Short path", "/hello"},
		{"Long path", "/hello/thisIsAPath/withDifferent/segments"},
		{"Path with slugs", "/hello/{slug}/world/{anotherSlug}"},
		{"Empty path", ""},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PathToTitle(benchmark.path)
			}
		})
	}
}

func BenchmarkExtractSlugs(b *testing.B) {
	benchmarks := []struct {
		name string
		path string
	}{
		{"No slugs", "/path/without/any/slugs"},
		{"One slug", "/path/with/{one}Slug"},
		{"Multiple slugs", "/path/with/{multiple}/{different}/{slugs}"},
		{"Slug only", "{slug}"},
		{"Empty path", ""},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ExtractSlugs(benchmark.path)
			}
		})
	}
}
