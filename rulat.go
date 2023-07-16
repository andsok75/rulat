package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var inputName string
	var simplified bool
	var fontSize int

	flag.StringVar(&inputName, "i", "-", "Input file name")
	flag.BoolVar(&simplified, "s", false, "Use simplified characters")
	flag.IntVar(&fontSize, "f", 12, "Font size")
	flag.Parse()

	var err error
	var input io.Reader = os.Stdin
	if inputName != "-" {
		input, err = os.Open(inputName)
		if err != nil {
			log.Fatal(err)
		}
	}
	content, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	items := getItems(string(content))

	fmt.Printf(header, fontSize)
	fmt.Print(chars)
	if simplified {
		fmt.Print(charsSimplified)
	}

	for _, item := range items {
		s := item2string(item)
		if h, ok := hyphenation()[s]; ok {
			fmt.Print(h)
		} else {
			fmt.Print(s)
		}
	}

	fmt.Print(footer)
}

func item2string(item Item) string {
	if !item.isWord {
		switch item.content {
		case "«":
			return "``"
		case "»":
			return "\""
		default:
			return item.content
		}
	}
	if v, ok := exceptions()[item.content]; ok {
		return v
	}

	word := []rune(item.content)
	return word2string(word)
}

func word2string(word []rune) string {
	s := []string{}
	wl := len(word)
	p := ""
	n := ""
	for i := 0; i < wl; i++ {
		c := string(word[i])
		if i > 0 {
			p = string(word[i-1])
		} else {
			p = ""
		}
		if i < wl-1 {
			n = string(word[i+1])
		} else {
			n = ""
		}
		switch c {
		case "-":
			s = append(s, "-")
		case "А":
			s = append(s, "A")
		case "Б":
			s = append(s, "B")
		case "В":
			s = append(s, "V")
		case "Г":
			s = append(s, "G")
		case "Д":
			s = append(s, "D")
		case "Е":
			s = append(s, "{\\Y}e")
		case "Ё":
			s = append(s, "{\\Y}o")
		case "Ж":
			s = append(s, "J")
		case "З":
			s = append(s, "Z")
		case "И":
			s = append(s, "I")
		case "Й":
			s = append(s, "{\\Y}")
		case "К":
			s = append(s, "K")
		case "Л":
			s = append(s, "L")
		case "М":
			s = append(s, "M")
		case "Н":
			s = append(s, "N")
		case "О":
			s = append(s, "O")
		case "П":
			s = append(s, "P")
		case "Р":
			s = append(s, "R")
		case "С":
			s = append(s, "S")
		case "Т":
			s = append(s, "T")
		case "У":
			s = append(s, "U")
		case "Ф":
			s = append(s, "F")
		case "Х":
			s = append(s, "H")
		case "Ц":
			s = append(s, "{\\C}")
		case "Ч":
			s = append(s, "C")
		case "Ш":
			s = append(s, "X")
		case "Щ":
			s = append(s, "{\\X}")
		case "Ъ":
			s = append(s, "Y")
		case "Ы":
			s = append(s, "YI")
		case "Ь":
			s = append(s, "Y")
		case "Э":
			s = append(s, "E")
		case "Ю":
			s = append(s, "{\\Y}u")
		case "Я":
			s = append(s, "{\\Y}a")
		case "а":
			s = append(s, "a")
		case "б":
			s = append(s, "b")
		case "в":
			s = append(s, "v")
		case "г":
			if p == "о" || p == "е" || p == "Е" {
				if i == wl-2 && string(word[i:]) == "го" {
					s = append(s, "vo")
					i += 1
				} else if i == wl-4 && string(word[i:]) == "гося" {
					s = append(s, "vosa")
					i += 3
				} else if i < wl-4 && string(word[i:i+3]) == "го-" {
					s = append(s, "vo-")
					i += 2
				} else {
					s = append(s, "g")
				}
			} else {
				s = append(s, "g")
			}
		case "д":
			s = append(s, "d")
		case "е":
			if p == "" || isVowel(p) {
				s = append(s, "{\\y}e")
			} else {
				s = append(s, "e")
			}
		case "ё":
			if p == "" || isVowel(p) {
				s = append(s, "{\\y}o")
			} else {
				if isFrict(p) {
					s = append(s, "o")
				} else {
					s = append(s, "{\\e}")
				}
			}
		case "ж":
			s = append(s, "j")
		case "з":
			s = append(s, "z")
		case "и":
			if isVowel(p) {
				if i == 3 && isPrefix(string(word[:3])) {
					s = append(s, "i")
				} else if i == 2 && isPrefix(string(word[:2])) {
					s = append(s, "i")
				} else {
					s = append(s, "{\\y}i")
				}
			} else {
				s = append(s, "i")
			}
		case "й":
			s = append(s, "{\\y}")
		case "к":
			s = append(s, "k")
		case "л":
			s = append(s, "l")
		case "м":
			s = append(s, "m")
		case "н":
			s = append(s, "n")
		case "о":
			s = append(s, "o")
		case "п":
			s = append(s, "p")
		case "р":
			s = append(s, "r")
		case "с":
			if i == wl-2 && string(word[i:]) == "ся" && (p == "у" || p == "ю" || p == "й" || p == "и" || p == "я" || p == "е" || p == "л" || p == "м" || p == "т" || p == "б" || p == "ь" || p == "х" || p == "г") {
				s = append(s, "sa")
				i += 1
			} else {
				s = append(s, "s")
			}
		case "т":
			s = append(s, "t")
		case "у":
			s = append(s, "u")
		case "ф":
			s = append(s, "f")
		case "х":
			s = append(s, "h")
		case "ц":
			s = append(s, "{\\c}")
		case "ч":
			s = append(s, "c")
		case "ш":
			s = append(s, "x")
		case "щ":
			s = append(s, "{\\x}")
		case "ы":
			if isFrict(p) {
				s = append(s, "i")
			} else {
				s = append(s, "{\\yi}")
			}
		case "ъ", "ь":
			if n == "ю" {
				s = append(s, "{\\y}u")
				i += 1
			} else if n == "я" {
				s = append(s, "{\\y}a")
				i += 1
			} else if n == "ё" {
				s = append(s, "{\\y}o")
				i += 1
			} else if n == "е" {
				s = append(s, "{\\y}e")
				i += 1
			} else if n == "и" {
				s = append(s, "{\\y}i")
				i += 1
			} else if i == wl-3 && string(word[i+1:]) == "ся" && (p == "т" || p == "ш") {
				// skip
			} else if i == wl-1 && isFrict(p) {
				// skip
			} else {
				s = append(s, "y")
			}
		case "э":
			s = append(s, "e")
		case "ю":
			if p == "" || isVowel(p) {
				s = append(s, "{\\y}u")
			} else {
				s = append(s, "{\\io}")
			}
		case "я":
			if p == "" || isVowel(p) {
				s = append(s, "{\\y}a")
			} else {
				s = append(s, "{\\ia}")
			}
		default:
			log.Fatalf("not valid: *%s*", c)
		}
	}
	return strings.Join(s, "")
}

func isVowel(c string) bool {
	if c == "а" || c == "е" || c == "ё" || c == "и" || c == "о" || c == "у" || c == "ы" || c == "э" || c == "ю" || c == "я" ||
		c == "А" || c == "Е" || c == "Ё" || c == "И" || c == "О" || c == "У" || c == "Ы" || c == "Э" || c == "Ю" || c == "Я" {
		return true
	} else {
		return false
	}
}

func isWord(c string) bool {
	switch c {
	case "-":
		fallthrough
	case "А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я":
		fallthrough
	case "а", "б", "в", "г", "д", "е", "ё", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ъ", "ы", "ь", "э", "ю", "я":
		return true
	default:
		return false
	}
}

func isFrict(c string) bool {
	switch c {
	case "Ж", "Ц", "Ч", "Ш", "Щ":
		fallthrough
	case "ж", "ц", "ч", "ш", "щ":
		return true
	default:
		return false
	}
}

type Item struct {
	isWord  bool
	content string
}

func getItems(s string) []Item {
	items := []Item{}
	word := ""
	for _, r := range s {
		c := string(r)
		if isWord(c) {
			word = word + c
			continue
		}
		if word != "" {
			items = append(items, Item{true, word})
			word = ""
		}
		items = append(items, Item{false, c})
	}
	return items
}

func isPrefix(s string) bool {
	return s == "про" || s == "по" || s == "за" || s == "на" || s == "не" || s == "Пред"
}

func exceptions() map[string]string {
	return map[string]string{
		"алхимией":            "alchimi{\\y}e{\\y}",
		"алхимии":             "alchimi{\\y}i",
		"алхимических":        "alchimiceskih",
		"архангела":           "archangela",
		"бухгалтерской":       "buchhaltersko{\\y}",
		"воцарился":           "votsarilsa",
		"гвардейцев":          "guarde{\\y}{\\c}ev",
		"гимн":                "hymn",
		"гимны":               "hymn{\\yi}",
		"дьявол":              "diavol",
		"дьяволопоклонника":   "diavolopoklonnika",
		"дьявольских":         "diavolyskih",
		"дьявольское":         "diavolysko{\\y}e",
		"епископ":             "episkop",
		"епископа":            "episkopa",
		"Епископская":         "Episkopska{\\y}a",
		"еретическая":         "hereticeska{\\y}a",
		"император":           "emperator",
		"императора":          "emperatora",
		"императору":          "emperatoru",
		"Империя":             "Emperi{\\y}a",
		"интерьеру":           "interieru",
		"кацеров":             "katzerov",
		"клиент":              "klient",
		"клиентов":            "klientov",
		"Клирик":              "Klerik",
		"клирика":             "klerika",
		"клирикам":            "klerikam",
		"клириками":           "klerikami",
		"клириках":            "klerikah",
		"клирики":             "kleriki",
		"клириков":            "klerikov",
		"Клириков":            "Klerikov",
		"кондотьер":           "kondotier",
		"Кондотьер":           "Kondotier",
		"кондотьера":          "kondotiera",
		"Курьер":              "Kurier",
		"курьера":             "kuriera",
		"меланхолично":        "melancholicno",
		"металл":              "metal",
		"металла":             "metalla",
		"механизм":            "mechanizm",
		"монах":               "monach",
		"монахов":             "monachov",
		"монаху":              "monachu",
		"мэтр":                "maître",
		"Мэтр":                "Maître",
		"ориентировку":        "orientirovku",
		"раухтопазов":         "rauchtopazov",
		"терьер":              "terrier",
		"техники":             "techniki",
		"ультразвук":          "ultrazvuk",
		"хаос":                "chaos",
		"хаотичных":           "chaoticn{\\yi}h",
		"характер":            "character",
		"характерную":         "characternu{\\y}u",
		"характерных":         "charactern{\\yi}h",
		"слабохарактерностью": "slabocharacternost{\\y}u",
		"Царём":               "Tsar{\\e}m",
		"Царь":                "Tsar",
		"школа":               "schola",
		"эхолокаторы":         "echolokator{\\yi}",
		"эхом":                "echom",

		"гвардейцев-альбаландцев": "guarde{\\y}{\\c}ev-albaland{\\c}ev",
		"клирики-тупицы":          "kleriki-tupi{\\c}i",
		"монах-привратник":        "monach-privratnik",
		"нехристю-хагжиту":        "nechrist{\\io}-hagjitu",
		"по-нарарски":             "po-nararrski",
		"царей-зверей":            "tsare{\\y}-zvere{\\y}",

		"Адиль":          "Adil",
		"Адиля":          "Adila",
		"Александр":      "Alexander",
		"Александра":     "Alexandera",
		"аль":            "al",
		"Альбаланд":      "Albaland",
		"Альбаланда":     "Albalanda",
		"альбаландец":    "albalande{\\c}",
		"альбаландских":  "albalandskih",
		"альбаландцев":   "albaland{\\c}ev",
		"Барбурге":       "Warburge",
		"Бробергер":      "Wroberger",
		"Бробергера":     "Wrobergera",
		"Валентин":       "Valentin",
		"Валентина":      "Valentina",
		"Вальтер":        "Walter",
		"Вальтера":       "Waltera",
		"Вальтером":      "Walterom",
		"Вальтеру":       "Walteru",
		"Витильск":       "Witilsk",
		"Витильска":      "Witilska",
		"Витильске":      "Witilske",
		"Вьюна":          "Vewna",
		"Вьюну":          "Vewnu",
		"Ганс":           "Hans",
		"Ганса":          "Hansa",
		"Гансом":         "Hansom",
		"Дель":           "Del",
		"Джулию":         "Juliyu",
		"Джума":          "Juma",
		"Дорч-ган-Тойн":  "Dortch-gan-Toyn",
		"Дорч-ган-Тойне": "Dortch-gan-Toyne",
		"Иисусе":         "Iesuse",
		"Каварзере":      "Cavarsere",
		"каликвец":       "kallikve{\\c}",
		"Каликвец":       "Kallikve{\\c}",
		"каликвеца":      "kallikve{\\c}a",
		"каликвецев":     "kallikve{\\c}ev",
		"Каликвецы":      "Kallikve{\\c}i",
		"Клеменз":        "Clemence",
		"Крусо":          "Caruso",
		"Лавендуззского": "Lavendusskovo",
		"Лезерберг":      "Leserberg",
		"Лезерберга":     "Leserberga",
		"Ливетте":        "Levette",
		"Лисецк":         "Lesetsk",
		"Лисецке":        "Lesetske",
		"Людвиг":         "Ludwig",
		"Михаила":        "Michaela",
		"Нарара":         "Nararra",
		"Нарары":         "Nararr{\\yi}",
		"Натан":          "Nathan",
		"Натана":         "Nathana",
		"нехристю":       "nechrist{\\io}",
		"Нормайенн":      "Normaenn",
		"Нормайенна":     "Normaenna",
		"Нормайенном":    "Normaennom",
		"Ньюгорт":        "Newgord",
		"Петра":          "Petera",
		"Прогансу":       "Progance",
		"Рози":           "Rosi",
		"Селико":         "Selico",
		"Травинно":       "Travino",
		"Уг":             "Ug",
		"Фабьен":         "Fabien",
		"Филипп":         "Philipp",
		"Филиппом":       "Philippom",
		"Филиппу":        "Philippu",
		"Фирвальден":     "Firvalden",
		"Фирвальдене":    "Firvaldene",
		"флотолийских":   "flottoli{\\y}skih",
		"Фрингбоу":       "Fringboe",
		"Ханна":          "Hannah",
		"Хартвиг":        "Hartwig",
		"Хартвига":       "Hartwiga",
		"Христа":         "Christa",
		"Христе":         "Christe",
		"Христово":       "Christovo",
		"Христовой":      "Christovo{\\y}",
		"Чак":            "Chuck",
		"Чака":           "Chucka",
		"Чезаре":         "Cesare",
		"Шоссии":         "Schossi{\\y}i",
		"Шуко":           "Schuco",

		"аист":          "aist",
		"безымянный":    "bezim{\\ia}nn{\\yi}{\\y}",
		"безымянная":    "bezim{\\ia}nna{\\y}a",
		"Воистину":      "Voistinu",
		"много":         "mnogo",
		"наивен":        "naiven",
		"намного":       "namnogo",
		"Неинтересно":   "Neiteresno",
		"немного":       "nemnogo",
		"ненамного":     "nenamnogo",
		"обыграть":      "obigraty",
		"почувствовал":  "pocustvoval",
		"почувствовать": "pocustvovaty",
		"предыдущая":    "predidu{\\x}a{\\y}a",
		"предыдущего":   "predidu{\\x}evo",
		"Предыстория":   "Predistori{\\y}a",
		"сегодня":       "sevodn{\\ia}",
		"Сегодня":       "Sevodn{\\ia}",
		"собиравшимися": "sobiravximisa",
		"стоит":         "stoit",
		"стоило":        "stoilo",
		"Стоило":        "Stoilo",
		"стоила":        "stoila",
		"удостоил":      "udostoil",
		"чувствовал":    "custvoval",
		"чувствовала":   "custvovala",
		"чувствуя":      "custvu{\\y}a",
	}
}

func hyphenation() map[string]string {
	return map[string]string{
		"bezim{\\ia}nna{\\y}a":      "bez\\-im{\\ia}n\\-na{\\y}a",
		"bezim{\\ia}nn{\\yi}{\\y}":  "bez\\-im{\\ia}n\\-n{\\yi}{\\y}",
		"bezoxibocno":               "bez\\-oxi\\-boc\\-no",
		"blagorodn{\\yi}h":          "blago\\-rod\\-n{\\yi}h",
		"blokiroval":                "bloki\\-ro\\-val",
		"b{\\yi}va{\\y}et":          "b{\\yi}\\-va\\-{\\y}et",
		"cerepicn{\\yi}mi":          "cerepic\\-n{\\yi}\\-mi",
		"cerv{\\e}m":                "cer\\-v{\\e}m",
		"dexovo{\\y}e":              "de\\-xo\\-vo\\-{\\y}e",
		"diavolyskih":               "diavoly\\-skih",
		"dogadalisy":                "do\\-ga\\-da\\-lisy",
		"doroge":                    "do\\-ro\\-ge",
		"dover{\\ia}ty":             "dove\\-r{\\ia}ty",
		"dovolen":                   "do\\-vo\\-len",
		"glazami":                   "glaza\\-mi",
		"golovo{\\y}":               "go\\-lo\\-vo{\\y}",
		"goroda":                    "go\\-ro\\-da",
		"govor{\\ia}":               "go\\-vo\\-r{\\ia}",
		"hoz{\\ia}{\\y}ina":         "ho\\-z{\\ia}\\-{\\y}i\\-na",
		"hoz{\\ia}{\\y}ku":          "ho\\-z{\\ia}{\\y}\\-ku",
		"Hrustalyn{\\yi}h":          "Hrus\\-taly\\-n{\\yi}h",
		"imenovala":                 "ime\\-no\\-va\\-la",
		"ime{\\y}etsa":              "ime\\-{\\y}et\\-sa",
		"isceznoveni{\\y}a":         "iscez\\-no\\-ve\\-ni\\-{\\y}a",
		"izurodovann{\\yi}m":        "izu\\-ro\\-do\\-van\\-n{\\yi}m",
		"kollek{\\c}ioner":          "kollek\\-{\\c}ioner",
		"kotor{\\yi}mi":             "ko\\-to\\-r{\\yi}\\-mi",
		"kotor{\\yi}{\\y}e":         "ko\\-to\\-r{\\yi}\\-{\\y}e",
		"kozlin{\\yi}mi":            "kozlin{\\yi}\\-mi",
		"krupn{\\yi}mi":             "krup\\-n{\\yi}\\-mi",
		"loxadymi":                  "lo\\-xady\\-mi",
		"malenyki{\\y}":             "ma\\-leny\\-ki{\\y}",
		"mertve{\\c}om":             "mer\\-tve\\-{\\c}om",
		"Miriam":                    "Mi\\-ri\\-am",
		"misku":                     "mis\\-ku",
		"Mor{\\x}ixsa":              "Mor\\-{\\x}ixsa",
		"muz{\\yi}kalynovo":         "muz{\\yi}kaly\\-no\\-vo",
		"nakone{\\c}":               "na\\-ko\\-ne{\\c}",
		"naroda":                    "na\\-ro\\-da",
		"na{\\y}omnika":             "na\\-{\\y}om\\-ni\\-ka",
		"na{\\y}omnikov":            "na{\\y}om\\-ni\\-kov",
		"necistoplotn{\\yi}m":       "ne\\-cisto\\-plot\\-n{\\yi}m",
		"nepri{\\y}atnosti":         "nepri\\-{\\y}at\\-nosti",
		"nesmotr{\\ia}":             "ne\\-smo\\-tr{\\ia}",
		"nesto{\\y}komu":            "ne\\-sto{\\y}\\-ko\\-mu",
		"neuhojenn{\\yi}{\\y}":      "neuhojen\\-n{\\yi}{\\y}",
		"neulovim{\\yi}m":           "ne\\-ulo\\-vi\\-m{\\yi}m",
		"nevajno":                   "ne\\-vaj\\-no",
		"normalynu{\\y}u":           "nor\\-maly\\-nu\\-{\\y}u",
		"obigraty":                  "ob\\-igraty",
		"obolocka":                  "obo\\-loc\\-ka",
		"odnajd{\\yi}":              "od\\-naj\\-d{\\yi}",
		"ognenn{\\yi}{\\y}":         "ognen\\-n{\\yi}{\\y}",
		"ogromn{\\yi}m":             "ogrom\\-n{\\yi}m",
		"organizovann{\\yi}{\\y}":   "organi\\-zo\\-van\\-n{\\yi}{\\y}",
		"ostavxevosa":               "ostav\\-xe\\-vosa",
		"osu{\\x}estvili":           "osu\\-{\\x}estvi\\-li",
		"Otpravl{\\ia}{\\y}tesy":    "Ot\\-prav\\-l{\\ia}{\\y}\\-tesy",
		"pere{\\x}egol{\\ia}ty":     "pere\\-{\\x}e\\-go\\-l{\\ia}ty",
		"p{\\ia}timinutki":          "p{\\ia}\\-ti\\-mi\\-nut\\-ki",
		"p{\\ia}timinutku":          "p{\\ia}\\-ti\\-mi\\-nut\\-ku",
		"Pocemu":                    "Po\\-ce\\-mu",
		"podobno{\\y}e":             "po\\-dob\\-no\\-{\\y}e",
		"podozreva{\\y}em{\\yi}m":   "po\\-do\\-zre\\-va\\-{\\y}e\\-m{\\yi}m",
		"pogovorim":                 "po\\-go\\-vo\\-rim",
		"pogovority":                "po\\-go\\-vo\\-rity",
		"pokrovitelystvom":          "po\\-kro\\-vi\\-tely\\-stvom",
		"pome{\\x}eni{\\y}e":        "po\\-me\\-{\\x}e\\-ni\\-{\\y}e",
		"poslani{\\y}e":             "po\\-sla\\-ni\\-{\\y}e",
		"po{\\y}avl{\\ia}lsa":       "po\\-{\\y}av\\-l{\\ia}l\\-sa",
		"poznakomilsa":              "po\\-znako\\-mil\\-sa",
		"predpos{\\yi}lki":          "pred\\-pos{\\yi}l\\-ki",
		"prekrasno":                 "pre\\-kras\\-no",
		"prenebrejitelyn{\\yi}m":    "pre\\-nebreji\\-tely\\-n{\\yi}m",
		"prezritelynu{\\y}u":        "pre\\-zri\\-tely\\-nu\\-{\\y}u",
		"prinimal":                  "pri\\-ni\\-mal",
		"pripodn{\\ia}lsa":          "pri\\-podn{\\ia}l\\-sa",
		"prixlosy":                  "pri\\-xlosy",
		"prokl{\\ia}t{\\yi}h":       "pro\\-kl{\\ia}\\-t{\\yi}h",
		"Propovednik":               "Pro\\-po\\-ved\\-nik",
		"prot{\\ia}nula":            "pro\\-t{\\ia}\\-nu\\-la",
		"prot{\\ia}nul":             "pro\\-t{\\ia}\\-nul",
		"protivopolojnu{\\y}u":      "pro\\-ti\\-vo\\-po\\-loj\\-nu\\-{\\y}u",
		"pust{\\yi}nnom":            "pu\\-st{\\yi}n\\-nom",
		"rabotal":                   "ra\\-bo\\-tal",
		"rashl{\\e}b{\\yi}vaty":     "ras\\-hl{\\e}\\-b{\\yi}\\-vaty",
		"raspolojenn{\\yi}{\\y}":    "raspolojen\\-n{\\yi}{\\y}",
		"Rasskaji":                  "Ras\\-ska\\-ji",
		"ravnoduxi{\\y}em":          "ravno\\-duxi\\-{\\y}em",
		"Realynu{\\y}u":             "Realy\\-nu\\-{\\y}u",
		"rexivxi{\\y}":              "re\\-xiv\\-xi{\\y}",
		"rukokr{\\yi}l{\\y}ami":     "ru\\-ko\\-kr{\\yi}\\-l{\\y}a\\-mi",
		"r{\\yi}jevat{\\yi}mi":      "r{\\yi}je\\-va\\-t{\\yi}\\-mi",
		"samoubi{\\y}stva":          "samo\\-ubi{\\y}\\-stva",
		"sbora":                     "sbo\\-ra",
		"scita{\\y}et":              "sci\\-ta\\-{\\y}et",
		"semides{\\ia}ti":           "semi\\-des{\\ia}\\-ti",
		"silyn{\\yi}m":              "sily\\-n{\\yi}m",
		"skolyko":                   "skoly\\-ko",
		"skr{\\yi}ty":               "skr{\\yi}\\-ty",
		"slabocharacternost{\\y}u":  "slabo\\-character\\-no\\-st{\\y}u",
		"soob{\\x}estve":            "so\\-ob\\-{\\x}e\\-stve",
		"spoko{\\y}no":              "spo\\-ko{\\y}\\-no",
		"sposobnost{\\ia}h":         "spo\\-sob\\-no\\-st{\\ia}h",
		"stalkivalisy":              "stalki\\-va\\-lisy",
		"stradani{\\y}a":            "stra\\-da\\-ni\\-{\\y}a",
		"sv{\\ia}te{\\y}xestvo":     "sv{\\ia}\\-te{\\y}\\-xestvo",
		"sv{\\ia}t{\\yi}m":          "sv{\\ia}\\-t{\\yi}m",
		"sv{\\ia}z{\\ia}mi":         "sv{\\ia}\\-z{\\ia}\\-mi",
		"svin{\\c}om":               "svin\\-{\\c}om",
		"svodn{\\yi}{\\y}":          "svod\\-n{\\yi}{\\y}",
		"t{\\e}mn{\\yi}mi":          "t{\\e}m\\-n{\\yi}\\-mi",
		"t{\\io}remn{\\yi}{\\y}":    "t{\\io}rem\\-n{\\yi}{\\y}",
		"ubl{\\io}dok":              "ubl{\\io}\\-dok",
		"unictojen{\\yi}":           "uni\\-cto\\-jen{\\yi}",
		"usmehnuvxevosa":            "usmeh\\-nuv\\-xe\\-vosa",
		"Uverena":                   "Uve\\-re\\-na",
		"vdohnovenno":               "vdoh\\-no\\-ven\\-no",
		"vlasty":                    "vlas\\-ty",
		"vozle":                     "voz\\-le",
		"Vozmojno":                  "Voz\\-moj\\-no",
		"vpecatleni{\\y}e":          "vpe\\-catle\\-ni\\-{\\y}e",
		"vrata":                     "vra\\-ta",
		"vsenarodna{\\y}a":          "vse\\-narod\\-na\\-{\\y}a",
		"vseusl{\\yi}xani{\\y}e":    "vse\\-usl{\\yi}\\-xa\\-ni\\-{\\y}e",
		"v{\\yi}da{\\y}u{\\x}imisa": "v{\\yi}\\-da\\-{\\y}u\\-{\\x}i\\-misa",
		"v{\\yi}goda":               "v{\\yi}\\-go\\-da",
		"v{\\yi}polnity":            "v{\\yi}\\-pol\\-nity",
		"v{\\yi}rezan{\\yi}":        "v{\\yi}\\-re\\-za\\-n{\\yi}",
		"v{\\yi}sokovo":             "v{\\yi}\\-so\\-ko\\-vo",
		"zat{\\yi}lke":              "za\\-t{\\yi}l\\-ke",
		"zna{\\y}ex":                "zna\\-{\\y}ex",
		"zvezdcat{\\yi}{\\y}":       "zvezdca\\-t{\\yi}{\\y}",

		"umol{\\ia}{\\y}u{\\x}e{\\y}e":       "umo\\-l{\\ia}\\-{\\y}u\\-{\\x}e\\-{\\y}e",
		"st{\\ia}giva{\\y}u{\\x}i{\\y}e":     "st{\\ia}gi\\-va\\-{\\y}u\\-{\\x}i\\-{\\y}e",
		"v{\\yi}sokopreosv{\\ia}{\\x}enstva": "v{\\yi}soko\\-preosv{\\ia}{\\x}en\\-st\\-va",
	}
}

const header = `\documentclass[%dpt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}

`

const chars = `\newcommand{\X}{X̹}
\newcommand{\x}{x̹}
\newcommand{\C}{C̹}
\renewcommand{\c}{c̹}

\newcommand{\e}{ë}
\newcommand{\yi}{\mbox{y\hspace{-0.55pt}ı}}
\newcommand{\ia}{\mbox{ı\hspace{-0.55pt}a}}
\newcommand{\io}{\mbox{ı\hspace{-0.55pt}o}}
\newcommand{\y}{y̆}
\newcommand{\Y}{Y̆}

`

const charsSimplified = `\renewcommand{\e}{e}
\renewcommand{\yi}{yi}
\renewcommand{\ia}{ia}
\renewcommand{\io}{io}
\renewcommand{\y}{y}
\renewcommand{\Y}{Y}

`

const footer = `
\end{document}
`
