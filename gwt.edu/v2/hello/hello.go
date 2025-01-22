package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishDefaultName = "World"

const norwegianBokmal = "nb"
const norwegianBokmalHelloPrefix = "Hei, "
const norwegianBokmalDefaultName = "verden"

const french = "fr"
const frenchHelloPrefix = "Bonjour, "
const frenchDefaultName = "le monde"

const spanish = "es"
const spanishHelloPrefix = "Hola, "
const spanishDefaultName = "mundo"

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
