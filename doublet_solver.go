package doublets_kata

import (
	"bufio"
	"os"
	"strings"
)

type Solver struct {
	StartElement string
	EndElement   string
	Dictionary   []string
	Output       []string
}

func CreateSolver(input string) (Solver, error) {
	s := Solver{}

	splitInput := strings.Split(input, " ")
	s.StartElement = splitInput[0]
	s.EndElement = splitInput[1]

	file, err := os.Open("dictionary.txt")
	if err != nil {
		return Solver{}, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s.Dictionary = append(s.Dictionary, strings.Trim(scanner.Text(), "\""))
	}

	err = file.Close()
	if err != nil {
		return Solver{}, err
	}
	return s, nil
}

func (s *Solver) Run() ([]string, error) {
	var potentialWords []string
	for _, st := range s.Dictionary {
		//check that 3 of the letters match current word and are in same place
		if st != s.StartElement {
			matchingLetters := 0
			for i, _ := range s.StartElement {
				if strings.Contains(st, string(s.StartElement[i])) && (string(st[i]) == string(s.StartElement[i])) {
					matchingLetters++
				}
			}
			if matchingLetters >= 3 {
				potentialWords = append(potentialWords, st)
			}
		}
	}
	return potentialWords, nil
}
