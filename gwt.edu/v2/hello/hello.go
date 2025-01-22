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

func Hello(name string, language string) string {

	prefix := englishHelloPrefix

	switch language {
	case norwegianBokmal:
		prefix = norwegianBokmalHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	if name == "" {
		switch language {
		case norwegianBokmal:
			name = norwegianBokmalDefaultName
		case french:
			name = frenchDefaultName
		default:
			name = englishDefaultName
		}
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
