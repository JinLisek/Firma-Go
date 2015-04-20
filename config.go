// config
package main

type tryb int

const (
	GADATLIWY tryb = 1 + iota
	SPOKOJNY
)

type dzialanie func(int, int) int

type zadanie struct {
	liczba1 int
	liczba2 int
	dzial   dzialanie
}

func suma(a int, b int) int {
	return a + b
}

func roznica(a, b int) int {
	return a - b
}

func iloczyn(a, b int) int {
	return a * b
}

var zadania []zadanie
var produkty []int
var ilosc_zadan int
var ilosc_produktow int

const tryb_symulacji tryb = SPOKOJNY
const pojemnosc_listy int = 10
const pojemnosc_magazynu int = 10
const liczba_pracownikow int = 3
const liczba_klientow int = 3
const opoznienie_prezesa int = 500
const opoznienie_pracownikow int = 2000
const opoznienie_klientow int = 10000
const interwal int = 1000
const trwanie_app int = 10000
