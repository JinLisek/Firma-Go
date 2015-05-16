package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SayHello(to string) {
	fmt.Printf("Hello, %s!\n", to)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	prez_lista := make(chan zadanie)
	lista_prez := make(chan bool)
	lista_prac := make(chan zadanie)
	prac_magazyn := make(chan int)
	magazyn_prac := make(chan bool)
	magazyn_klient := make(chan int)
	koniec := make(chan bool)
	prac_sumer := make(chan zadanie)
	sumer_prac := make(chan zadanie)
	sukces_sumer := make(chan bool)
	sumer_serwisant := make(chan bool)
	serwisant_sumer := make(chan bool)
	prac_iloczyner := make(chan zadanie)
	iloczyner_prac := make(chan zadanie)
	sukces_iloczyner := make(chan bool)
	iloczyner_serwisant := make(chan bool)
	serwisant_iloczyner := make(chan bool)
	pomoz := make(chan bool)

	go prezes(prez_lista, lista_prez)

	go lista_zadan(prez_lista, lista_prez, lista_prac)

	for i := 1; i <= liczba_pracownikow; i++ {
		go pracownik(i, lista_prac, prac_sumer, sukces_sumer, sumer_prac, prac_iloczyner, sukces_iloczyner, iloczyner_prac, prac_magazyn, magazyn_prac, pomoz)
	}

	for i := 1; i <= liczba_sumerow; i++ {
		go sumer(i, prac_sumer, sukces_sumer, sumer_prac, sumer_serwisant, serwisant_sumer)
	}

	for i := 1; i <= liczba_iloczynerow; i++ {
		go iloczyner(i, prac_iloczyner, sukces_iloczyner, iloczyner_prac, iloczyner_serwisant, serwisant_iloczyner, pomoz)
	}

	go serwisant(sumer_serwisant, serwisant_sumer, iloczyner_serwisant, serwisant_iloczyner)

	go magazyn(prac_magazyn, magazyn_prac, magazyn_klient)

	for i := 1; i <= liczba_klientow; i++ {
		go klient(i, magazyn_klient)
	}

	if tryb_symulacji == SPOKOJNY {
		go interfejs(koniec)
		var b bool = <-koniec
		if b {
			return
		}
		//time.Sleep(time.Millisecond * time.Duration(trwanie_app))
	} else {
		var input string
		fmt.Scanln(&input)
	}

}
