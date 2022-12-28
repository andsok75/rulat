package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const input_file = "./text"

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
				if i == wl-2 && string(word[i:]) == "ся" && (p == "у" || p == "ю" || p == "й" || p == "я" || p == "е" || p == "л" || p == "м" || p == "т" || p == "б" || p == "ь" || p == "х" || p == "г") {
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
		"апокалипсис":         "apocalypsis",
		"континент":           "continent",
		"континента":          "continenta",
		"пеликан":             "pelican",
		"пеликану":            "pelicanu",
		"пеликанье":           "pelican{\\y}e",
		"пеликана":            "pelicana",
		"пеликаном":           "pelicanom",
		"климатом":            "climatom",
		"теологический":       "theologiceski{\\y}",
		"фактах":              "factah",
		"факты":               "fact{\\yi}",
		"артефакт":            "artefact",
		"фальцетом":           "falsettom",
		"гимн":                "hymn",
		"гимны":               "hymn{\\yi}",
		"миф":                 "myth",
		"мифических":          "mythiceskih",
		"мифический":          "mythiceski{\\y}",
		"мифами":              "mythami",
		"критическую":         "criticesku{\\y}u",
		"бухгалтерской":       "buchhaltersko{\\y}",
		"шифр":                "chiffre",
		"коллекция":           "collecti{\\y}a",
		"коллекцию":           "collecti{\\y}u",
		"Коллекционер":        "Collectioner",
		"коллекционер":        "collectioner",
		"коллекционера":       "collectionera",
		"коллекционером":      "collectionerom",
		"коллекционеров":      "collectionerov",
		"курьера":             "couriera",
		"Курьер":              "Courier",
		"серафима":            "seraphima",
		"клиент":              "client",
		"клиентов":            "clientov",
		"репутация":           "reputati{\\y}a",
		"репутацией":          "reputati{\\y}e{\\y}",
		"контура":             "contoura",
		"архангела":           "archangela",
		"аккуратно":           "accuratno",
		"Тракт":               "Tract",
		"тракту":              "tractu",
		"тракте":              "tracte",
		"трактирах":           "tractirah",
		"энтузиазма":          "enthusiasma",
		"Кардинал":            "Cardinal",
		"кардинал":            "cardinal",
		"кардинала":           "cardinala",
		"кардиналу":           "cardinalu",
		"кардиналов":          "cardinalov",
		"кардинальская":       "cardinalska{\\y}a",
		"Кардинальская":       "Cardinalska{\\y}a",
		"лабиринт":            "labyrinth",
		"музыка":              "musica",
		"музыку":              "musicu",
		"Музыкант":            "Musicant",
		"музыканта":           "musicanta",
		"музыкантов":          "musicantov",
		"музыкального":        "musicalnovo",
		"музицирования":       "musicirovani{\\y}a",
		"балкон":              "balcon",
		"колоннаду":           "colonnadu",
		"грифелем":            "griffelem",
		"гильдии":             "gildi{\\y}i",
		"раухтопазов":         "rauchtopazov",
		"браслет":             "bracelet",
		"аптекой":             "apotheko{\\y}",
		"аптеке":              "apotheke",
		"аптеку":              "apotheku",
		"аптекарь":            "apothekary",
		"Аптекарь":            "Apothekary",
		"аптекарем":           "apothekarem",
		"аптекаря":            "apothekar{\\ia}",
		"Аптека":              "Apotheka",
		"комод":               "commode",
		"бургомистр":          "burgemeester",
		"бургомистра":         "burgemeestera",
		"конфликт":            "conflict",
		"Конфликты":           "Conflict{\\yi}",
		"истории":             "istori{\\y}i",
		"историю":             "istori{\\y}u",
		"историков":           "istorikov",
		"Предыстория":         "Predistori{\\y}a",
		"серьёзный":           "serieusn{\\yi}{\\y}",
		"серьёзно":            "serieusno",
		"Серьёзно":            "Serieusno",
		"серьёзности":         "serieusnosti",
		"серьёзная":           "serieusna{\\y}a",
		"серьёзные":           "serieusn{\\yi}{\\y}e",
		"серьёзного":          "serieusnovo",
		"информация":          "informati{\\y}a",
		"информации":          "informati{\\y}i",
		"информацию":          "informati{\\y}u",
		"сэкономим":           "seconomim",
		"сэкономлю":           "seconoml{\\iu}",
		"коллегой":            "collego{\\y}",
		"специфической":       "specificesko{\\y}",
		"теорию":              "theori{\\y}u",
		"теория":              "theori{\\y}a",
		"теории":              "theori{\\y}i",
		"Теория":              "Theori{\\y}a",
		"Катастрофы":          "Catastroph{\\yi}",
		"катастрофы":          "catastroph{\\yi}",
		"еретическая":         "hereticeska{\\y}a",
		"Епископская":         "Episkopska{\\y}a",
		"епископ":             "episkop",
		"епископа":            "episkopa",
		"секта":               "secta",
		"манускрипты":         "manuscript{\\yi}",
		"император":           "emperator",
		"императора":          "emperatora",
		"императору":          "emperatoru",
		"Империя":             "Emperi{\\y}a",
		"экспедиций":          "expediti{\\y}",
		"суккубов":            "succubov",
		"логику":              "logicu",
		"логике":              "logice",
		"Мэтр":                "Maitre",
		"мэтр":                "maitre",
		"Кондотьер":           "Condottier",
		"кондотьер":           "condottier",
		"кондотьера":          "condottiera",
		"терьер":              "terrier",
		"кацеров":             "katzerov",
		"секунд":              "secund",
		"секунду":             "secundu",
		"гвардейцев":          "guarde{\\y}{\\c}ev",
		"кобра":               "cobra",
		"экземпляром":         "exemplarom",
		"скептицизма":         "scepticizma",
		"специалист":          "specialist",
		"специалисты":         "specialist{\\yi}",
		"инквизитор":          "inquisitor",
		"инквизитором":        "inquisitorom",
		"инквизиция":          "inquisiti{\\y}a",
		"Инквизиция":          "Inquisiti{\\y}a",
		"клирики":             "clerici",
		"клириков":            "clericov",
		"клирикам":            "clericam",
		"клириками":           "clericami",
		"клирика":             "clerica",
		"клириках":            "clericah",
		"Клириков":            "Clericov",
		"Клирик":              "Cleric",
		"безапелляционно":     "bezappellationno",
		"картограф":           "cartograph",
		"картографом":         "cartographom",
		"ситуации":            "situati{\\y}i",
		"ситуация":            "situati{\\y}a",
		"ситуацию":            "situati{\\y}u",
		"магистров":           "magisterov",
		"Магистры":            "Magister{\\yi}",
		"магистрами":          "magisterami",
		"перспектива":         "perspectiva",
		"перспективе":         "perspective",
		"Конкретизируй":       "Concretiziru{\\y}",
		"конкретики":          "concretiki",
		"канонизация":         "canoniza{\\c}i{\\y}a",
		"каноник":             "canonik",
		"каноника":            "canonika",
		"дукат":               "ducat",
		"дукатов":             "ducatov",
		"монографий":          "monographi{\\y}",
		"клерк":               "clerk",
		"клан":                "clan",
		"фраза":               "phrasa",
		"фразу":               "phrasu",
		"библиотека":          "bibliotheca",
		"библиотеки":          "bibliotheci",
		"регулярно":           "regularno",
		"сапфир":              "sapphir",
		"сапфиром":            "sapphirom",
		"сапфирами":           "sapphirami",
		"копию":               "copi{\\y}u",
		"копия":               "copi{\\y}a",
		"Копия":               "Copi{\\y}a",
		"контраргумент":       "contrargument",
		"дискредитировать":    "discreditirovaty",
		"дискредитируют":      "discreditiru{\\y}ut",
		"характерную":         "characternu{\\y}u",
		"характерных":         "charactern{\\yi}h",
		"слабохарактерностью": "slabocharacternost{\\y}u",
		"эмоций":              "emoti{\\y}",
		"эмоции":              "emoti{\\y}i",
		"эмоциям":             "emoti{\\y}am",
		"нотацию":             "notati{\\y}u",
		"этаже":               "etage",
		"этажа":               "etaga",
		"этаж":                "etage",
		"компания":            "compani{\\y}a",
		"компании":            "compani{\\y}i",
		"алхимией":            "alchimi{\\y}e{\\y}",
		"алхимии":             "alchimi{\\y}i",
		"алхимических":        "alchimiceskih",
		"комментариев":        "commentari{\\y}ev",
		"актёр":               "acteur",
		"депозите":            "deposite",
		"фантазия":            "fantasi{\\y}a",
		"фантазёры":           "fantas{\\e}r{\\yi}",
		"техники":             "techniki",
		"грандиозную":         "grandiosnu{\\y}u",
		"грандиозной":         "grandiosno{\\y}",
		"аферу":               "affairu",
		"кретины":             "cretin{\\yi}",
		"школа":               "schola",
		"Процессия":           "Processi{\\y}a",
		"процессия":           "processi{\\y}a",
		"шпиона":              "spiona",
		"шпионов":             "spionov",
		"жест":                "geste",
		"жесте":               "geste",
		"Трактаты":            "Tractat{\\yi}",
		"эхом":                "echom",
		"горнистов":           "hornistov",
		"маршировала":         "marschierovala",
		"кирасах":             "cuirassah",
		"алебардами":          "hallebardami",
		"алебарды":            "hallebard{\\yi}",
		"кавалькада":          "cavalcada",
		"кавалькадой":         "cavalcado{\\y}",
		"делегации":           "delegati{\\y}i",
		"сутана":              "soutana",
		"фальшивка":           "falsivka",
		"эффект":              "effect",
		"кантонские":          "cantonski{\\y}e",
		"Кантонец":            "Cantone{\\c}",
		"кантонских":          "cantonskih",
		"капитан":             "capitan",
		"шапке":               "chapke",
		"шапку":               "chapku",
		"субъект":             "subject",
		"центру":              "centru",
		"центре":              "centre",
		"центральных":         "centraln{\\yi}h",
		"центральную":         "centralnu{\\y}u",
		"центральной":         "centralno{\\y}",
		"зафиксированы":       "zafixirovan{\\yi}",
		"механизм":            "machanism",
		"активнее":            "activne{\\y}e",
		"кабинет":             "cabinet",
		"официальная":         "officialna{\\y}a",
		"контексте":           "contexte",
		"велефа":              "veletha",
		"атаки":               "attacki",
		"локальный":           "localn{\\yi}y",
		"гримуары":            "grimoir{\\yi}",
		"кандидатура":         "candidatura",
		"хаос":                "chaos",
		"хаотичных":           "chaoticn{\\yi}h",
		"иллюзии":             "illusi{\\y}i",
		"мебель":              "meuble",
		"инициатива":          "initiativa",
		"сюртуке":             "surtuke",
		"тюрьмы":              "turm{\\yi}",
		"тюремный":            "turemn{\\yi}{\\y}",
		"дракона":             "dracona",
		"силуэт":              "silhouette",
		"силуэта":             "silhouetta",
		"силуэты":             "silhouett{\\yi}",
		"металл":              "metal",
		"металла":             "metalla",
		"наркотик":            "narcotic",
		"наркотика":           "narcotica",
		"монастыря":           "monaster{\\ia}",
		"монастырём":          "monaster{\\e}m",
		"монастыре":           "monastere",
		"монастырь":           "monastery",
		"монах":               "monach",
		"монахов":             "monachov",
		"монаху":              "monachu",
		"тифа":                "typha",
		"капюшон":             "capuchon",
		"реальности":          "realnosti",
		"реальность":          "realnosty",
		"реальная":            "realna{\\y}a",
		"Реальную":            "Realnu{\\y}u",
		"интерьеру":           "interieuru",
		"бесцеремонно":        "besceremonno",
		"комбинацией":         "combinati{\\y}e{\\y}",
		"акцентом":            "accentom",
		"короной":             "corono{\\y}",
		"коридор":             "corridor",
		"бутылку":             "butelku",
		"бутылки":             "butelki",
		"персональный":        "personaln{\\yi}{\\y}",
		"дьявол":              "diavol",
		"дьявольское":         "diavolsko{\\y}e",
		"дьявольских":         "diavolskih",
		"дьяволопоклонника":   "diavolopoklonnika",
		"паузу":               "pausu",
		"секреты":             "secret{\\yi}",
		"секретов":            "secretov",
		"курсе":               "curse",
		"меланхолично":        "melancholicno",
		"реликвию":            "reliqui{\\y}u",
		"реликвии":            "reliqui{\\y}i",
		"скандал":             "scandal",
		"скандалов":           "scandalov",
		"реформации":          "reformati{\\y}i",
		"банально":            "banalno",
		"результат":           "resultat",
		"безрезультатно":      "bezresultatno",
		"нейтральной":         "neutralno{\\y}",
		"рецепт":              "recept",
		"Резервный":           "Reservn{\\yi}{\\y}",
		"компенсировать":      "compensirovat{\\y}",
		"гротескная":          "grottescna{\\y}a",
		"камере":              "camere",
		"тему":                "themu",
		"ратуше":              "rathuse",
		"блокировал":          "blockieroval",
		"вулкан":              "vulcan",
		"команды":             "command{\\yi}",
		"Команда":             "Commanda",
		"капризов":            "capricov",

		"гвардейцев-альбаландцев": "guarde{\\y}{\\c}ev-albaland{\\c}ev",
		"по-нарарски":             "po-nararrski",
		"нехристю-хагжиту":        "nechrist{\\iu}-hagjitu",
		"женщину-музыканта":       "jen{\\x}inu-musicanta",
		"беднягу-аптекаря":        "bedn{\\ia}gu-apothekar{\\ia}",
		"клирики-тупицы":          "clerici-tupi{\\c}i",
		"монах-привратник":        "monach-privratnik",

		"Иисусе":         "Iesuse",
		"Христа":         "Christa",
		"Христе":         "Christe",
		"Христово":       "Christovo",
		"Христовой":      "Christovo{\\y}",
		"нехристю":       "nechrist{\\iu}",
		"Михаила":        "Michaela",
		"Петра":          "Petera",
		"Дель":           "Del",
		"Людвиг":         "Ludwig",
		"Лезерберг":      "Leserberg",
		"Лезерберга":     "Leserberga",
		"Фирвальден":     "Firvalden",
		"Фирвальдене":    "Firvaldene",
		"Фабьен":         "Fabien",
		"Клеменз":        "Clemence",
		"Ганс":           "Hans",
		"Ганса":          "Hansa",
		"Гансом":         "Hansom",
		"Альбаланд":      "Albaland",
		"Альбаланда":     "Albalanda",
		"Лисецк":         "Lesetsk",
		"Лисецке":        "Lesetske",
		"Витильск":       "Witilsk",
		"Витильска":      "Witilska",
		"Витильске":      "Witilske",
		"Дорч-ган-Тойн":  "Dortch-gan-Toyn",
		"Дорч-ган-Тойне": "Dortch-gan-Toyne",
		"Бробергер":      "Wroberger",
		"Бробергера":     "Wrobergera",
		"Хартвиг":        "Hartwig",
		"Хартвига":       "Hartwiga",
		"Шуко":           "Schuco",
		"Рози":           "Rosi",
		"Шоссии":         "Schossi{\\y}i",
		"Валентин":       "Valentin",
		"Валентина":      "Valentina",
		"Вальтер":        "Walter",
		"Вальтера":       "Waltera",
		"Вальтеру":       "Walteru",
		"Вальтером":      "Walterom",
		"Нормайенн":      "Normaenn",
		"Нормайенном":    "Normaennom",
		"Нормайенна":     "Normaenna",
		"Филипп":         "Philipp",
		"Филиппу":        "Philippu",
		"Филиппом":       "Philippom",
		"Натан":          "Nathan",
		"Натана":         "Nathana",
		"Фрингбоу":       "Fringboe",
		"Травинно":       "Travino",
		"Джулию":         "Juliyu",
		"Адиль":          "Adil",
		"Адиля":          "Adil{\\ia}",
		"аль":            "al",
		"Джума":          "Juma",
		"Лавендуззского": "Lavendusskovo",
		"Чезаре":         "Cesare",
		"Каварзере":      "Cavarsere",
		"Ньюгорт":        "Newgord",
		"Ханна":          "Hannah",
		"Селико":         "Selico",
		"Микель":         "Michel",
		"Барбурге":       "Warburge",
		"Константин":     "Constantin",
		"Константина":    "Constantina",
		"Константину":    "Constantinu",
		"Крусо":          "Caruso",
		"Нарара":         "Nararra",
		"Нарары":         "Nararr{\\yi}",
		"флотолийских":   "flottoli{\\y}skih",
		"альбаландцев":   "albaland{\\c}ev",
		"альбаландец":    "albalande{\\c}",
		"альбаландских":  "albalandskih",
		"Александр":      "Alexander",
		"Александра":     "Alexandera",
		"Ливетте":        "Levette",
		"Вьюну":          "Vewnu",
		"Вьюна":          "Vewna",
		"Каликвец":       "Callicve{\\c}",
		"Каликвецы":      "Callicve{\\c}i",
		"каликвец":       "callicve{\\c}",
		"каликвецев":     "callicve{\\c}ev",
		"каликвеца":      "callicve{\\c}a",
		"Прогансу":       "Progance",

		"сегодня":       "sevodn{\\ia}",
		"Сегодня":       "Sevodn{\\ia}",
		"немного":       "nemnogo",
		"ненамного":     "nenamnogo",
		"намного":       "namnogo",
		"много":         "mnogo",
		"аист":          "aist",
		"наивен":        "naiven",
		"Воистину":      "Voistinu",
		"чувствуя":      "custvu{\\y}a",
		"чувствовал":    "custvoval",
		"чувствовала":   "custvovala",
		"почувствовать": "pocustvovaty",
		"почувствовал":  "pocustvoval",
		"предыдущая":    "predidu{\\x}a{\\y}a",
		"предыдущего":   "predidu{\\x}evo",
		"собиравшимися": "sobiravximisa",
		"стоит":         "stoit",
	}
}

const header = `\documentclass[10pt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}
`

const commands = `
\newcommand{\e}{ë}
%\newcommand{\e}{e}
%\newcommand{\e}{é}
%\newcommand{\e}{ó}

\newcommand{\yi}{yı}
%\newcommand{\yi}{yi}
%\newcommand{\yi}{ǝ}

\newcommand{\ia}{ıa}
%\newcommand{\ia}{ia}
%\newcommand{\ia}{ía}
%\newcommand{\ia}{á}

\newcommand{\iu}{ıo}
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
