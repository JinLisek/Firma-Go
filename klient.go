// pracownik
package main

import (
	"fmt"
	"time"
)

func klient(id int, magazyn_klient <-chan int) {
	var produkt int

	for {

		select {
		case produkt = <-magazyn_klient:
			if tryb_symulacji == GADATLIWY {
				fmt.Println("Klient", id, "wybral produkt z magazynu:", produkt)
			}
		case <-time.After(time.Millisecond * time.Duration(opoznienie_klientow)):
			if tryb_symulacji == GADATLIWY {
				//fmt.Println("Klient", id, "probowal wybrac produkt z magazunu, jednak magazyn jest pusty")
			}
		}

		time.Sleep(time.Millisecond * time.Duration(opoznienie_klientow))
	}
}
