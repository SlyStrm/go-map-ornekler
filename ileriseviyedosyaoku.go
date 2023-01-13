package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	fmt.Println("|İSSİM|- |YAŞ| - |TAKIM|")
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "-")

		fmt.Printf("| %s | %s | %s |\n", arr[0], arr[1], arr[2])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
