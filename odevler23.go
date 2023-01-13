package main

import (
	"fmt"
	"os"
)

const CikolataFiyata = 10
const CipsFiyata = 5
const KolaFiyata = 3

func main() {

	stok := make(map[string]int)
	stok["cikolata"] = 10
	stok["kola"] = 10
	stok["cips"] = 10

	urunler := make(map[string]int)

	var girilenUrun string
	var girilenAdet int
	var cikis string

	for {
		println("lütfen hangi ürünü almak  istediğini girin")
		fmt.Scan(&girilenUrun)
		if girilenUrun == "ÇİKOLATA" || girilenUrun == "CİKOLATA" {
			girilenUrun = "cikolata"
		}

		if girilenUrun == "KOLA" || girilenUrun == "Kola" {
			girilenUrun = "kola"
		}

		if girilenUrun == "CİPS" || girilenUrun == "Cips" {
			girilenUrun = "cips"
		}

		if girilenUrun != "cips" && girilenUrun != "kola" && girilenUrun != "cikolata" {
			fmt.Println("yalnızca cips,kola ve cikolata alabilirsiniz")
			continue
		}

		println("lütfen kaç adet  istediğini girin")
		fmt.Scan(&girilenAdet)
		if stok[girilenUrun] < (urunler[girilenUrun] + girilenAdet) {
			fmt.Printf(" en fazla %d adet %s alabilirsiniz", stok[girilenUrun], girilenUrun)
			continue
		}

		println("başka bir ürün istermisiniz")
		fmt.Scan(&cikis)

		urunler[girilenUrun] = urunler[girilenUrun] + girilenAdet

		if cikis == "Hayır" {
			break
		}

	}

	toplam := 0

	for hangiUrun, Adet := range urunler {

		if hangiUrun == "cikolata" {

			if girilenAdet < 3 {
				toplam += CikolataFiyata * Adet
			} else {
				toplam += CikolataFiyata*Adet - ((CikolataFiyata*Adet)*10)/100

				dosyayaYaz("cikolata.txt", "cikolataya kampanya uygulandı\n")
			}

		}

		if hangiUrun == "kola" {
			if girilenAdet < 3 {
				toplam += KolaFiyata * Adet
			} else {
				toplam += KolaFiyata*Adet - ((KolaFiyata*Adet)*5)/100

				dosyayaYaz("kola.txt", "kolaya kampanya uygulandı\n")
			}
		}

		if hangiUrun == "cips" {
			if girilenAdet < 3 {
				toplam += CipsFiyata * Adet
			} else {
				toplam += CipsFiyata*Adet - ((CipsFiyata*Adet)*15)/100

				dosyayaYaz("cips.txt", "cipse kampanya uygulandı\n")
			}
		}

	}

	fmt.Printf(" toplam ödenecek miktar %d TL", toplam)
}

func dosyayaYaz(dosyaIsmi string, yazi string) {
	f, err := os.OpenFile(dosyaIsmi, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(yazi); err != nil {
		panic(err)
	}
}
