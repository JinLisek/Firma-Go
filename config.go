// config
package main

type tryb int

const (
	GADATLIWY tryb = 1 + iota
	SPOKOJNY
)

type dzialanie int

const (
	SUMA dzialanie = 1 + iota
	ILOCZYN
)

type zadanie struct {
	liczba1 int
	liczba2 int
	dzial   dzialanie
	wynik   int
}

var zadania []zadanie
var produkty []int
var ilosc_zadan int
var ilosc_produktow int

const tryb_symulacji tryb = GADATLIWY
const pojemnosc_listy int = 2
const pojemnosc_magazynu int = 3
const liczba_pracownikow int = 3
const liczba_klientow int = 2
const liczba_sumerow int = 2
const liczba_iloczynerow int = 2
const opoznienie_prezesa int = 3000
const opoznienie_pracownikow int = 3000
const opoznienie_klientow int = 5000
const opoznienie_sumera int = 500
const opoznienie_iloczynera int = 500
const opoznienie_serwisanta int = 10000
const psucie_sumera int = 10
const psucie_iloczynera int = 10
const oczekiwanie_na_sumer int = opoznienie_sumera * 8
const oczekiwanie_na_iloczyner int = opoznienie_iloczynera * 8
const interwal int = 1000
const trwanie_app int = 30000
