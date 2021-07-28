package systemdconf

import (
	"strings"
)

var escapeChars = [256]byte{
	'.': '.',
	'/': '-',
	'0': '0', '1': '1', '2': '2', '3': '3', '4': '4', '5': '5', '6': '6', '7': '7', '8': '8', '9': '9',
	':': ':',
	'A': 'A', 'B': 'B', 'C': 'C', 'D': 'D', 'E': 'E', 'F': 'F', 'G': 'G', 'H': 'H', 'I': 'I', 'J': 'J',
	'K': 'K', 'L': 'L', 'M': 'M', 'N': 'N', 'O': 'O', 'P': 'P', 'Q': 'Q', 'R': 'R', 'S': 'S', 'T': 'T',
	'U': 'U', 'V': 'V', 'W': 'W', 'X': 'X', 'Y': 'Y', 'Z': 'Z',
	'_': '_',
	'a': 'a', 'b': 'b', 'c': 'c', 'd': 'd', 'e': 'e', 'f': 'f', 'g': 'g', 'h': 'h', 'i': 'i', 'j': 'j',
	'k': 'k', 'l': 'l', 'm': 'm', 'n': 'n', 'o': 'o', 'p': 'p', 'q': 'q', 'r': 'r', 's': 's', 't': 't',
	'u': 'u', 'v': 'v', 'w': 'w', 'x': 'x', 'y': 'y', 'z': 'z',
}

const hextable = "0123456789abcdef"

// Escape applies the systemd escaping rules to the given non-path unit name.
func Escape(unit string) (escaped string) {
	if len(unit) == 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(escapedLength(unit))

	var i int
	if unit[0] == '.' {
		b.WriteString(`\x2e`)
		i = 1
	} else {
		i = 0
	}

	for i < len(unit) {
		char := unit[i]
		if unescaped := escapeChars[char]; unescaped > 0 {
			b.WriteByte(unescaped)
		} else {
			b.WriteString(`\x`)
			b.WriteByte(hextable[char>>4])
			b.WriteByte(hextable[char&0x0f])
		}
		i++
	}

	return b.String()
}

func escapedLength(unit string) int {
	if len(unit) == 0 {
		return 0
	}

	n := 0

	var i int
	if unit[0] == '.' {
		n += 4
		i = 1
	} else {
		i = 0
	}

	for i < len(unit) {
		if escapeChars[unit[i]] > 0 {
			n += 1
		} else {
			n += 4
		}
		i++
	}

	return n
}
