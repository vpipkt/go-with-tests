package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishDefaultName = "World"

const norwegianBokmal = "nb"
const norwegianBokmalHelloPrefix = "Hei, "
const norwegianBokmalDefaultName = "verden"

func Hello(name string, language string) string {

	if language == norwegianBokmal {
		if name == "" {
			name = norwegianBokmalDefaultName
		}
		return norwegianBokmalHelloPrefix + name
	}

	// default or unknown language: en
	if name == "" {
		name = englishDefaultName
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
