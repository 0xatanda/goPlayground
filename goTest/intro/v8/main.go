package main

const (
	French        = "French"
	Spanish       = "Spanish"
	EnglishPrefix = "Hello, "
	FrenchPrefix  = "Bonjour, "
	SpanishPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case Spanish:
		prefix = SpanishPrefix
	case French:
		prefix = FrenchPrefix
	default:
		prefix = EnglishPrefix
	}
	return
}
