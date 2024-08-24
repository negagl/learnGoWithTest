package helloworld

import "fmt"

const (
	spanish = "ES"
	english = "EN"
	korean  = "KR"
	helloPrefixEn = "Hello, "
	helloPrefixEs = "Hola, "
	helloPrefixKr = "안녕하세요, "
	exclamation   = "!"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return LanguageSwitcher(language) + name + exclamation
}

func LanguageSwitcher(language string) (prefix string) {
	switch language {
		case spanish:
			prefix = helloPrefixEs
		case korean:
			prefix = helloPrefixKr
		default:
			prefix = helloPrefixEn
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
