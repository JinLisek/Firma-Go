// lista_zadan
package main

import (
	"time"
)

func lista_zadan(prez_lista <-chan zadanie, lista_prez chan<- bool, lista_prac chan<- zadanie) {
	var z zadanie
	zadania = make([]zadanie, pojemnosc_listy)

	for {
		select {
		case z = <-prez_lista:
			if ilosc_zadan != pojemnosc_listy {
				zadania[ilosc_zadan] = z
				ilosc_zadan++
				lista_prez <- true
			} else {
				lista_prez <- false
			}
		case <-time.After(time.Millisecond * time.Duration(interwal/2)):
		}

		if ilosc_zadan != 0 {
			select {
			case lista_prac <- zadania[ilosc_zadan-1]:
				ilosc_zadan--
			case <-time.After(time.Millisecond * time.Duration(interwal/2)):
			}

		}
	}
}
