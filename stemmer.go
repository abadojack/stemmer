package stemmer

import (
	"regexp"
	"strings"
)

//roots represents standalone root words
var roots = []string{"ĉar", "ĉi", "ĉu", "kaj", "ke", "la", "minus", "plus",
	"se", "ĉe", "da", "de", "el", "ekster", "en", "ĝis", "je", "kun", "na",
	"po", "pri", "pro", "sen", "tra", "ajn", "do", "ja", "jen", "ju", "ne",
	"pli", "tamen", "tre", "tro", "ci", "ĝi", "ili", "li", "mi", "ni", "oni",
	"ri", "si", "ŝi", "ŝli", "vi", "unu", "du", "tri", "kvin", "ĵus", "nun", "plu",
	"tuj", "amen", "bis", "boj", "fi", "ha", "he", "ho", "hu", "hura", "nu", "ve",
	"esperanto",
}

//correlative roots
var correlatives = []string{
	"kia", "kial", "kiam", "kie", "kiel", "kies", "kio", "kiom", "kiu",
	"tia", "tial", "tiam", "tie", "tiel", "ties", "tio", "tiom", "tiu",
	"ia", "ial", "iam", "ie", "iel", "ies", "io", "iom", "iu",
	"ĉia", "ĉial", "ĉiam", "ĉie", "ĉiel", "ĉies", "ĉio", "ĉiom", "ĉiu",
	"nenia", "nenial", "neniam", "nenie", "neniel", "nenies", "nenio", "neniom", "neniu",
}

//Stem returns the word's base form.
func Stem(word string) string {
	word = strings.TrimSpace(strings.ToLower(word))
	//standalone roots
	for _, root := range roots {
		if word == root {
			return word
		}
	}
	//l’ l' → la
	if word == "l’" || word == "l'" {
		return "la"
	}

	//un’ un' unuj → unu
	if word == "un’" || word == "un'" || word == "unuj" {
		return "unu"
	}

	//-’ -' → -o
	word = replaceSuffix(word, "'", "o")
	word = replaceSuffix(word, "’", "o")

	//’st- 'st- → est-
	word = replacePrefix(word, "’st", "est")
	word = replacePrefix(word, "'st", "est")

	// nouns, adjectives, -u correlatives:
	// -oj -on -ojn → o
	// -aj -an -ajn → a
	// -uj -un -ujn → u
	word = replaceSuffix(word, "oj", "o")
	word = replaceSuffix(word, "on", "o")
	word = replaceSuffix(word, "ojn", "o")
	word = replaceSuffix(word, "aj", "a")
	word = replaceSuffix(word, "an", "a")
	word = replaceSuffix(word, "ajn", "a")
	word = replaceSuffix(word, "uj", "u")
	word = replaceSuffix(word, "un", "u")
	word = replaceSuffix(word, "ujn", "u")

	//correlatives: -en → -e
	for _, s := range []string{"kien", "tien", "ien", "nenien", "ĉien"} {
		if word == s {
			return word[:len(word)-1]
		}
	}

	//correlative roots
	for _, s := range correlatives {
		if word == s {
			return word
		}
	}

	//accusative pronouns: -in → -i
	if strings.HasSuffix(word, "in") {
		return replaceSuffix(word, "in", "i")
	}

	//accusative adverbs: -en → -o
	word = replaceSuffix(word, "en", "o")

	//verbs: -is -as -os -us -u → -i
	word = replaceSuffix(word, "is", "i")
	word = replaceSuffix(word, "as", "i")
	word = replaceSuffix(word, "os", "i")
	word = replaceSuffix(word, "us", "i")
	word = replaceSuffix(word, "u", "i")

	//lexical aspect: ek- el-
	if !strings.HasPrefix(word, "ekscit") {
		word = strings.TrimPrefix(word, "ek")
	}
	if !strings.HasPrefix(word, "elefant") {
		word = strings.TrimPrefix(word, "el")
	}

	//imperfective verbs & action nouns: -adi -ado → -i
	if strings.HasSuffix(word, "adi") {
		return replaceSuffix(word, "adi", "i")
	} else if strings.HasSuffix(word, "ado") {
		return replaceSuffix(word, "ado", "i")
	}

	// compound verbs:
	// -inti -anti -onti -iti -ati -oti → -i
	// -inte -ante -onte -ite -ate -ote → -i
	// -inta -anta -onta -ita -ata -ota → -i
	switch {
	case strings.HasSuffix(word, "inti"):
		return replaceSuffix(word, "inti", "i")
	case strings.HasSuffix(word, "anti"):
		return replaceSuffix(word, "anti", "i")
	case strings.HasSuffix(word, "onti"):
		return replaceSuffix(word, "onti", "i")
	case strings.HasSuffix(word, "iti"):
		return replaceSuffix(word, "iti", "i")
	case strings.HasSuffix(word, "ati"):
		return replaceSuffix(word, "ati", "i")
	case strings.HasSuffix(word, "oti"):
		return replaceSuffix(word, "oti", "i")
	case strings.HasSuffix(word, "inte"):
		return replaceSuffix(word, "inte", "i")
	case strings.HasSuffix(word, "ante"):
		return replaceSuffix(word, "ante", "i")
	case strings.HasSuffix(word, "onte"):
		return replaceSuffix(word, "onte", "i")
	case strings.HasSuffix(word, "ite"):
		return replaceSuffix(word, "ite", "i")
	case strings.HasSuffix(word, "ate"):
		return replaceSuffix(word, "ate", "i")
	case strings.HasSuffix(word, "ote"):
		return replaceSuffix(word, "ote", "i")
	case strings.HasSuffix(word, "inta"):
		return replaceSuffix(word, "inta", "i")
	case strings.HasSuffix(word, "anta"):
		return replaceSuffix(word, "anta", "i")
	case strings.HasSuffix(word, "onta"):
		return replaceSuffix(word, "onta", "i")
	case strings.HasSuffix(word, "ita"):
		return replaceSuffix(word, "ita", "i")
	case strings.HasSuffix(word, "ata"):
		return replaceSuffix(word, "ata", "i")
	case strings.HasSuffix(word, "ota"):
		return replaceSuffix(word, "ota", "i")
	}

	// participle nouns:
	// -into -anto -onto → -anto
	// -ito  -ato  -oto  → -ato
	switch {
	case strings.HasSuffix(word, "into"):
		return replaceSuffix(word, "into", "anto")
	case strings.HasSuffix(word, "anto"):
		return replaceSuffix(word, "anto", "anto")
	case strings.HasSuffix(word, "onto"):
		return replaceSuffix(word, "onto", "anto")
	case strings.HasSuffix(word, "ito"):
		return replaceSuffix(word, "ito", "ato")
	case strings.HasSuffix(word, "ato"):
		return replaceSuffix(word, "ato", "ato")
	case strings.HasSuffix(word, "oto"):
		return replaceSuffix(word, "oto", "ato")
	}

	return word
}

//StemAggressive performs a aggressive stemming on the word and returns the simplest
//form of it's  root form.
func StemAggressive(word string) string {
	word = Stem(word)

	//root words
	for _, root := range roots {
		if word == root {
			return word
		}
	}

	//remove final suffix if it's a vowel
	if strings.HasSuffix(word, "a") || strings.HasSuffix(word, "e") || strings.HasSuffix(word, "i") ||
		strings.HasSuffix(word, "o") || strings.HasSuffix(word, "u") {
		word = word[:len(word)-1]
	}

	//remove suffix for participle nouns:
	//-int- -ant- -ont- -it- -at- -ot-
	matched, err := regexp.MatchString("[aeiou].*(int|ant|ont|it|at|ot)$", word)
	if err == nil && matched {

		switch {
		case strings.HasSuffix(word, "int") || strings.HasSuffix(word, "ant") || strings.HasSuffix(word, "ont"):
			word = word[:len(word)-3] // I chose not to use strings.TrimSuffix because because it calls strings.HasSuffix again.
		case strings.HasSuffix(word, "it") || strings.HasSuffix(word, "at") || strings.HasSuffix(word, "ot"):
			word = word[:len(word)-2]
		}

	}

	return word
}

//replaceSuffix returns s with the trailing oldSuffix replaced with the newSuffix.
//If s doesn't end with the suffix, s is returned unchanged.
func replaceSuffix(s, oldSuffix, newSuffix string) string {
	if strings.HasSuffix(s, oldSuffix) {
		return s[:len(s)-len(oldSuffix)] + newSuffix
	}
	return s
}

//replacePrefix returns s with the leading oldPrefix replaced with the newPrefix.
//If s doesn't start with the prefix, s is returned unchanged.
func replacePrefix(s, oldPrefix, newPrefix string) string {
	if strings.HasPrefix(s, oldPrefix) {
		return newPrefix + s[len(oldPrefix):]
	}
	return s
}
