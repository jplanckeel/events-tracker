package print

import (
	"fmt"
	"strconv"
)

// formatLine formate une ligne avec une largeur spécifique
func FormatLine(data interface{}, width int) string {
	return padRight(toString(data), width)
}

// padRight complète la chaîne avec des espaces à droite pour atteindre une largeur spécifique
func padRight(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	return s + spaces(width-len(s))
}

// spaces génère une chaîne d'espaces
func spaces(count int) string {
	return "                                                                "[:count]
}

// toString convertit une valeur en chaîne
func toString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	default:
		return fmt.Sprint(value)
	}
}
