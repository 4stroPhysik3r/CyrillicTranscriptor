package dictionary

import "strings"

func ReplaceCharacters(input string, replacements map[string]string) string {
	var output strings.Builder
	for _, char := range input {
		charStr := string(char)
		if replacement, exists := replacements[charStr]; exists {
			output.WriteString(replacement)
		} else {
			output.WriteString(charStr)
		}
	}
	return output.String()
}

var Replacements = map[string]string{
	"Б": "b",
	"Ц": "x",
	"Д": "d",
	"Ф": "f",
	"Г": "g",
	"Х": "h",
	"И": "i",
	"Й": "j",
	"Л": "l",
	"Н": "n",
	"П": "p",
	"Р": "r",
	"С": "s",
	"У": "u",
	"В": "w",
	"Ы": "y",
	"З": "z",
	"Ж": "zh",
	"Ё": "jo",
	"Ч": "ch",
	"Э": "je",
	"Ю": "ju",
	"Я": "ja",
}

/*
cyrillic alphabet
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
А Б Ц Д Е Ф Г Х И Й К Л М Н О П Я Р С Т У В В Х Ы З
а б ц д е ф г х и й к л м н о п я р с т у в в х ы з
--------------------
unique letters
B C D F G H I J L N P Q R S U V W Y Z
Б Ц Д Ф Г Х И Й Л Н П Я Р С У В В Ы З
б ц д ф г х и й л н п я р с у в в ы з
--------------------
special cases (2 letters)
Ж   Ё  Ч   Э  Ю  Я
zh jo ch je/ä ju ja
*/
