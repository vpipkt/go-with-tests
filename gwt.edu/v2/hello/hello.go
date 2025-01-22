package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "nb" {
		return "Hei, " + name
	}

	// default or unknown language: en
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
