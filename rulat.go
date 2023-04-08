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

	flag.StringVar(&input, "i", "text", "Input text")
	flag.BoolVar(&detskiy, "d", false, "Detskiy text")
	flag.Parse()

	content, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	items := getItems(string(content))

	if detskiy {
		fmt.Print(headerDetskiy)
	} else {
		fmt.Print(header)
	}

	for _, item := range items {
		s := item2string(item)
		fmt.Print(s)
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

func hythenation() map[string]string {
	return map[string]string{
		"bezoxibocno":             "bez\\-oxi\\-boc\\-no",
		"rukokr{\\yi}l{\\y}ami":   "ru\\-ko\\-kr{\\yi}\\-l{\\y}a\\-mi",
		"prezritelynu{\\y}u":      "pre\\-zri\\-tely\\-nu\\-{\\y}u",
		"imenovala":               "ime\\-no\\-va\\-la",
		"emotionalyn{\\yi}{\\y},": "emo\\-ti\\-o\\-naly\\-n{\\yi}{\\y}",
		"p{\\ia}timinutku:":       "p{\\ia}\\-ti\\-mi\\-nut\\-ku",
		"skolyko":                 "skoly\\-ko",
		"odnajd{\\yi}":            "od\\-naj\\-d{\\yi}",
		"normalynu{\\y}u":         "nor\\-maly\\-nu\\-{\\y}u",
	}
}

func exceptions() map[string]string {
	return map[string]string{
		"аккуратно":           "accuratno",
		"актёр":               "acteur",
		"активнее":            "activne{\\y}e",
		"акцентом":            "accentom",
		"алебардами":          "hallebardami",
		"алебарды":            "hallebard{\\yi}",
		"алхимией":            "alchimi{\\y}e{\\y}",
		"алхимии":             "alchimi{\\y}i",
		"алхимических":        "alchimiceskih",
		"апокалипсис":         "apocalypsis",
		"Аптека":              "Apotheka",
		"аптекарем":           "apothekarem",
		"аптекарь":            "apothekary",
		"Аптекарь":            "Apothekary",
		"аптекаря":            "apothekar{\\ia}",
		"аптеке":              "apotheke",
		"аптекой":             "apotheko{\\y}",
		"аптеку":              "apotheku",
		"артефакт":            "artefact",
		"архангела":           "archangela",
		"атаки":               "attacki",
		"аферу":               "affairu",
		"балкон":              "balcon",
		"банально":            "banalyno",
		"безапелляционно":     "bezappellationno",
		"безрезультатно":      "bezresultatno",
		"бесцеремонно":        "besceremonno",
		"библиотека":          "bibliotheca",
		"библиотеки":          "bibliotheci",
		"блокировал":          "blockieroval",
		"браслет":             "bracelet",
		"бургомистр":          "burgemeester",
		"бургомистра":         "burgemeestera",
		"бутылки":             "butelki",
		"бутылку":             "butelku",
		"бухгалтерской":       "buchhaltersko{\\y}",
		"велефа":              "veletha",
		"воцарился":           "votsarilsa",
		"вулкан":              "vulcan",
		"гвардейцев":          "guarde{\\y}{\\c}ev",
		"гильдии":             "gildi{\\y}i",
		"гимн":                "hymn",
		"гимны":               "hymn{\\yi}",
		"горнистов":           "hornistov",
		"грандиозной":         "grandioseno{\\y}",
		"грандиозную":         "grandiosenu{\\y}u",
		"гримаса":             "grimasa",
		"гримасу":             "grimasu",
		"гримуары":            "grimoir{\\yi}",
		"грифелем":            "griffelem",
		"гротескная":          "grottescna{\\y}a",
		"делегации":           "delegati{\\y}i",
		"депозите":            "deposite",
		"дискредитировать":    "discreditirovaty",
		"дискредитируют":      "discreditiru{\\y}ut",
		"дракона":             "dracona",
		"дукат":               "ducat",
		"дукатов":             "ducatov",
		"дьявол":              "diavol",
		"дьяволопоклонника":   "diavolopoklonnika",
		"дьявольских":         "diavolskih",
		"дьявольское":         "diavolsko{\\y}e",
		"епископ":             "episkop",
		"епископа":            "episkopa",
		"Епископская":         "Episkopska{\\y}a",
		"еретическая":         "hereticeska{\\y}a",
		"жест":                "geste",
		"жесте":               "geste",
		"жестикулируя":        "gesticuliru{\\y}a",
		"зафиксированы":       "zafixirovan{\\yi}",
		"иллюзии":             "illusi{\\y}i",
		"император":           "emperator",
		"императора":          "emperatora",
		"императору":          "emperatoru",
		"Империя":             "Emperi{\\y}a",
		"инициатива":          "initiativa",
		"инквизитор":          "inquisitor",
		"инквизитором":        "inquisitorom",
		"инквизиция":          "inquisiti{\\y}a",
		"Инквизиция":          "Inquisiti{\\y}a",
		"интерьеру":           "interieuru",
		"информации":          "informati{\\y}i",
		"информацию":          "informati{\\y}u",
		"информация":          "informati{\\y}a",
		"истерически":         "hystericeski",
		"истории":             "istori{\\y}i",
		"историков":           "istorikov",
		"историю":             "istori{\\y}u",
		"кабинет":             "cabinet",
		"кавалькада":          "cavalcada",
		"кавалькадой":         "cavalcado{\\y}",
		"камере":              "camere",
		"кандидатура":         "candidatura",
		"канонизация":         "canoniza{\\c}i{\\y}a",
		"каноник":             "canonik",
		"каноника":            "canonika",
		"Кантонец":            "Cantone{\\c}",
		"кантонские":          "cantonski{\\y}e",
		"кантонских":          "cantonskih",
		"капитан":             "capitan",
		"капризов":            "capricov",
		"капюшон":             "capuchon",
		"кардинал":            "cardinal",
		"Кардинал":            "Cardinal",
		"кардинала":           "cardinala",
		"кардиналов":          "cardinalov",
		"кардиналу":           "cardinalu",
		"кардинальская":       "cardinalska{\\y}a",
		"Кардинальская":       "Cardinalska{\\y}a",
		"картограф":           "cartograph",
		"картографом":         "cartographom",
		"катастрофы":          "catastroph{\\yi}",
		"Катастрофы":          "Catastroph{\\yi}",
		"кацеров":             "katzerov",
		"кирасах":             "cuirassah",
		"клан":                "clan",
		"клерк":               "clerk",
		"клиент":              "client",
		"клиентов":            "clientov",
		"климатом":            "climatom",
		"Клирик":              "Cleric",
		"клирика":             "clerica",
		"клирикам":            "clericam",
		"клириками":           "clericami",
		"клириках":            "clericah",
		"клирики":             "clerici",
		"клириков":            "clericov",
		"Клириков":            "Clericov",
		"кобра":               "cobra",
		"коллегой":            "collego{\\y}",
		"коллекционер":        "collectioner",
		"Коллекционер":        "Collectioner",
		"коллекционера":       "collectionera",
		"коллекционеров":      "collectionerov",
		"коллекционером":      "collectionerom",
		"коллекцию":           "collecti{\\y}u",
		"коллекция":           "collecti{\\y}a",
		"колоннаду":           "colonnadu",
		"Команда":             "Commanda",
		"команды":             "command{\\yi}",
		"комбинацией":         "combinati{\\y}e{\\y}",
		"комментариев":        "commentari{\\y}ev",
		"комод":               "commode",
		"компании":            "compani{\\y}i",
		"компания":            "compani{\\y}a",
		"компенсировать":      "compensirovat{\\y}",
		"кондотьер":           "condottier",
		"Кондотьер":           "Condottier",
		"кондотьера":          "condottiera",
		"Конкретизируй":       "Concretiziru{\\y}",
		"конкретики":          "concretiki",
		"конкретно":           "concretno",
		"контакт":             "contact",
		"контексте":           "contexte",
		"континент":           "continent",
		"континента":          "continenta",
		"контраргумент":       "contrargument",
		"контура":             "contoura",
		"конфликт":            "conflict",
		"Конфликты":           "Conflict{\\yi}",
		"концентратом":        "concentratom",
		"копию":               "copi{\\y}u",
		"копия":               "copi{\\y}a",
		"Копия":               "Copi{\\y}a",
		"коридор":             "corridor",
		"корневизору":         "kornevisoru",
		"короной":             "corono{\\y}",
		"кретины":             "cretin{\\yi}",
		"критическую":         "criticesku{\\y}u",
		"курсе":               "kurse",
		"Курьер":              "Courier",
		"курьера":             "couriera",
		"лабиринт":            "labyrinth",
		"логике":              "logice",
		"логику":              "logicu",
		"локальный":           "localyn{\\yi}y",
		"магистрами":          "magisterami",
		"магистров":           "magisterov",
		"Магистры":            "Magister{\\yi}",
		"манере":              "maniere",
		"манускрипты":         "manuscript{\\yi}",
		"маньяк":              "maniac",
		"маньяка":             "maniaca",
		"маньяков":            "maniacov",
		"маразме":             "marasme",
		"маршировала":         "marschierovala",
		"мебель":              "meuble",
		"меланхолично":        "melancholicno",
		"металл":              "metal",
		"металла":             "metalla",
		"механизм":            "mechanism",
		"миф":                 "myth",
		"мифами":              "mythami",
		"мифический":          "mythiceski{\\y}",
		"мифических":          "mythiceskih",
		"монастыре":           "monastere",
		"монастырём":          "monaster{\\e}m",
		"монастырь":           "monastery",
		"монастыря":           "monaster{\\ia}",
		"монах":               "monach",
		"монахов":             "monachov",
		"монаху":              "monachu",
		"монографий":          "monographi{\\y}",
		"музицирования":       "musicirovani{\\y}a",
		"музыка":              "musica",
		"музыкального":        "musicalynovo",
		"Музыкант":            "Musicant",
		"музыканта":           "musicanta",
		"музыкантов":          "musicantov",
		"музыку":              "musicu",
		"мэтр":                "maitre",
		"Мэтр":                "Maitre",
		"наркотик":            "narcotic",
		"наркотика":           "narcotica",
		"нейтральной":         "neutralyno{\\y}",
		"нормальную":          "normalynu{\\y}u",
		"нотацию":             "notati{\\y}u",
		"ориентировку":        "orientirovku",
		"официальная":         "officialyna{\\y}a",
		"паузу":               "pausu",
		"пеликан":             "pelican",
		"пеликана":            "pelicana",
		"пеликаном":           "pelicanom",
		"пеликану":            "pelicanu",
		"пеликанье":           "pelican{\\y}e",
		"персональный":        "personalyn{\\yi}{\\y}",
		"перспектива":         "perspectiva",
		"перспективе":         "perspective",
		"полицейского":        "police{\\y}skovo",
		"полиции":             "polici{\\y}i",
		"Полиции":             "Polici{\\y}i",
		"полицию":             "polici{\\y}u",
		"практически":         "practiceski",
		"процессия":           "processi{\\y}a",
		"Процессия":           "Processi{\\y}a",
		"ратуше":              "rathuse",
		"раухтопазов":         "rauchtopazov",
		"рациям":              "raci{\\y}am",
		"реальная":            "realyna{\\y}a",
		"реальности":          "realynosti",
		"реальность":          "realynosty",
		"Реальную":            "Realynu{\\y}u",
		"регистр":             "register",
		"регулярно":           "regularno",
		"Резервный":           "Reservn{\\yi}{\\y}",
		"результат":           "resultat",
		"реликвии":            "reliqui{\\y}i",
		"реликвию":            "reliqui{\\y}u",
		"репутацией":          "reputati{\\y}e{\\y}",
		"репутация":           "reputati{\\y}a",
		"рефлекторно":         "reflectorno",
		"реформации":          "reformati{\\y}i",
		"рецепт":              "recept",
		"ритмично":            "rythmicno",
		"сапфир":              "sapphir",
		"сапфирами":           "sapphirami",
		"сапфиром":            "sapphirom",
		"секретарь":           "secretary",
		"секретаря":           "secretar{\\ia}",
		"секретов":            "secretov",
		"секреты":             "secret{\\yi}",
		"секта":               "secta",
		"секунд":              "secund",
		"секунду":             "secundu",
		"серафима":            "seraphima",
		"серьёзная":           "serieusena{\\y}a",
		"серьёзно":            "serieuseno",
		"Серьёзно":            "Serieuseno",
		"серьёзного":          "serieusenovo",
		"серьёзности":         "serieusenosti",
		"серьёзные":           "serieusen{\\yi}{\\y}e",
		"серьёзный":           "serieusen{\\yi}{\\y}",
		"силуэт":              "silhouette",
		"силуэта":             "silhouetta",
		"силуэты":             "silhouett{\\yi}",
		"ситуации":            "situati{\\y}i",
		"ситуацию":            "situati{\\y}u",
		"ситуация":            "situati{\\y}a",
		"скандал":             "scandal",
		"скандалов":           "scandalov",
		"скептицизма":         "scepticizma",
		"спецагент":           "specagent",
		"Спецагент":           "Specagent",
		"спецагента":          "specagenta",
		"спецагентов":         "specagentov",
		"специалист":          "specialist",
		"специалисты":         "specialist{\\yi}",
		"специальный":         "specialyn{\\yi}{\\y}",
		"специфической":       "specificesko{\\y}",
		"статистика":          "statistica",
		"статистике":          "statistice",
		"стерильности":        "sterilynosti",
		"субъект":             "subject",
		"суккубов":            "succubov",
		"сутана":              "soutana",
		"сэкономим":           "seconomim",
		"сэкономлю":           "seconoml{\\io}",
		"сюртуке":             "surtuke",
		"тему":                "themu",
		"теологический":       "theologiceski{\\y}",
		"Теоретически":        "Theoreticeski",
		"теории":              "theori{\\y}i",
		"теорию":              "theori{\\y}u",
		"теория":              "theori{\\y}a",
		"Теория":              "Theori{\\y}a",
		"терьер":              "terrier",
		"техники":             "techniki",
		"тифа":                "typha",
		"Тракт":               "Tract",
		"Трактаты":            "Tractat{\\yi}",
		"тракте":              "tracte",
		"трактирах":           "tractirah",
		"тракту":              "tractu",
		"тюремный":            "turemn{\\yi}{\\y}",
		"тюрьмы":              "turm{\\yi}",
		"ультразвук":          "ultrazvuk",
		"фактах":              "factah",
		"факты":               "fact{\\yi}",
		"фальцетом":           "falsettom",
		"фальшивка":           "falsivka",
		"фантазёры":           "fantas{\\e}r{\\yi}",
		"фантазия":            "fantasi{\\y}a",
		"формальность":        "formalynosty",
		"фраза":               "phrasa",
		"фразу":               "phrasu",
		"хаос":                "chaos",
		"хаотичных":           "chaoticn{\\yi}h",
		"характер":            "character",
		"характерную":         "characternu{\\y}u",
		"характерных":         "charactern{\\yi}h",
		"слабохарактерностью": "slabocharacternost{\\y}u",
		"Царём":               "Tsar{\\e}m",
		"Царь":                "Tsar",
		"центр":               "center",
		"центральной":         "centralyno{\\y}",
		"центральную":         "centralynu{\\y}u",
		"центральных":         "centralyn{\\yi}h",
		"центре":              "centre",
		"центру":              "centru",
		"шанс":                "chance",
		"Шанс":                "Chance",
		"шапке":               "chapke",
		"шапку":               "chapku",
		"шифр":                "chiffre",
		"школа":               "schola",
		"шоке":                "schoke",
		"шпиона":              "spiona",
		"шпионов":             "spionov",
		"экземпляром":         "exemplarom",
		"экспедиций":          "expediti{\\y}",
		"эмоции":              "emoti{\\y}i",
		"эмоций":              "emoti{\\y}",
		"эмоциональный":       "emotionalyn{\\yi}{\\y}",
		"эмоциям":             "emoti{\\y}am",
		"энтузиазма":          "enthusiasma",
		"этаж":                "etage",
		"этажа":               "etaga",
		"этаже":               "etage",
		"эффект":              "effect",
		"эхолокаторы":         "echolocator{\\yi}",
		"эхом":                "echom",

		"беднягу-аптекаря":        "bedn{\\ia}gu-apothekar{\\ia}",
		"гвардейцев-альбаландцев": "guarde{\\y}{\\c}ev-albaland{\\c}ev",
		"женщину-музыканта":       "jen{\\x}inu-musicanta",
		"клирики-тупицы":          "clerici-tupi{\\c}i",
		"монах-привратник":        "monach-privratnik",
		"нехристю-хагжиту":        "nechrist{\\io}-hagjitu",
		"по-нарарски":             "po-nararrski",
		"Стриж-парикмахер":        "Strij-perukmacher",
		"Стрижу-парикмахеру":      "Striju-perukmacheru",
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
		"каликвец":       "callicve{\\c}",
		"Каликвец":       "Callicve{\\c}",
		"каликвеца":      "callicve{\\c}a",
		"каликвецев":     "callicve{\\c}ev",
		"Каликвецы":      "Callicve{\\c}i",
		"Клеменз":        "Clemence",
		"Константин":     "Constantin",
		"Константина":    "Constantina",
		"Константину":    "Constantinu",
		"Крусо":          "Caruso",
		"Лавендуззского": "Lavendusskovo",
		"Лезерберг":      "Leserberg",
		"Лезерберга":     "Leserberga",
		"Ливетте":        "Levette",
		"Лисецк":         "Lesetsk",
		"Лисецке":        "Lesetske",
		"Людвиг":         "Ludwig",
		"Микель":         "Michel",
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
		"Воистину":      "Voistinu",
		"много":         "mnogo",
		"наивен":        "naiven",
		"намного":       "namnogo",
		"Неинтересно":   "Neiteresno",
		"немного":       "nemnogo",
		"ненамного":     "nenamnogo",
		"почувствовал":  "pocustvoval",
		"почувствовать": "pocustvovaty",
		"предыдущая":    "predidu{\\x}a{\\y}a",
		"предыдущего":   "predidu{\\x}evo",
		"Предыстория":   "Predistori{\\y}a",
		"сегодня":       "sevodn{\\ia}",
		"Сегодня":       "Sevodn{\\ia}",
		"собиравшимися": "sobiravximisa",
		"стоит":         "stoit",
		"удостоил":      "udostoil",
		"чувствовал":    "custvoval",
		"чувствовала":   "custvovala",
		"чувствуя":      "custvu{\\y}a",
	}
}

const header = `\documentclass[10pt]{book}
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

const headerDetskiy = `\documentclass[12pt]{book}
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
