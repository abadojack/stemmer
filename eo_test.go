package stemmer

import "testing"

func Test_StemEO(t *testing.T) {
	stemmer := NewStemmer("eo", false)

	//Test simple stemming.
	for word, want := range testsSimple {
		got := stemmer.Stem(word)
		if want != got {
			t.Fatalf("Simple: word: %s, want: %s, got: %s", word, want, got)
		}
	}
	/*
		//Test aggressive stemming.
		stemmer.Aggressive = true
		for word, want := range testsAggressive {
			got := stemmer.Stem(word)
			if want != got {
				t.Fatalf("Aggressive: word: %s, want: %s, got: %s", word, want, got)
			}
		}
	*/
	//Test standalone root words.
	for _, word := range standAloneRootWords {
		got := stemmer.Stem(word)
		if word != got {
			t.Fatalf("Stand alone root words: word: %s, want: %s, got: %s", word, word, got)
		}
	}

	//Test pronouns.
	for _, word := range pronouns {
		//Prepositions
		got := stemmer.Stem(word)
		if word != got {
			t.Fatalf("Stand alone root words: word: %s, want: %s, got: %s", word, word, got)
		}

		//Accusative Prepositions
		got = stemmer.Stem(word + "n")
		if word != got {
			t.Fatalf("Stand alone root words: word: %s, want: %s, got: %s", word, word, got)
		}
	}

	for _, start := range []string{"ki", "ti", "i", "ĉi", "neni"} {
		for _, end := range []string{"a", "al", "am", "e", "el", "es", "o", "om", "u"} {
			word := start + end
			got := stemmer.Stem(word)
			if word != got {
				t.Fatalf("Correlative: word: %s, want: %s, got: %s", word, word, got)
			}

			if end == "a" || end == "o" || end == "u" {
				for _, end1 := range []string{"j", "n", "jn"} {
					word1 := word + end1
					got = stemmer.Stem(word1)
					if word != got {
						t.Fatalf("Correlative -%s: word: %s, want: %s, got: %s", end1, word1, word, got)
					}
				}
			} else if end == "e" {
				word1 := word + "n"
				got = stemmer.Stem(word1)
				if word != got {
					t.Fatalf("Correlative -%s: word: %s, want: %s, got: %s", "en", word1, word, got)
				}
			}
		}
	}

}

var testsSimple = map[string]string{
	"mia":         "mia",       //-a possesive adjective
	"miaj":        "mia",       //-aj plural possesive adjective
	"mian":        "mia",       //-an accusative possesive adjective
	"miajn":       "mia",       //-ajn accusative plural possesive adjective
	"hundo":       "hundo",     //-o noun
	"hundoj":      "hundo",     //-oj plural noun
	"hundon":      "hundo",     //-on accusative noun
	"hundojn":     "hundo",     //-ojn accusative plural noun
	"longa":       "longa",     //-a adjective
	"longaj":      "longa",     //-aj plural adjective
	"longan":      "longa",     //-an accusative adjective
	"longajn":     "longa",     //-ajn accusative plural adjective
	"labori":      "labori",    //-i infinitive verb
	"laboris":     "labori",    //-is past indicative verb
	"laboras":     "labori",    //-as present indicative verb
	"laboros":     "labori",    //-os future indicative verb
	"laborus":     "labori",    //-us conditional verb
	"laboru":      "labori",    //-u jussive verb
	"labore":      "labore",    //-e adverb
	"hejmen":      "hejmo",     //-en accusative adverb
	"laborinti":   "labori",    //-inti
	"laboranti":   "labori",    //-anti
	"laboronti":   "labori",    //-onti
	"laboriti":    "labori",    //-iti
	"laborati":    "labori",    //-ati
	"laboroti":    "labori",    //-oti
	"laborintis":  "labori",    //-intis
	"laborantis":  "labori",    //-antis
	"laborontis":  "labori",    //-ontis
	"laboritis":   "labori",    //-itis
	"laboratis":   "labori",    //-atis
	"laborotis":   "labori",    //-otis
	"laborintas":  "labori",    //-intas
	"laborantas":  "labori",    //-antas
	"laborontas":  "labori",    //-ontas
	"laboritas":   "labori",    //-itas
	"laboratas":   "labori",    //-atas
	"laborotas":   "labori",    //-otas
	"laborintos":  "labori",    //-intos
	"laborantos":  "labori",    //-antos
	"laborontos":  "labori",    //-ontos
	"laboritos":   "labori",    //-itos
	"laboratos":   "labori",    //-atos
	"laborotos":   "labori",    //-otos
	"laborintus":  "labori",    //-intus
	"laborantus":  "labori",    //-antus
	"laborontus":  "labori",    //-ontus
	"laboritus":   "labori",    //-itus
	"laboratus":   "labori",    //-atus
	"laborotus":   "labori",    //-otus
	"laborintu":   "labori",    //-intu
	"laborantu":   "labori",    //-antu
	"laborontu":   "labori",    //-ontu
	"laboritu":    "labori",    //-itu
	"laboratu":    "labori",    //-atu
	"laborotu":    "labori",    //-otu
	"laborinte":   "labori",    //-inte
	"laborante":   "labori",    //-ante
	"laboronte":   "labori",    //-onte
	"laborite":    "labori",    //-ite
	"laborate":    "labori",    //-ate
	"laborote":    "labori",    //-ote
	"laborinta":   "labori",    //-inta
	"laboranta":   "labori",    //-anta
	"laboronta":   "labori",    //-onta
	"laborita":    "labori",    //-ita
	"laborata":    "labori",    //-ata
	"laborota":    "labori",    //-ota
	"laborintaj":  "labori",    //-intaj
	"laborantaj":  "labori",    //-antaj
	"laborontaj":  "labori",    //-ontaj
	"laboritaj":   "labori",    //-itaj
	"laborataj":   "labori",    //-ataj
	"laborotaj":   "labori",    //-otaj
	"laborintan":  "labori",    //-intan
	"laborantan":  "labori",    //-antan
	"laborontan":  "labori",    //-ontan
	"laboritan":   "labori",    //-itan
	"laboratan":   "labori",    //-atan
	"laborotan":   "labori",    //-otan
	"laborintajn": "labori",    //-intajn
	"laborantajn": "labori",    //-antajn
	"laborontajn": "labori",    //-ontajn
	"laboritajn":  "labori",    //-itajn
	"laboratajn":  "labori",    //-atajn
	"laborotajn":  "labori",    //-otajn
	"laboradi":    "labori",    //-adi
	"laboradas":   "labori",    //-adas
	"laborado":    "labori",    //-ado
	"eklabori":    "labori",    //ek-
	"eksciti":     "eksciti",   //ek- exception
	"ellabori":    "labori",    //el-
	"elefanto":    "elefanto",  //el- exception
	"laboranto":   "laboranto", //-anto
	"laborinto":   "laboranto", //-into
	"laboronto":   "laboranto", //-onto
	"laborato":    "laborato",  //-ato
	"laborito":    "laborato",  //-ito
	"laboroto":    "laborato",  //-oto
	"esperanto":   "esperanto", //esperanto: keep -anto
	"kanto":       "kanto",     //kanto: keep -anto
	"hund’":       "hundo",     //-’ noun with typographic apostrophe
	"hund'":       "hundo",     //-' noun with typewriter apostrophe
	"l’":          "la",        //l’ with typographic apostrophe
	"l'":          "la",        //l' with typewriter apostrophe
	"’stas":       "esti",      //’stas with typographic apostrophe
	"'stas":       "esti",      //'stas with typewriter apostrophe
	"un’":         "unu",       //un’ with typographic apostrophe
	"un'":         "unu",       //un' with typewriter apostrophe
	"unuj":        "unu",       //plural unu
}

var testsAggressive = map[string]string{
	"mia":         "mi",       //-a possesive adjective
	"miaj":        "mi",       //-aj plural possesive adjective
	"mian":        "mi",       //-an accusative possesive adjective
	"miajn":       "mi",       //-ajn accusative plural possesive adjective
	"hundo":       "hund",     //-o noun
	"hundoj":      "hund",     //-oj plural noun
	"hundon":      "hund",     //-on accusative noun
	"hundojn":     "hund",     //-ojn accusative plural noun
	"longa":       "long",     //-a adjective
	"longaj":      "long",     //-aj plural adjective
	"longan":      "long",     //-an accusative adjective
	"longajn":     "long",     //-ajn accusative plural adjective
	"labor":       "labor",    //-i infinitive verb
	"laboris":     "labor",    //-is past indicative verb
	"laboras":     "labor",    //-as present indicative verb
	"laboros":     "labor",    //-os future indicative verb
	"laborus":     "labor",    //-us conditional verb
	"laboru":      "labor",    //-u jussive verb
	"labore":      "labor",    //-e adverb
	"hejmen":      "hejm",     //-en accusative adverb
	"laborinti":   "labor",    //-inti
	"laboranti":   "labor",    //-anti
	"laboronti":   "labor",    //-onti
	"laboriti":    "labor",    //-iti
	"laborati":    "labor",    //-ati
	"laboroti":    "labor",    //-oti
	"laborintis":  "labor",    //-intis
	"laborantis":  "labor",    //-antis
	"laborontis":  "labor",    //-ontis
	"laboritis":   "labor",    //-itis
	"laboratis":   "labor",    //-atis
	"laborotis":   "labor",    //-otis
	"laborintas":  "labor",    //-intas
	"laborantas":  "labor",    //-antas
	"laborontas":  "labor",    //-ontas
	"laboritas":   "labor",    //-itas
	"laboratas":   "labor",    //-atas
	"laborotas":   "labor",    //-otas
	"laborintos":  "labor",    //-intos
	"laborantos":  "labor",    //-antos
	"laborontos":  "labor",    //-ontos
	"laboritos":   "labor",    //-itos
	"laboratos":   "labor",    //-atos
	"laborotos":   "labor",    //-otos
	"laborintus":  "labor",    //-intus
	"laborantus":  "labor",    //-antus
	"laborontus":  "labor",    //-ontus
	"laboritus":   "labor",    //-itus
	"laboratus":   "labor",    //-atus
	"laborotus":   "labor",    //-otus
	"laborintu":   "labor",    //-intu
	"laborantu":   "labor",    //-antu
	"laborontu":   "labor",    //-ontu
	"laboritu":    "labor",    //-itu
	"laboratu":    "labor",    //-atu
	"laborotu":    "labor",    //-otu
	"laborinte":   "labor",    //-inte
	"laborante":   "labor",    //-ante
	"laboronte":   "labor",    //-onte
	"laborite":    "labor",    //-ite
	"laborate":    "labor",    //-ate
	"laborote":    "labor",    //-ote
	"laborinta":   "labor",    //-inta
	"laboranta":   "labor",    //-anta
	"laboronta":   "labor",    //-onta
	"laborita":    "labor",    //-ita
	"laborata":    "labor",    //-ata
	"laborota":    "labor",    //-ota
	"laborintaj":  "labor",    //-intaj
	"laborantaj":  "labor",    //-antaj
	"laborontaj":  "labor",    //-ontaj
	"laboritaj":   "labor",    //-itaj
	"laborataj":   "labor",    //-ataj
	"laborotaj":   "labor",    //-otaj
	"laborintan":  "labor",    //-intan
	"laborantan":  "labor",    //-antan
	"laborontan":  "labor",    //-ontan
	"laboritan":   "labor",    //-itan
	"laboratan":   "labor",    //-atan
	"laborotan":   "labor",    //-otan
	"laborintajn": "labor",    //-intajn
	"laborantajn": "labor",    //-antajn
	"laborontajn": "labor",    //-ontajn
	"laboritajn":  "labor",    //-itajn
	"laboratajn":  "labor",    //-atajn
	"laborotajn":  "labor",    //-otajn
	"laboradi":    "labor",    //-adi
	"laboradas":   "labor",    //-adas
	"laborado":    "labor",    //-ado
	"eklabori":    "labor",    //ek-
	"eksciti":     "ekscit",   //ek- exception
	"ellabori":    "labor",    //el-
	"elefanto":    "elefant",  //el- exception
	"laboranto":   "labor",    //-anto
	"laborinto":   "labor",    //-into
	"laboronto":   "labor",    //-onto
	"laborato":    "labor",    //-ato
	"laborito":    "labor",    //-ito
	"laboroto":    "labor",    //-oto
	"esperanto":   "esperant", //esperanto: keep -anto
	"kanto":       "kant",     //kanto: keep -anto
	"hund’":       "hund",     //-’ noun with typographic apostrophe
	"hund'":       "hund",     //-' noun with typewriter apostrophe
	"l’":          "la",       //l’ with typographic apostrophe
	"l'":          "la",       //"l' with typewriter apostrophe
	"un’":         "unu",      //"un’ with typographic apostrophe
	"un'":         "unu",      //"un' with typewriter apostrophe
	"unuj":        "unu",      //plural unu
}

var standAloneRootWords = []string{
	"adiaŭ", "ajn", "al", "almenaŭ", "ambaŭ", "amen", "ankaŭ", "ankoraŭ", "anstataŭ", "antaŭ",
	"apenaŭ", "apud", "aŭ", "baldaŭ", "bis", "boj", "ĉar", "ĉe", "cent", "ĉi", "ĉirkaŭ", "ĉu",
	"da", "de", "dek", "des", "do", "du", "dum", "eĉ", "ekster", "el", "en", "fi", "for", "ĝis",
	"ha", "he", "hieraŭ", "ho", "hodiaŭ", "hu", "hura", "ja", "jam", "je", "jen", "jes", "ju",
	"ĵus", "kaj", "ke", "kontraŭ", "krom", "kun", "kvankam", "kvar", "kvazaŭ", "kvin", "la",
	"laŭ", "malgraŭ", "mem", "mil", "minus", "morgaŭ", "na", "naŭ", "ne", "nek", "nu", "nul",
	"nun", "nur", "ok", "ol", "per", "plej", "pli", "plu", "plus", "po", "por", "post", "preskaŭ",
	"preter", "pri", "pro", "se", "sed", "sen", "sep", "ses", "sub", "super", "sur", "tamen",
	"tra", "trans", "tre", "tri", "tro", "tuj", "unu", "ve",
}

var pronouns = []string{"ci", "ĝi", "ili", "li", "mi", "ni", "oni", "ri", "si", "ŝi", "ŝli", "vi"}
