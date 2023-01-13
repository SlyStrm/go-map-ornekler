package main

import (
	"fmt"
)

func main() {

	baraj := make(map[string]int)
	baraj["matematik"] = 50
	baraj["fizik"] = 45
	baraj["tarih"] = 40

	Notlar := make(map[string]int)

	var girilenDers string
	var girilenNot int
	var cikis string

	for {
		println("lütfen Hangi Derse Girmek  İstediğini girin")
		fmt.Scan(&girilenDers)

		if girilenDers == "Matematik" || girilenDers == "MATEMATİK" {
			girilenDers = "matematik"
		}

		if girilenDers == "fizik" || girilenDers == "FİZİK" {
			girilenDers = "fizik"
		}

		if girilenDers == "Tarih" || girilenDers == "TARİH" {
			girilenDers = "tarih"

		}

		if girilenDers != "matematik" && girilenDers != "fizik" && girilenDers != "tarih" {
			fmt.Println("yalnızca Seçili Dersler Girilebilir")
			continue
		}

		println("lütfen Notunuzu Girin")

		fmt.Scan(&girilenNot)

		if girilenDers == "Matematik" && girilenNot >= 80 {
			println("Helals")
		}

		if baraj[girilenDers] > (Notlar[girilenDers] + girilenNot) {
			fmt.Printf("geçmek için en az %d notunu %s dersinden almalısınız ", baraj[girilenDers], girilenDers)
			continue
		}

		println("Başka Derse Girmek istermisiniz")
		fmt.Scan(&cikis)

		Notlar[girilenDers] = Notlar[girilenDers] + girilenNot

		if cikis == "Hayır" {
			break
		}

	}

	toplam := 0

	for hangiUrun, not := range Notlar {

		if hangiUrun == "Matematik" {

			if girilenNot < 80 {
				toplam += not
			} else {
				println("Matematik dersiniz 80 üstü yani çok iyi")
			}

		}

		if hangiUrun == "Fizik" {
			toplam += not
		}

		if hangiUrun == "Tarih" {
			toplam += not
		}

	}

	for Ders, Not := range Notlar {
		fmt.Printf("%s dersinizin notu =  %d\n", Ders, Not)

	}
}
