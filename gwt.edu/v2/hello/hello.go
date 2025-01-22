package main

import (
	"fmt"
	"os"
)

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
	prefix := greetingPrefix(language)

	name = nameOrDefault(name, language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case norwegianBokmal:
		prefix = norwegianBokmalHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func nameOrDefault(name, language string) (displayName string) {
	displayName = name
	if name == "" {
		switch language {
		case norwegianBokmal:
			displayName = norwegianBokmalDefaultName
		case french:
			displayName = frenchDefaultName
		case spanish:
			displayName = spanishDefaultName
		default:
			displayName = englishDefaultName
		}
	}
	return
}

func parseArgs() (result [2]string) {
	// read args from command, omitting program name
	args := os.Args[1:]

	// set defaults
	result = [2]string{"", ""}

	// copy args into result
	copy(result[:], args)
	return result
}

func main() {
	args := parseArgs()
	name := args[0]
	lang := args[1]
	fmt.Println(Hello(name, lang))
}
