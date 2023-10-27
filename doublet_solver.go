package doublets_kata

import (
	"bufio"
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
			if matchingLetters >= 3 {
				potentialWords = append(potentialWords, st)
			}
		}
	}
	if len(potentialWords) == 1 {
		return potentialWords[0]
	} else {
		fmt.Println(potentialWords)
		fmt.Println(potentialWords[0])
		fmt.Println(potentialWords[1])

		word1 := 0
		word2 := 0
		for i, _ := range s.EndElement {
			if string(potentialWords[0][i]) == string(s.EndElement[i]) {
				word1++
			}
			if string(potentialWords[1][i]) == string(s.EndElement[i]) {
				word2++
			}
		}
		if word1 > word2 {
			return potentialWords[0]
		} else {
			return potentialWords[1]
		}
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

func (s *Solver) CheckWhichWordIsCloserToFinalWord(potentialWords []string) string {
	//words := map[string]int{}
	//matchingLetters := 0
	fmt.Println("potential words")
	fmt.Println(potentialWords)
	//for i, _ := range potentialWords {
	//	fmt.Println("word")
	//	fmt.Println(potentialWords[i])
	//	//for _, letter := range s.EndElement {
	//	//	if strings.Contains(word, string(letter)) {
	//	//		fmt.Println("element")
	//	//		fmt.Println(string(letter))
	//	//		matchingLetters++
	//	//	}
	//	//}
	//	//words[word] = matchingLetters
	//}
	//fmt.Println(words)
	return "test"
}
