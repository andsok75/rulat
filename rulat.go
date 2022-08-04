package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Item struct {
	isWord  bool
	content string
}

func main() {
	content, err := ioutil.ReadFile("./text")
	if err != nil {
		log.Fatal(err)
	}

	items := []Item{}
	word := ""
	for _, r := range string(content) {
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

	fmt.Print(`\documentclass[10pt]{book}
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

\newcommand{\Y}{Y̆}
%\newcommand{\Y}{Y}

\newcommand{\C}{C}
\renewcommand{\c}{c}

\newcommand{\X}{X̨}
\newcommand{\x}{x̨}
\newcommand{\Q}{Q}
\newcommand{\q}{q}

% % ogonek
% \newcommand{\X}{X̨}
% \newcommand{\x}{x̨}
% \newcommand{\Q}{C̨}
% \newcommand{\q}{c̨}
% 
% % retroflex hook
% \newcommand{\X}{X̢}
% \newcommand{\x}{x̢}
% \newcommand{\Q}{C̢}
% \renewcommand{\q}{c̢}
% 
% % cedilla
% \newcommand{\X}{X̧}
% \newcommand{\x}{x̧}
% \newcommand{\Q}{Ç}
% \renewcommand{\q}{ç}
% 
% % hook
% \newcommand{\X}{X̡}
% \newcommand{\x}{x̡}
% \newcommand{\Q}{C̡}
% \renewcommand{\q}{c̡}
% 
% % acute accent
% \newcommand{\X}{X̗}
% \newcommand{\x}{x̗}
% \newcommand{\Q}{C̗}
% \renewcommand{\q}{c̗}

`)

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
		switch item.content {
		case "Христа":
			fmt.Print("Christa")
			continue
		case "Христово":
			fmt.Print("Christovo")
			continue
		case "Христовой":
			fmt.Print("Christovo{\\y}")
			continue
		case "нехристю":
			fmt.Print("nechrist{\\i}u")
			continue

		case "немного":
			fmt.Print("nemnogo")
			continue
		case "много":
			fmt.Print("mnogo")
			continue
		case "аист":
			fmt.Print("aist")
			continue

		case "наизнанку":
			fmt.Print("naiznanku")
			continue
		case "происходит":
			fmt.Print("proishodit")
			continue
		case "заинтересовали":
			fmt.Print("zainteresovali")
			continue
		case "заинтересованно":
			fmt.Print("zainteresovanno")
			continue
		case "Воистину":
			fmt.Print("Voistinu")
			continue
		case "наивен":
			fmt.Print("naiven")
			continue

		case "несущаяся":
			fmt.Print("nesu{\\x}a{\\y}asa")
			continue

		case "бургомистра":
			fmt.Print("burgermeistera")
			continue
		case "шифр":
			fmt.Print("chiffre")
			continue
		case "бухгалтерской":
			fmt.Print("buchhalterskoy")
			continue
		case "кацеров":
			fmt.Print("katzerov")
			continue
		case "миф":
			fmt.Print("myth")
			continue
		case "тифа":
			fmt.Print("typha")
			continue
		case "капюшон":
			fmt.Print("capuchon")
			continue
		case "коллекцию":
			fmt.Print("collecti{\\y}u")
			continue

		case "Аванс":
			fmt.Print("Avance")
			continue
		case "Михаила":
			fmt.Print("Michaela")
			continue
		case "Дель":
			fmt.Print("Del")
			continue
		case "Людвиг":
			fmt.Print("Ludwig")
			continue
		case "Лезерберг":
			fmt.Print("Leserberg")
			continue
		case "Фирвальдене":
			fmt.Print("Firvaldene")
			continue
		case "Фабьен":
			fmt.Print("Fabien")
			continue
		case "Клеменз":
			fmt.Print("Clemence")
			continue
		case "Иисусе":
			fmt.Print("Iesuse")
			continue
		default:
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
				fmt.Print("{\\C}")
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
					if isFrik(p) {
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
					fmt.Print("{\\y}i")
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
				if i == wl-2 && n == "я" && (p == "л" || p == "м" || p == "т" || p == "б" || p == "ь") {
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
				fmt.Print("{\\c}")
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
				} else if i == wl-3 && (p == "т" || p == "ш") && n == "с" && string(word[wl-1]) == "я" {
				} else if isFrik(p) && n == "" {
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
				log.Fatalf("not a word: *%s*", c)
			}
		}
	}

	fmt.Print(`
\end{document}
`)
}

func isVowel(c string) bool {
	if c == "а" || c == "е" || c == "и" || c == "о" || c == "у" || c == "ы" || c == "э" || c == "ю" || c == "я" {
		return true
	} else {
		return false
	}
}

func isWord(c string) bool {
	switch c {
	case "А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я":
		fallthrough
	case "а", "б", "в", "г", "д", "е", "ё", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ъ", "ы", "ь", "э", "ю", "я":
		return true
	default:
		return false
	}
}

func isFrik(c string) bool {
	switch c {
	case "Ж", "Ц", "Ч", "Ш", "Щ":
		fallthrough
	case "ж", "ц", "ч", "ш", "щ":
		return true
	default:
		return false
	}
}
