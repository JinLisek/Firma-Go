package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumer(id int, prac_sumer <-chan zadanie, sukces_sumer chan<- bool, sumer_prac chan<- zadanie, sumer_serwisant chan<- bool, serwisant_sumer <-chan bool) {
	var losowe_pstwo int
	var z zadanie
	var sukces bool
	var zepsuta bool

	for {
		z = <-prac_sumer
		//fmt.Println("ŁĄCZY SIĘ SUMER", id)

		time.Sleep(time.Millisecond * time.Duration(opoznienie_sumera))

		losowe_pstwo = rand.Intn(100) + 1
		if losowe_pstwo <= psucie_sumera {
			sukces = false
			zepsuta = true
			sukces_sumer <- sukces

			if tryb_symulacji == GADATLIWY {
				fmt.Println("Sumer", id, "się zepsuł.")
			}
		} else {
			sukces = true
			zepsuta = false
			z.wynik = z.liczba1 + z.liczba2
			sukces_sumer <- sukces
			sumer_prac <- z
		}

		if zepsuta {
			zepsuta = false
			sumer_serwisant <- true
			<-serwisant_sumer
			if tryb_symulacji == GADATLIWY {
				fmt.Println("Sumer", id, "został naprawiony przez serwisanta.")
			}
		}
	}
}
