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
			}
			if i == wl-1 {
				n = ""
			}
			if i < wl-1 {
				n = string(word[i+1])
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
				fmt.Print("{\\Q}")
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
				if i == wl-2 && n == "о" && (p == "о" || p == "е" || p == "Е") {
					fmt.Print("v")
				} else if i < wl-2 && string(word[i:i+3]) == "го-" && (p == "о" || p == "е" || p == "Е") {
					fmt.Print("v")
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
				if i == wl-2 && n == "я" && (p == "й" || p == "я" || p == "е" || p == "л" || p == "м" || p == "т" || p == "б" || p == "ь") {
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
				fmt.Print("{\\q}")
			case "ч":
				fmt.Print("c")
			case "ш":
				fmt.Print("x")
			case "щ":
				fmt.Print("{\\x}")
			case "ы":
				if p == "ц" {
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
				} else if i == wl-3 && (p == "т" || p == "ш") && n == "с" && string(word[wl-1]) == "я" {
				} else if isFrict(p) && n == "" {
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
	return s == "про" || s == "по" || s == "за" || s == "на"
}

func exceptions() map[string]string {
	return map[string]string{
		"сегодня":          "sevodn{\\ia}",
		"немного":          "nemnogo",
		"много":            "mnogo",
		"аист":             "aist",
		"наивен":           "naiven",
		"Воистину":         "Voistinu",
		"кацеров":          "katzerov",
		"чувствовал":       "custvoval",
		"Иисусе":           "Iesuse",
		"Христа":           "Christa",
		"Христово":         "Christovo",
		"Христовой":        "Christovo{\\y}",
		"нехристю-хагжиту": "nechrist{\\iu}-hagjitu",
		"Михаила":          "Michaela",
		"Петра":            "Petera",
		"Дель":             "Del",
		"Людвиг":           "Ludwig",
		"Лезерберг":        "Leserberg",
		"Фирвальден":       "Firvalden",
		"Фирвальдене":      "Firvaldene",
		"Фабьен":           "Fabien",
		"Клеменз":          "Clemence",
		"Ганс":             "Hans",
		"Ганса":            "Hansa",
		"Гансом":           "Hansom",
		"Альбаланд":        "Albaland",
		"Альбаланда":       "Albalanda",
		"Лисецк":           "Lisetsk",
		"Лисецке":          "Lisetske",
		"Витильск":         "Witilsk",
		"Витильска":        "Witilska",
		"Дорч-ган-Тойн":    "Dortch-gan-Toyn",
		"Дорч-ган-Тойне":   "Dortch-gan-Toyne",
		"Бробергер":        "Broberger",
		"Бробергера":       "Brobergera",
		"Хартвиг":          "Hartwig",
		"Хартвига":         "Hartwiga",
		"Шуко":             "Shuko",
		"Рози":             "Rosi",
		"Шоссии":           "Shossi{\\y}a",
		"Валентин":         "Valentin",
		"Валентина":        "Valentina",
		"Вальтер":          "Walter",
		"Нормайенн":        "Normaenn",
	}
}

const header = `\documentclass[10pt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}

\newcommand{\e}{ë}
%\newcommand{\e}{e}
%\newcommand{\e}{é}
%\newcommand{\e}{ó}

\renewcommand{\i}{ı}
%\renewcommand{\i}{i}

\newcommand{\yi}{yı}
%\newcommand{\yi}{yi}
%\newcommand{\yi}{ǝ}

\newcommand{\ia}{ıa}
%\newcommand{\ia}{ia}
%\newcommand{\ia}{ía}
%\newcommand{\ia}{á}

\newcommand{\iu}{ıo}
%\newcommand{\iu}{ıu}
%\newcommand{\iu}{iu}
%\newcommand{\iu}{io}
%\newcommand{\iu}{ío}
%\newcommand{\iu}{íu}
%\newcommand{\iu}{ú}

\newcommand{\y}{y̆}
%\newcommand{\y}{y}
\newcommand{\yf}{y̆}

\newcommand{\Y}{Y̆}
%\newcommand{\Y}{Y}

\newcommand{\X}{X̹}
\newcommand{\x}{x̹}
\newcommand{\Q}{C̹}
\newcommand{\q}{c̹}

% % mixed ogonek and q
% \newcommand{\X}{X̨}
% \newcommand{\x}{x̨}
% \newcommand{\Q}{Q}
% \newcommand{\q}{q}
%
% % ogonek
% \newcommand{\X}{X̨}
% \newcommand{\x}{x̨}
% \newcommand{\Q}{C̨}
% \newcommand{\q}{c̨}
% 
% % right half ring
% \newcommand{\X}{X̹}
% \newcommand{\x}{x̹}
% \newcommand{\Q}{C̹}
% \newcommand{\q}{c̹}
%
% % retroflex hook
% \newcommand{\X}{X̢}
% \newcommand{\x}{x̢}
% \newcommand{\Q}{C̢}
% \newcommand{\q}{c̢}
% 
% % cedilla
% \newcommand{\X}{X̧}
% \newcommand{\x}{x̧}
% \newcommand{\Q}{Ç}
% \newcommand{\q}{ç}
% 
% % hook
% \newcommand{\X}{X̡}
% \newcommand{\x}{x̡}
% \newcommand{\Q}{C̡}
% \newcommand{\q}{c̡}
% 
% % acute accent below
% \newcommand{\X}{X̗}
% \newcommand{\x}{x̗}
% \newcommand{\Q}{C̗}
% \newcommand{\q}{c̗}

`

const footer = `
\end{document}
`
