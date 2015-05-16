package main

import (
	"time"
)

func serwisant(sumer_serwisant <-chan bool, serwisant_sumer chan<- bool, iloczyner_serwisant <-chan bool, serwisant_iloczyner chan<- bool) {
	for {
		select {
		case <-sumer_serwisant:
			time.Sleep(time.Millisecond * time.Duration(opoznienie_serwisanta))
			serwisant_sumer <- true
		case <-iloczyner_serwisant:
			time.Sleep(time.Millisecond * time.Duration(opoznienie_serwisanta))
			serwisant_iloczyner <- true
		}

	}
}
