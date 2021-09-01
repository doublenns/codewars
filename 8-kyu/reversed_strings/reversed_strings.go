package reversedstrings

// Solution that works w/ a string that includes grapheme clustered chars
// (combined chars) such as 🙏2️⃣👌👁 which reverses to 👁👌2️⃣🙏.
// https://stackoverflow.com/a/44350406/15246278

import "unicode"

func reverse(runes []rune, length int) {
	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func Solution(word string) string {
	textRunes := []rune(word)
	textRunesLength := len(textRunes)
	if textRunesLength <= 1 {
		return word
	}

	i, j := 0, 0
	for i < textRunesLength && j < textRunesLength {
		j = i + 1
		for j < textRunesLength && unicode.IsMark(textRunes[j]) {
			j++
		}

		if unicode.IsMark(textRunes[j-1]) {
			// Reverses Combined Characters
			reverse(textRunes[i:j], j-i)
		}

		i = j
	}

	// Reverses the entire array
	reverse(textRunes, textRunesLength)

	return string(textRunes)
}
