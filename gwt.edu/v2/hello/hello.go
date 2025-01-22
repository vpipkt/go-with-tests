package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	englishDefaultName = "World"

	norwegianBokmal            = "nb"
	norwegianBokmalHelloPrefix = "Hei, "
	norwegianBokmalDefaultName = "verden"

	french            = "fr"
	frenchHelloPrefix = "Bonjour, "
	frenchDefaultName = "le monde"

	spanish            = "es"
	spanishHelloPrefix = "Hola, "
	spanishDefaultName = "mundo"
)

func Hello(name string, language string) string {

	prefix := englishHelloPrefix

	switch language {
	case norwegianBokmal:
		prefix = norwegianBokmalHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	if name == "" {
		switch language {
		case norwegianBokmal:
			name = norwegianBokmalDefaultName
		case french:
			name = frenchDefaultName
		case spanish:
			name = spanishDefaultName
		default:
			name = englishDefaultName
		}
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
