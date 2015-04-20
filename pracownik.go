// pracownik
package main

import (
	"fmt"
	"time"
)

func pracownik(id int, lista_prac <-chan zadanie, prac_magazyn chan<- int, magazyn_prac <-chan bool) {
	var z zadanie
	var produkt int

	for {

		select {
		case z = <-lista_prac:
			produkt = z.dzial(z.liczba1, z.liczba2)

			if tryb_symulacji == GADATLIWY {
				fmt.Println(id, "pracownik wykonal produkt:", produkt)
			}

			prac_magazyn <- produkt
			var sukces bool = <-magazyn_prac

			if tryb_symulacji == GADATLIWY {
				if sukces {
					fmt.Println(id, "pracownik dodal do magazynu:", produkt)
				} else {
					fmt.Println(id, "pracownik probowal dodac do magazynu:", produkt, "jednak magazyn jest pelny")
				}
			}

		case <-time.After(time.Millisecond * time.Duration(opoznienie_pracownikow)):
			if tryb_symulacji == GADATLIWY {
				fmt.Println(id, "pracownik probowal wziac zadanie z listy, jednak lista jest pusta")
			}
		}

		time.Sleep(time.Millisecond * time.Duration(opoznienie_pracownikow))
	}
}
