package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("cikolata.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	toplamKampanyaSayisi := 0

	for scanner.Scan() {
		toplamKampanyaSayisi++
	}

	fmt.Printf("Cikolata için toplam yapılan kampanya sayısı %d adet.", toplamKampanyaSayisi)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
