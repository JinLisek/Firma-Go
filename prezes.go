// prezes.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func prezes(prez_lista chan<- zadanie, lista_prez <-chan bool) {
	var z zadanie
	var dzialanieLos int
	var znak string
	var sukces bool

	for {
		dzialanieLos = rand.Intn(2)

		z.liczba1 = rand.Intn(200) - 100
		z.liczba2 = rand.Intn(200) - 100

		if dzialanieLos == 0 {
			z.dzial = SUMA
			znak = "+"
		} else {
			z.dzial = ILOCZYN
			znak = "*"
		}

		sukces = false

		for sukces == false {
			prez_lista <- z
			sukces = <-lista_prez

			if tryb_symulacji == GADATLIWY {
				if sukces {
					fmt.Println("Prezes dodal do listy zadan:", z.liczba1, znak, z.liczba2)
				} else {
					//fmt.Println("Prezes probowal dodac do listy zadan:", z.liczba1, znak, z.liczba2, "jednak lista jest pelna")
				}
			}

			time.Sleep(time.Millisecond * time.Duration(opoznienie_prezesa))
		}

	}
}
