package main

import "fmt"

const englishHelloPrefix = "Hello, "

const norwegianBokmal = "nb"
const norwegianBokmalHelloPrefix = "Hei, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == norwegianBokmal {
		return norwegianBokmalHelloPrefix + name
	}

	// default or unknown language: en
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
