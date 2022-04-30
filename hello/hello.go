package hello

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

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
