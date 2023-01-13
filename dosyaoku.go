package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	dosyayaYaz2("okunacakDosya.txt", "beni oku!\n")

	file, err := os.Open("okunacakDosya.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func dosyayaYaz2(dosyaIsmi string, yazi string) {
	f, err := os.OpenFile(dosyaIsmi, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(yazi); err != nil {
		panic(err)
	}
}
