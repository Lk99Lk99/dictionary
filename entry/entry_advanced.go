package entry

import (
	"strings"
)

// FromWordPairCsv erzeugt ein neues Entry-Objekt aus einem String,
// der ein Wortpaar enthält. Das deutsche und das englische Wort sind
// durch ein Komma getrennt.
// Gibt es keines oder mehrere Kommas im String, wird ein leerer Eintrag zurückgegeben.
func FromWordPairCsv(pair string) Entry {
	// TODO
	count := 0
	x := ","
	for i := 0; i < len(pair); i++ {

		if pair[i] == x[0] {
			count++
		}
	}

	if count != 1 {
		return Empty()
	}
	wort := strings.Split(pair, ",")
	return New(wort[0], wort[1])

}
