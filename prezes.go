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

	for {
		dzialanieLos = rand.Intn(3)

		z.liczba1 = rand.Intn(200) - 100
		z.liczba2 = rand.Intn(200) - 100

		if dzialanieLos == 0 {
			if z.liczba2 >= 0 {
				z.dzial = suma
				znak = "+"
			} else {
				z.liczba2 *= -1
				z.dzial = roznica
				znak = "-"
			}

		} else if dzialanieLos == 1 {
			if z.liczba2 >= 0 {
				z.dzial = roznica
				znak = "-"
			} else {
				z.liczba2 *= -1
				z.dzial = suma
				znak = "+"
			}
		} else {
			z.dzial = iloczyn
			znak = "*"
		}

		prez_lista <- z
		var sukces bool = <-lista_prez

		if tryb_symulacji == GADATLIWY {
			if sukces {
				fmt.Println("Prezes dodal do listy zadan:", z.liczba1, znak, z.liczba2)
			} else {
				fmt.Println("Prezes probowal dodac do listy zadan:", z.liczba1, znak, z.liczba2, "jednak lista jest pelna")
			}
		}

		time.Sleep(time.Millisecond * time.Duration(opoznienie_prezesa))
	}
}
