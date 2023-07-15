package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var input string
	var detskiy bool
	var size int

	flag.StringVar(&input, "i", "text", "Input text file name")
	flag.BoolVar(&detskiy, "d", false, "Detskiy text")
	flag.IntVar(&size, "s", 0, "Font size")
	flag.Parse()

	content, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	items := getItems(string(content))

	if detskiy {
		fmt.Printf(headerDetskiy, size)
	} else {
		fmt.Printf(header, size)
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
				s = append(s, "{\\yf}i")
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
		"krupn{\\yi}mi":                "krup\\-n{\\yi}\\-mi",
		"bezoxibocno":                  "bez\\-oxi\\-boc\\-no",
		"rukokr{\\yi}l{\\y}ami":        "ru\\-ko\\-kr{\\yi}\\-l{\\y}a\\-mi",
		"prezritelynu{\\y}u":           "pre\\-zri\\-tely\\-nu\\-{\\y}u",
		"imenovala":                    "ime\\-no\\-va\\-la",
		"emotionalyn{\\yi}{\\y}":       "emo\\-ti\\-o\\-naly\\-n{\\yi}{\\y}",
		"p{\\ia}timinutku":             "p{\\ia}\\-ti\\-mi\\-nut\\-ku",
		"p{\\ia}timinutki":             "p{\\ia}\\-ti\\-mi\\-nut\\-ki",
		"skolyko":                      "skoly\\-ko",
		"odnajd{\\yi}":                 "od\\-naj\\-d{\\yi}",
		"normalynu{\\y}u":              "nor\\-maly\\-nu\\-{\\y}u",
		"kotor{\\yi}mi":                "ko\\-to\\-r{\\yi}\\-mi",
		"v{\\yi}sokovo":                "v{\\yi}\\-so\\-ko\\-vo",
		"cerv{\\e}m":                   "cer\\-v{\\e}m",
		"podozreva{\\y}em{\\yi}m":      "po\\-do\\-zre\\-va\\-{\\y}e\\-m{\\yi}m",
		"pogovority":                   "po\\-go\\-vo\\-rity",
		"v{\\yi}da{\\y}u{\\x}imisa":    "v{\\yi}\\-da\\-{\\y}u\\-{\\x}i\\-misa",
		"svodn{\\yi}{\\y}":             "svod\\-n{\\yi}{\\y}",
		"Otpravl{\\ia}{\\y}tesy":       "Ot\\-prav\\-l{\\ia}{\\y}\\-tesy",
		"svin{\\c}om":                  "svin\\-{\\c}om",
		"misku":                        "mis\\-ku",
		"po{\\y}avl{\\ia}lsa":          "po\\-{\\y}av\\-l{\\ia}l\\-sa",
		"scita{\\y}et":                 "sci\\-ta\\-{\\y}et",
		"dovolen":                      "do\\-vo\\-len",
		"vdohnovenno":                  "vdoh\\-no\\-ven\\-no",
		"prixlosy":                     "pri\\-xlosy",
		"nesmotr{\\ia}":                "ne\\-smo\\-tr{\\ia}",
		"hoz{\\ia}{\\y}ku":             "ho\\-z{\\ia}{\\y}\\-ku",
		"sv{\\ia}te{\\y}xestvo":        "sv{\\ia}\\-te{\\y}\\-xestvo",
		"sv{\\ia}z{\\ia}mi":            "sv{\\ia}\\-z{\\ia}\\-mi",
		"prekrasno":                    "pre\\-kras\\-no",
		"izurodovann{\\yi}m":           "izu\\-ro\\-do\\-van\\-n{\\yi}m",
		"silyn{\\yi}m":                 "sily\\-n{\\yi}m",
		"mertve{\\c}om":                "mer\\-tve\\-{\\c}om",
		"podobno{\\y}e":                "po\\-dob\\-no\\-{\\y}e",
		"v{\\yi}rezan{\\yi}":           "v{\\yi}\\-re\\-za\\-n{\\yi}",
		"dexovo{\\y}e":                 "de\\-xo\\-vo\\-{\\y}e",
		"prinimal":                     "pri\\-ni\\-mal",
		"hoz{\\ia}{\\y}ina":            "ho\\-z{\\ia}\\-{\\y}i\\-na",
		"malenyki{\\y}":                "ma\\-leny\\-ki{\\y}",
		"Rasskaji":                     "Ras\\-ska\\-ji",
		"t{\\e}mn{\\yi}mi":             "t{\\e}m\\-n{\\yi}\\-mi",
		"protivopolojnu{\\y}u":         "pro\\-ti\\-vo\\-po\\-loj\\-nu\\-{\\y}u",
		"ubl{\\io}dok":                 "ubl{\\io}\\-dok",
		"dogadalisy":                   "do\\-ga\\-da\\-lisy",
		"sposobnost{\\ia}h":            "spo\\-sob\\-no\\-st{\\ia}h",
		"samoubi{\\y}stva":             "samo\\-ubi{\\y}\\-stva",
		"golovo{\\y}":                  "go\\-lo\\-vo{\\y}",
		"Propovednik":                  "Pro\\-po\\-ved\\-nik",
		"ostavxevosa":                  "ostav\\-xe\\-vosa",
		"poslani{\\y}e":                "po\\-sla\\-ni\\-{\\y}e",
		"pogovorim":                    "po\\-go\\-vo\\-rim",
		"soob{\\x}estve":               "so\\-ob\\-{\\x}e\\-stve",
		"v{\\yi}goda":                  "v{\\yi}\\-go\\-da",
		"vrata":                        "vra\\-ta",
		"kotor{\\yi}{\\y}e":            "ko\\-to\\-r{\\yi}\\-{\\y}e",
		"zna{\\y}ex":                   "zna\\-{\\y}ex",
		"nakone{\\c}":                  "na\\-ko\\-ne{\\c}",
		"vlasty":                       "vlas\\-ty",
		"slabocharacternost{\\y}u":     "slabo\\-character\\-no\\-st{\\y}u",
		"na{\\y}omnikov":               "na{\\y}om\\-ni\\-kov",
		"usmehnuvxevosa":               "usmeh\\-nuv\\-xe\\-vosa",
		"skr{\\yi}ty":                  "skr{\\yi}\\-ty",
		"spoko{\\y}no":                 "spo\\-ko{\\y}\\-no",
		"Uverena":                      "Uve\\-re\\-na",
		"predpos{\\yi}lki":             "pred\\-pos{\\yi}l\\-ki",
		"nesto{\\y}komu":               "ne\\-sto{\\y}\\-ko\\-mu",
		"pokrovitelystvom":             "po\\-kro\\-vi\\-tely\\-stvom",
		"isceznoveni{\\y}a":            "iscez\\-no\\-ve\\-ni\\-{\\y}a",
		"rashl{\\e}b{\\yi}vaty":        "ras\\-hl{\\e}\\-b{\\yi}\\-vaty",
		"Miriam":                       "Mi\\-ri\\-am",
		"naroda":                       "na\\-ro\\-da",
		"Vozmojno":                     "Voz\\-moj\\-no",
		"nepri{\\y}atnosti":            "nepri\\-{\\y}at\\-nosti",
		"poznakomilsa":                 "po\\-znako\\-mil\\-sa",
		"neulovim{\\yi}m":              "ne\\-ulo\\-vi\\-m{\\yi}m",
		"vozle":                        "voz\\-le",
		"sv{\\ia}t{\\yi}m":             "sv{\\ia}\\-t{\\yi}m",
		"vsenarodna{\\y}a":             "vse\\-narod\\-na\\-{\\y}a",
		"pere{\\x}egol{\\ia}ty":        "pere\\-{\\x}e\\-go\\-l{\\ia}ty",
		"semides{\\ia}ti":              "semi\\-des{\\ia}\\-ti",
		"vpecatleni{\\y}e":             "vpe\\-catle\\-ni\\-{\\y}e",
		"goroda":                       "go\\-ro\\-da",
		"loxadymi":                     "lo\\-xady\\-mi",
		"doroge":                       "do\\-ro\\-ge",
		"obigraty":                     "ob\\-igraty",
		"blokiroval":                   "bloki\\-ro\\-val",
		"glazami":                      "glaza\\-mi",
		"Hrustalyn{\\yi}h":             "Hrus\\-taly\\-n{\\yi}h",
		"ravnoduxi{\\y}em":             "ravno\\-duxi\\-{\\y}em",
		"stradani{\\y}a":               "stra\\-da\\-ni\\-{\\y}a",
		"pome{\\x}eni{\\y}e":           "po\\-me\\-{\\x}e\\-ni\\-{\\y}e",
		"umol{\\ia}{\\y}u{\\x}e{\\y}e": "umo\\-l{\\ia}\\-{\\y}u\\-{\\x}e\\-{\\y}e",
		"rexivxi{\\y}":                 "re\\-xiv\\-xi{\\y}",
		"zat{\\yi}lke":                 "za\\-t{\\yi}l\\-ke",
		"vseusl{\\yi}xani{\\y}e":       "vse\\-usl{\\yi}\\-xa\\-ni\\-{\\y}e",
		"prokl{\\ia}t{\\yi}h":          "pro\\-kl{\\ia}\\-t{\\yi}h",
		"na{\\y}omnika":                "na\\-{\\y}om\\-ni\\-ka",
		"nevajno":                      "ne\\-vaj\\-no",
		"b{\\yi}va{\\y}et":             "b{\\yi}\\-va\\-{\\y}et",
		"organizovann{\\yi}{\\y}":      "organi\\-zo\\-van\\-n{\\yi}{\\y}",
		"v{\\yi}polnity":               "v{\\yi}\\-pol\\-nity",
		"prot{\\ia}nula":               "pro\\-t{\\ia}\\-nu\\-la",
		"raspolojenn{\\yi}{\\y}":       "raspolojen\\-n{\\yi}{\\y}",
		"cerepicn{\\yi}mi":             "cerepic\\-n{\\yi}\\-mi",
		"neuhojenn{\\yi}{\\y}":         "neuhojen\\-n{\\yi}{\\y}",
		"muz{\\yi}kalynovo":            "muz{\\yi}kaly\\-no\\-vo",
		"govor{\\ia}":                  "go\\-vo\\-r{\\ia}",
		"rabotal":                      "ra\\-bo\\-tal",
		"kozlin{\\yi}mi":               "kozlin{\\yi}\\-mi",
		"unictojen{\\yi}":              "uni\\-cto\\-jen{\\yi}",
		"Realynu{\\y}u":                "Realy\\-nu\\-{\\y}u",
		"dover{\\ia}ty":                "dove\\-r{\\ia}ty",
		"diavolyskih":                  "diavoly\\-skih",
		"ogromn{\\yi}m":                "ogrom\\-n{\\yi}m",
		"sbora":                        "sbo\\-ra",
		"r{\\yi}jevat{\\yi}mi":         "r{\\yi}je\\-va\\-t{\\yi}\\-mi",
		"pripodn{\\ia}lsa":             "pri\\-podn{\\ia}l\\-sa",
		"zvezdcat{\\yi}{\\y}":          "zvezdca\\-t{\\yi}{\\y}",
		"prenebrejitelyn{\\yi}m":       "pre\\-nebreji\\-tely\\-n{\\yi}m",
		"ime{\\y}etsa":                 "ime\\-{\\y}et\\-sa",
		"prot{\\ia}nul":                "pro\\-t{\\ia}\\-nul",
		"ognenn{\\yi}{\\y}":            "ognen\\-n{\\yi}{\\y}",
		"turemn{\\yi}{\\y}":            "turem\\-n{\\yi}{\\y}",
		"pust{\\yi}nnom":               "pu\\-st{\\yi}n\\-nom",
		"necistoplotn{\\yi}m":          "ne\\-cisto\\-plot\\-n{\\yi}m",
		"kollek{\\c}ioner":             "kollek\\-{\\c}ioner",
		"osu{\\x}estvili":              "osu\\-{\\x}estvi\\-li",
		"stalkivalisy":                 "stalki\\-va\\-lisy",
		"blagorodn{\\yi}h":             "blago\\-rod\\-n{\\yi}h",
		"t{\\io}remn{\\yi}{\\y}":       "t{\\io}rem\\-n{\\yi}{\\y}",
		"Mor{\\x}ixsa":                 "Mor\\-{\\x}ixsa",
		"obolocka":                     "obo\\-loc\\-ka",
		"Pocemu":                       "Po\\-ce\\-mu",

		"bezim{\\ia}nn{\\yi}{\\y}":           "bez\\-im{\\ia}n\\-n{\\yi}{\\y}",
		"bezim{\\ia}nna{\\y}a":               "bez\\-im{\\ia}n\\-na{\\y}a",
		"st{\\ia}giva{\\y}u{\\x}i{\\y}e":     "st{\\ia}gi\\-va\\-{\\y}u\\-{\\x}i\\-{\\y}e",
		"v{\\yi}sokopreosv{\\ia}{\\x}enstva": "v{\\yi}soko\\-preosv{\\ia}{\\x}en\\-st\\-va",
	}
}

const header = `\documentclass[%dpt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}

\newcommand{\e}{e}
\newcommand{\yi}{yi}
\newcommand{\ia}{ia}
\newcommand{\io}{io}
\newcommand{\y}{y}
\newcommand{\Y}{Y}

\newcommand{\yf}{y̆}

\newcommand{\X}{X̹}
\newcommand{\x}{x̹}
\newcommand{\C}{C̹}
\renewcommand{\c}{c̹}

`

const headerDetskiy = `\documentclass[%dpt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}

\newcommand{\e}{ë}
\newcommand{\yi}{\mbox{y\hspace{-0.55pt}ı}}
\newcommand{\ia}{\mbox{ı\hspace{-0.55pt}a}}
\newcommand{\io}{\mbox{ı\hspace{-0.55pt}o}}
\newcommand{\y}{y̆}
\newcommand{\Y}{Y̆}

\newcommand{\yf}{y̆}

\newcommand{\X}{X̹}
\newcommand{\x}{x̹}
\newcommand{\C}{C̹}
\renewcommand{\c}{c̹}

`

const footer = `
\end{document}
`
