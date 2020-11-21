package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	// CSV Dateiname als Flag realisieren
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	flag.Parse()

	// Öffnen der Datei
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Fehler beim Öffnen der Datei: %s", *csvFileName))
	}

	// Parsen der csv-Datei
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Fehler beim Parsen der Datei.")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			fmt.Printf("Korrekt\n")
			correct++
		} else {
			fmt.Printf("Falsch: %s\n", answer)
		}
	}

	fmt.Printf("Du hast %d Fragen von %d richtig beantwortet", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
