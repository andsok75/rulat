package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("./text")
	if err != nil {
		log.Fatal(err)
	}

	items := []string{}
	word := ""
	for _, r := range string(content) {
		c := string(r)
		if isWord(c) {
			word = word + c
			continue
		}
		if word != "" {
			items = append(items, word)
			word = ""
		}
		items = append(items, c)
	}

	fmt.Print(`\documentclass[10pt]{book}
\usepackage{fontspec}
\setmainfont{Linux Libertine O}
\begin{document}

`)

	p := ""
	for _, r := range string(content) {
		c := string(r)
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
			fmt.Print("E")
		case "Ж":
			fmt.Print("J")
		case "З":
			fmt.Print("Z")
		case "И":
			fmt.Print("I")
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
			fmt.Print("Q")
		case "Ч":
			fmt.Print("C")
		case "Ш":
			fmt.Print("X")
		case "Э":
			fmt.Print("E")
		case "Я":
			fmt.Print("Ya")
		case "а":
			fmt.Print("a")
		case "б":
			fmt.Print("b")
		case "в":
			fmt.Print("v")
		case "г":
			fmt.Print("g")
		case "д":
			fmt.Print("d")
		case "е":
			if p == " " || isVowel(p) {
				fmt.Print("y̆e")
			} else {
				fmt.Print("e")
			}
		case "ж":
			fmt.Print("j")
		case "з":
			fmt.Print("z")
		case "и":
			if isVowel(p) {
				fmt.Print("y̆i")
			} else {
				fmt.Print("i")
			}
		case "й":
			fmt.Print("y̆")
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
			fmt.Print("s")
		case "т":
			fmt.Print("t")
		case "у":
			fmt.Print("u")
		case "ф":
			fmt.Print("f")
		case "х":
			fmt.Print("h")
		case "ц":
			fmt.Print("q")
		case "ч":
			fmt.Print("c")
		case "ш":
			fmt.Print("x")
		case "щ":
			fmt.Print("x̨")
		case "ы":
			fmt.Print("yı")
		case "ь":
			fmt.Print("y")
		case "э":
			fmt.Print("e")
		case "ю":
			if p == " " || isVowel(p) {
				fmt.Print("y̆u")
			} else {
				fmt.Print("ıu")
			}
		case "я":
			if p == " " || isVowel(p) {
				fmt.Print("y̆a")
			} else {
				fmt.Print("ıa")
			}
		default:
			fmt.Print(c)
		}
		p = c
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
	case "А", "Б", "В", "Г", "Д", "Е", "Ж", "З", "И", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Э", "Я":
		fallthrough
	case "а", "б", "в", "г", "д", "е", "ж", "з", "и", "й", "к", "л", "м", "н", "о", "п", "р", "с", "т", "у", "ф", "х", "ц", "ч", "ш", "щ", "ы", "ь", "э", "ю", "я":
		return true
	default:
		return false
	}
}
