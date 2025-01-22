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

func parseArgs() (args []string) {
	args = os.Args[1:]
	if len(args) == 0 {
		args = append(args, "")
		args = append(args, "")
	}
	if len(args) == 1 {
		args = append(args, "")
	}
	return
}

func main() {
	args := parseArgs()
	name := args[0]
	lang := args[1]
	fmt.Println(Hello(name, lang))
}
