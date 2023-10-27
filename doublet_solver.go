package doublets_kata

import (
	"bufio"
	"errors"
	"fmt"
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
	returnedWord, err := s.FindNextWord(s.StartElement)
	if err != nil {
		return nil, err
	}

	for {
		if returnedWord == s.EndElement {
			s.Output = append(s.Output, returnedWord)
			break
		}
		s.Output = append(s.Output, returnedWord)
		fmt.Println("word")
		fmt.Println(returnedWord)
		returnedWord, err = s.FindNextWord(returnedWord)
		if err != nil {
			return nil, err
		}
	}
	return s.Output, nil
}

func (s *Solver) FindNextWord(CurrentWord string) (string, error) {
	var potentialWords []string
	for _, st := range s.Dictionary {
		if s.CheckIfWordIsValid(st) {
			matchingLetters := 0
			for i, _ := range CurrentWord {
				if strings.Contains(st, string(CurrentWord[i])) && (string(st[i]) == string(CurrentWord[i])) {
					matchingLetters++
				}
			}
			if matchingLetters >= 3 {
				potentialWords = append(potentialWords, st)
			}
		}
	}
	if len(potentialWords) == 1 {
		return potentialWords[0], nil
	} else {
		fmt.Println(potentialWords)
		return "", errors.New("more than one option found for next word")
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
