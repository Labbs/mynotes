package block

func GenerateBlockPosition(a, b string) string {
	// Si pas de borne inférieure, on commence par 'a'
	if a == "" {
		return "a"
	}

	// Si pas de borne supérieure, on ajoute 'a' à la borne inférieure
	if b == "" {
		return a + "a"
	}

	// Trouve le préfixe commun
	prefixLen := 0
	for prefixLen < len(a) && prefixLen < len(b) && a[prefixLen] == b[prefixLen] {
		prefixLen++
	}

	// Si a est un préfixe de b, ajoute un caractère à a
	if prefixLen == len(a) {
		return a + "a"
	}

	// Sinon, incrémente le premier caractère différent
	newChar := byte((int(a[prefixLen]) + int(b[prefixLen])) / 2)
	if newChar == a[prefixLen] {
		return a + "a"
	}

	return a[:prefixLen] + string(newChar)
}
