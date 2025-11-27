package convector_

import (
	"fmt"
	"strconv"
	"unicode"
)

func StringToRunes(text string) (int, string, int, error) {
	runes := []rune(text)

	for i, char := range runes {
		if char == '+' || char == '-' || char == '/' || char == '*' {
			// Проверяем что соседние символы - цифры
			if i > 0 && unicode.IsDigit(runes[i-1]) && i < len(runes)-1 && unicode.IsDigit(runes[i+1]) {
				left := string(runes[:i])
				operator := string(char)
				right := string(runes[i+1:])

				a, _ := strconv.Atoi(left)
				b, _ := strconv.Atoi(right)

				return a, operator, b, nil
			} else {
				return 0, "", 0, fmt.Errorf("недопустимые символы")
			}
		}
	}
	return 0, "", 0, fmt.Errorf("нет выражения")
}
