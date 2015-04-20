// interfejs
package main

import (
	"fmt"
	"strings"
)

func wypisz_polecenia() {
	fmt.Println("l - lista zadan")
	fmt.Println("m - magazyn")
	fmt.Println("k - koniec")
}

func interfejs() {
	for {
		wypisz_polecenia()
		var znak string
		fmt.Scan(&znak)

		if strings.EqualFold(znak, "l") {
			fmt.Println("Ilosc zadan na liscie to", ilosc_zadan)
		} else if strings.EqualFold(znak, "m") {
			fmt.Println("Ilosc produktow w magazynie to", ilosc_produktow)
		} else if strings.EqualFold(znak, "k") {
			return
		}
	}
}
