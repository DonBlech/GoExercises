package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	// Einlesen des CSV Datei
	fn := "problems.csv"

	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Start reading from the file with a reader.
	gesamtPunkte := 0
	reader := bufio.NewReader(file)
	fmt.Println(reader)
	var line string
	for i := 0; i < 10; i++ {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// Wandeln CSV => String-Array
		r := csv.NewReader(strings.NewReader(line))
		records, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		// Lesen von StdIn
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Frage: %s = ", records[0][0])
		antwort, _ := reader.ReadString('\n')

		// Prüfen, ob die Antwort richtig ist
		if strings.TrimSpace(antwort) == strings.TrimSpace(records[0][1]) {
			fmt.Printf("Die Anwort ist richtig:  %s \n", records[0][1])
			gesamtPunkte++
		} else {
			fmt.Printf("Die Anwort ist falsch. Richtig ist:  %s \n", records[0][1])
		}
	}

	fmt.Printf("Herzlichen Glückwunsch! Sie haben insgesamt % Punkte erreicht \n", gesamtPunkte)

}
