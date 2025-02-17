package main

const spanish = "Spanish"
const french = "French"
const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name, language string) string {

	if name == "" {
		return "Hello, World"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}

	return
}
