package main

import "fmt"

const (
	Spanish       = "Spanish"
	French        = "French"
	EnglishPrefix = "Hello, "
	SpanishPrefix = "Hola, "
	FrenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == Spanish {
		return SpanishPrefix + name
	}

	if language == French {
		return FrenchPrefix + name
	}
	return EnglishPrefix + name
}

func main() {
	fmt.Println(Hello("Atanda", ""))
}
