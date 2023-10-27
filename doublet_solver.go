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
	secondWord, err := s.FindNextWord(s.StartElement)
	if err != nil {
		return nil, err
	}

	s.Output = append(s.Output, secondWord)
	fmt.Println("second word")
	fmt.Println(secondWord)
	thirdWord, err := s.FindNextWord(secondWord)
	if err != nil {
		return nil, err
	}

	s.Output = append(s.Output, thirdWord)
	fmt.Println("third word")
	fmt.Println(thirdWord)

	fourthWord, err := s.FindNextWord(thirdWord)
	if err != nil {
		return nil, err
	}

	s.Output = append(s.Output, fourthWord)
	fmt.Println("fourth word")
	fmt.Println(fourthWord)

	fifthWord, err := s.FindNextWord(fourthWord)
	if err != nil {
		return nil, err
	}

	s.Output = append(s.Output, fifthWord)
	fmt.Println("fifthWord word")
	fmt.Println(fifthWord)

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
		return "", errors.New("more than one option found for next word")
	}
}

func (s *Solver) CheckIfWordIsValid(newWord string) bool {
	for _, v := range s.Output {
		if v == newWord {
			return false
		}
	}
	return true
}
