package hello

const (
	spanish = "Spanish"
	french  = "French"
)

// Say hello by passing in `name` and `language`.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// Generate greeting prefix based on provided language.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = "Hola, "
	case french:
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}
	return
}
