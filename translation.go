package main

import (
	"fmt"
	"strings"

	"github.com/4stroPhysik3r/CyrilicTranslator.git/dictionary"
)

var input string

func main() {
	fmt.Println(">> Welcome to the cyrillic translator!")
	println(">> Type here the name you want to translate to the english alphabet:")

	fmt.Scan(&input)
	output := strings.Title(strings.ToLower(dictionary.ReplaceCharacters(strings.ToUpper(input), dictionary.Replacements)))
	println(">> Translated to the english alphabet:")
	println(output)
}
