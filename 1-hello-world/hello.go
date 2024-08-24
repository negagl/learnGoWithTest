package main

import "fmt"

const (
	helloPrefixEn = "Hello, "
	helloPrefixEs = "Hola, "
	helloPrefixKr = " 안녕하세요"
	exclamation   = "!"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "ES":
		return helloPrefixEs + name + exclamation
	case "KR":
		return name + helloPrefixKr + exclamation
	}

	return helloPrefixEn + name + exclamation
}

func main() {
	fmt.Println(Hello("", ""))
}
