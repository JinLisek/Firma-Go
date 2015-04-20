// magazyn
package main

import (
	"time"
)

func magazyn(prac_magazyn <-chan int, magazyn_prac chan<- bool, magazyn_klient chan<- int) {
	var produkt int

	produkty = make([]int, pojemnosc_magazynu)

	for {
		select {
		case produkt = <-prac_magazyn:
			if ilosc_produktow != pojemnosc_magazynu {
				produkty[ilosc_produktow] = produkt
				ilosc_produktow++
				magazyn_prac <- true
			} else {
				magazyn_prac <- false
			}
		case <-time.After(time.Millisecond * time.Duration(interwal/2)):
		}

		if ilosc_produktow != 0 {
			select {
			case magazyn_klient <- produkty[ilosc_produktow-1]:
				ilosc_produktow--
			case <-time.After(time.Millisecond * time.Duration(interwal/2)):
			}

		}
	}

}
