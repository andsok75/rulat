package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const input_file = "./detskiy"

func main() {
	content, err := ioutil.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	items := getItems(string(content))

	fmt.Print(header)
	fmt.Print(commands)

	for _, item := range items {
		if !item.isWord {
			switch item.content {
			case "«":
				fmt.Print("``")
			case "»":
				fmt.Print("\"")
			default:
				fmt.Print(item.content)
			}
			continue
		}
		if v, ok := exceptions()[item.content]; ok {
			fmt.Print(v)
			continue
		}

		log.Printf("processing *%s*", item.content)
		word := []rune(item.content)
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
				fmt.Print("-")
			case "А":
				fmt.Print("A")
			case "Б":
				fmt.Print("B")
			case "В":
				fmt.Print("V")
			case "Г":
				fmt.Print("G")
			case "Д":
				fmt.Print("D")
			case "Е":
				fmt.Print("{\\Y}e")
			case "Ё":
				fmt.Print("{\\Y}o")
			case "Ж":
				fmt.Print("J")
			case "З":
				fmt.Print("Z")
			case "И":
				fmt.Print("I")
			case "Й":
				fmt.Print("{\\Y}")
			case "К":
				fmt.Print("K")
			case "Л":
				fmt.Print("L")
			case "М":
				fmt.Print("M")
			case "Н":
				fmt.Print("N")
			case "О":
				fmt.Print("O")
			case "П":
				fmt.Print("P")
			case "Р":
				fmt.Print("R")
			case "С":
				fmt.Print("S")
			case "Т":
				fmt.Print("T")
			case "У":
				fmt.Print("U")
			case "Ф":
				fmt.Print("F")
			case "Х":
				fmt.Print("H")
			case "Ц":
				fmt.Print("{\\C}")
			case "Ч":
				fmt.Print("C")
			case "Ш":
				fmt.Print("X")
			case "Щ":
				fmt.Print("{\\X}")
			case "Ъ":
				fmt.Print("Y")
			case "Ы":
				fmt.Print("YI")
			case "Ь":
				fmt.Print("Y")
			case "Э":
				fmt.Print("E")
			case "Ю":
				fmt.Print("{\\Y}u")
			case "Я":
				fmt.Print("{\\Y}a")
			case "а":
				fmt.Print("a")
			case "б":
				fmt.Print("b")
			case "в":
				fmt.Print("v")
			case "г":
				if p == "о" || p == "е" || p == "Е" {
					if i == wl-2 && string(word[i:]) == "го" {
						fmt.Print("vo")
						i += 1
					} else if i == wl-4 && string(word[i:]) == "гося" {
						fmt.Print("vosa")
						i += 3
					} else if i < wl-4 && string(word[i:i+3]) == "го-" {
						fmt.Print("vo-")
						i += 2
					} else {
						fmt.Print("g")
					}
				} else {
					fmt.Print("g")
				}
			case "д":
				fmt.Print("d")
			case "е":
				if p == "" || isVowel(p) {
					fmt.Print("{\\y}e")
				} else {
					fmt.Print("e")
				}
			case "ё":
				if p == "" || isVowel(p) {
					fmt.Print("{\\y}o")
				} else {
					if isFrict(p) {
						fmt.Print("o")
					} else {
						fmt.Print("{\\e}")
					}
				}
			case "ж":
				fmt.Print("j")
			case "з":
				fmt.Print("z")
			case "и":
				if isVowel(p) {
					if i == 3 && isPrefix(string(word[:3])) {
						fmt.Print("i")
					} else if i == 2 && isPrefix(string(word[:2])) {
						fmt.Print("i")
					} else {
						fmt.Print("{\\y}i")
					}
				} else {
					fmt.Print("i")
				}
			case "й":
				fmt.Print("{\\y}")
			case "к":
				fmt.Print("k")
			case "л":
				fmt.Print("l")
			case "м":
				fmt.Print("m")
			case "н":
				fmt.Print("n")
			case "о":
				fmt.Print("o")
			case "п":
				fmt.Print("p")
			case "р":
				fmt.Print("r")
			case "с":
				if i == wl-2 && string(word[i:]) == "ся" && (p == "у" || p == "ю" || p == "й" || p == "и" || p == "я" || p == "е" || p == "л" || p == "м" || p == "т" || p == "б" || p == "ь" || p == "х" || p == "г") {
					fmt.Print("sa")
					i += 1
				} else {
					fmt.Print("s")
				}
			case "т":
				fmt.Print("t")
			case "у":
				fmt.Print("u")
			case "ф":
				fmt.Print("f")
			case "х":
				fmt.Print("h")
			case "ц":
				fmt.Print("{\\c}")
			case "ч":
				fmt.Print("c")
			case "ш":
				fmt.Print("x")
			case "щ":
				fmt.Print("{\\x}")
			case "ы":
				if isFrict(p) {
					fmt.Print("i")
				} else {
					fmt.Print("{\\yi}")
				}
			case "ъ", "ь":
				if n == "ю" {
					fmt.Print("{\\y}u")
					i += 1
				} else if n == "я" {
					fmt.Print("{\\y}a")
					i += 1
				} else if n == "ё" {
					fmt.Print("{\\y}o")
					i += 1
				} else if n == "е" {
					fmt.Print("{\\y}e")
					i += 1
				} else if n == "и" {
					fmt.Print("{\\yf}i")
					i += 1
				} else if i == wl-3 && string(word[i+1:]) == "ся" && (p == "т" || p == "ш") {
					// skip
				} else if i == wl-1 && isFrict(p) {
					// skip
				} else {
					fmt.Print("y")
				}
			case "э":
				fmt.Print("e")
			case "ю":
				if p == "" || isVowel(p) {
					fmt.Print("{\\y}u")
				} else {
					fmt.Print("{\\iu}")
				}
			case "я":
				if p == "" || isVowel(p) {
					fmt.Print("{\\y}a")
				} else {
					fmt.Print("{\\ia}")
				}
			default:
				log.Fatalf("not valid: *%s*", c)
			}
		}
	}

	fmt.Print(footer)
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
		"банально":            "banalno",
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
		"грандиозной":         "grandiosno{\\y}",
		"грандиозную":         "grandiosnu{\\y}u",
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
		"курсе":               "curse",
		"Курьер":              "Courier",
		"курьера":             "couriera",
		"лабиринт":            "labyrinth",
		"логике":              "logice",
		"логику":              "logicu",
		"локальный":           "localn{\\yi}y",
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
		"музыкального":        "musicalnovo",
		"Музыкант":            "Musicant",
		"музыканта":           "musicanta",
		"музыкантов":          "musicantov",
		"музыку":              "musicu",
		"мэтр":                "maitre",
		"Мэтр":                "Maitre",
		"наркотик":            "narcotic",
		"наркотика":           "narcotica",
		"нейтральной":         "neutralno{\\y}",
		"нормальную":          "normalnu{\\y}u",
		"нотацию":             "notati{\\y}u",
		"ориентировку":        "orientirovku",
		"официальная":         "officialna{\\y}a",
		"паузу":               "pausu",
		"пеликан":             "pelican",
		"пеликана":            "pelicana",
		"пеликаном":           "pelicanom",
		"пеликану":            "pelicanu",
		"пеликанье":           "pelican{\\y}e",
		"персональный":        "personaln{\\yi}{\\y}",
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
		"реальная":            "realna{\\y}a",
		"реальности":          "realnosti",
		"реальность":          "realnosty",
		"Реальную":            "Realnu{\\y}u",
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
		"серьёзная":           "serieusna{\\y}a",
		"серьёзно":            "serieusno",
		"Серьёзно":            "Serieusno",
		"серьёзного":          "serieusnovo",
		"серьёзности":         "serieusnosti",
		"серьёзные":           "serieusn{\\yi}{\\y}e",
		"серьёзный":           "serieusn{\\yi}{\\y}",
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
		"специальный":         "specialn{\\yi}{\\y}",
		"специфической":       "specificesko{\\y}",
		"статистика":          "statistica",
		"статистике":          "statistice",
		"стерильности":        "sterilnosti",
		"субъект":             "subject",
		"суккубов":            "succubov",
		"сутана":              "soutana",
		"сэкономим":           "seconomim",
		"сэкономлю":           "seconoml{\\iu}",
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
		"формальность":        "formalnosty",
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
		"центральной":         "centralno{\\y}",
		"центральную":         "centralnu{\\y}u",
		"центральных":         "centraln{\\yi}h",
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
		"эмоциональный":       "emotionaln{\\yi}{\\y}",
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
		"нехристю-хагжиту":        "nechrist{\\iu}-hagjitu",
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
		"нехристю":       "nechrist{\\iu}",
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

const header = `\documentclass[12pt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}
`

const commands = `
\newcommand{\e}{ë}
%\newcommand{\e}{e}
%\newcommand{\e}{é}
%\newcommand{\e}{ó}

\newcommand{\yi}{\mbox{y\hspace{-0.55pt}ı}}
%\newcommand{\yi}{yi}
%\newcommand{\yi}{ǝ}

\newcommand{\ia}{\mbox{ı\hspace{-0.55pt}a}}
%\newcommand{\ia}{ia}
%\newcommand{\ia}{ía}
%\newcommand{\ia}{á}

\newcommand{\iu}{\mbox{ı\hspace{-0.55pt}o}}
%\newcommand{\iu}{io}
%\newcommand{\iu}{ío}
%\newcommand{\iu}{ıu}
%\newcommand{\iu}{iu}
%\newcommand{\iu}{íu}
%\newcommand{\iu}{ú}

\newcommand{\y}{y̆}
\newcommand{\yf}{y̆}
%\newcommand{\y}{y}

\newcommand{\Y}{Y̆}
%\newcommand{\Y}{Y}

\newcommand{\X}{X̹}
\newcommand{\x}{x̹}
\newcommand{\C}{C̹}
\renewcommand{\c}{c̹}

% % mixed ogonek and q
% \newcommand{\X}{X̨}
% \newcommand{\x}{x̨}
% \newcommand{\C}{Q}
% \renewcommand{\c}{q}
%
% % ogonek
% \newcommand{\X}{X̨}
% \newcommand{\x}{x̨}
% \newcommand{\C}{C̨}
% \renewcommand{\c}{c̨}
% 
% % right half ring
% \newcommand{\X}{X̹}
% \newcommand{\x}{x̹}
% \newcommand{\C}{C̹}
% \renewcommand{\c}{c̹}
%
% % retroflex hook
% \newcommand{\X}{X̢}
% \newcommand{\x}{x̢}
% \newcommand{\C}{C̢}
% \renewcommand{\c}{c̢}
% 
% % cedilla
% \newcommand{\X}{X̧}
% \newcommand{\x}{x̧}
% \newcommand{\C}{Ç}
% \renewcommand{\c}{ç}
% 
% % hook
% \newcommand{\X}{X̡}
% \newcommand{\x}{x̡}
% \newcommand{\C}{C̡}
% \renewcommand{\c}{c̡}
% 
% % acute accent below
% \newcommand{\X}{X̗}
% \newcommand{\x}{x̗}
% \newcommand{\C}{C̗}
% \renewcommand{\c}{c̗}
`

const footer = `
\end{document}
`
