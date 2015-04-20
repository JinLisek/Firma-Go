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

	go prezes(prez_lista, lista_prez)

	go lista_zadan(prez_lista, lista_prez, lista_prac)

	for i := 1; i <= liczba_pracownikow; i++ {
		go pracownik(i, lista_prac, prac_magazyn, magazyn_prac)
	}

	go magazyn(prac_magazyn, magazyn_prac, magazyn_klient)

	for i := 1; i <= liczba_klientow; i++ {
		go klient(i, magazyn_klient)
	}

	if tryb_symulacji == SPOKOJNY {
		go interfejs()
		time.Sleep(time.Millisecond * time.Duration(trwanie_app))
	} else {
		var input string
		fmt.Scanln(&input)
	}

}
