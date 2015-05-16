package main

import (
	"fmt"
	"math/rand"
	"time"
)

func iloczyner(id int, prac_iloczyner <-chan zadanie, sukces_iloczyner chan<- bool, iloczyner_prac chan<- zadanie,
	iloczyner_serwisant chan<- bool, serwisant_iloczyner <-chan bool, pomoz <-chan bool) {
	var losowe_pstwo int
	var z zadanie
	var sukces bool
	var zepsuta bool

	for {
		z = <-prac_iloczyner
		<-pomoz
		//fmt.Println("ŁĄCZY SIĘ ILOCZYNER", id)

		time.Sleep(time.Millisecond * time.Duration(opoznienie_iloczynera))

		losowe_pstwo = rand.Intn(100) + 1
		if losowe_pstwo <= psucie_iloczynera {
			sukces = false
			zepsuta = true
			sukces_iloczyner <- sukces

			if tryb_symulacji == GADATLIWY {
				fmt.Println("Iloczyner", id, "się zepsuł.")
			}
		} else {
			sukces = true
			zepsuta = false
			z.wynik = z.liczba1 * z.liczba2
			sukces_iloczyner <- sukces
			iloczyner_prac <- z
		}

		if zepsuta {
			zepsuta = false
			iloczyner_serwisant <- true
			<-serwisant_iloczyner
			if tryb_symulacji == GADATLIWY {
				fmt.Println("Iloczyner", id, "został naprawiony przez serwisanta.")
			}
		}
	}
}
