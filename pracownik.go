// pracownik
package main

import (
	"fmt"
	"time"
)

func pracownik(id int, lista_prac <-chan zadanie,
	prac_sumer chan<- zadanie, sukces_sumer <-chan bool, sumer_prac <-chan zadanie,
	prac_iloczyner chan<- zadanie, sukces_iloczyner <-chan bool, iloczyner_prac <-chan zadanie,
	prac_magazyn chan<- int, magazyn_prac <-chan bool, pomoz chan<- bool) {
	var z zadanie
	var sukces bool
	var znak string

	for {

		select {
		case z = <-lista_prac:
			if z.dzial == SUMA {
				znak = "+"
			} else {
				znak = "*"
			}
			if tryb_symulacji == GADATLIWY {
				fmt.Println("Pracownik", id, "wzial z listy zadanie:", z.liczba1, znak, z.liczba2)
			}
			sukces = false

			if z.dzial == SUMA {
				for sukces == false {
					prac_sumer <- z
					sukces = <-sukces_sumer
					if sukces {
						z = <-sumer_prac

						//fmt.Println("Pracownik", id, "odebral zadanie:", z.liczba1, znak, z.liczba2, "=", z.wynik, "od sumera")
					} else {
						//fmt.Println("Pracownik", id, "czeka na naprawe sumera")
					}
				}

				//fmt.Println("OBLICZONO", z.liczba1, z.liczba2, z.wynik)
			} else {
				for sukces == false {
					select {
					case pomoz <- true:
						//fmt.Println("Pracownik", id, "pomaga przy iloczynerze.")
					case <-time.After(time.Millisecond * time.Duration(interwal/2)):
					}

					select {
					case prac_iloczyner <- z:
					case <-time.After(time.Millisecond * time.Duration(interwal/2)):
						continue
					}
					//fmt.Println(id)
					sukces = <-sukces_iloczyner
					if sukces {
						z = <-iloczyner_prac

						//fmt.Println("Pracownik", id, "odebral zadanie:", z.liczba1, znak, z.liczba2, "=", z.wynik, "od iloczynera")
					} else {
						//fmt.Println("Pracownik", id, "czeka na naprawe iloczynera")
					}
				}

				//fmt.Println("OBLICZONO", z.liczba1, z.liczba2, z.wynik)
			}

			if tryb_symulacji == GADATLIWY {
				fmt.Println("Pracownik", id, "wykonal produkt:", z.liczba1, znak, z.liczba2, "=", z.wynik)
			}

			sukces = false

			for sukces == false {
				prac_magazyn <- z.wynik
				sukces = <-magazyn_prac

				if tryb_symulacji == GADATLIWY {
					if sukces {
						fmt.Println("Pracownik", id, "dodal do magazynu:", z.wynik)
					} else {
						fmt.Println("Pracownik", id, "probowal dodac do magazynu:", z.wynik, "jednak magazyn jest pelny")
					}
				}
				time.Sleep(time.Millisecond * time.Duration(opoznienie_pracownikow))
			}

		case <-time.After(time.Millisecond * time.Duration(opoznienie_pracownikow)):
			if tryb_symulacji == GADATLIWY {
				//fmt.Println("Pracownik", id, "probowal wziac zadanie z listy, jednak lista jest pusta")
			}
		}
	}
}
