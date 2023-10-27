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
	s.Output = append(s.Output, s.StartElement)
	returnedWord := s.FindNextWord(s.StartElement)

	for {
		if returnedWord == s.EndElement {
			s.Output = append(s.Output, returnedWord)
			break
		}
		s.Output = append(s.Output, returnedWord)
		returnedWord = s.FindNextWord(returnedWord)
	}
	return s.Output, nil
}

func (s *Solver) FindNextWord(CurrentWord string) string {
	var potentialWords []string
	for _, st := range s.Dictionary {
		if s.CheckIfWordIsValid(st) {
			matchingLetters := 0
			for i, _ := range CurrentWord {
				if strings.Contains(st, string(CurrentWord[i])) && (string(st[i]) == string(CurrentWord[i])) {
					matchingLetters++
				}
			}
			if matchingLetters >= len(CurrentWord)-1 {
				potentialWords = append(potentialWords, st)
			}
		}
	}
	if len(potentialWords) == 1 {
		return potentialWords[0]
	} else {
		wordValues := map[string]int{}
		for _, s := range potentialWords {
			wordValues[s] = 0
		}

		for i, _ := range s.EndElement {
			for i2, _ := range wordValues {
				if string(i2[i]) == string(s.EndElement[i]) {
					wordValues[i2]++
				}
			}
		}

		max := 0
		returnableValue := ""
		for s, i := range wordValues {
			if i > max {
				returnableValue = s
				max = i
			}
		}

		return returnableValue
	}
}

func (s *Solver) CheckIfWordIsValid(newWord string) bool {
	for _, v := range s.Output {
		if v == newWord {
			return false
		}
	}
	if len(newWord) != len(s.StartElement) {
		return false
	}
	return true
}
