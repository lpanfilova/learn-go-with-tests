package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// public functions start with a capital letter
func Hello(name string, language string) string {

	if name == "" {
		name = worldTranslate(language)
	}

	return greetingPrefix(language) + name
}

// private functions start with a lowercase letter
// named return: prefix string
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func worldTranslate(language string) (name string) {
	switch language {
	case spanish:
		name = "el Mundo"
	case french:
		name = "le Monde"
	default:
		name = "World"
	}

	return
}

func main() {
	fmt.Println(Hello("Jane", "Spanish"))
}
