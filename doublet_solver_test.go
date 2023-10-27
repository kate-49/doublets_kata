package doublets_kata

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_E2E_Simple_Inputs(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput []string
	}{
		//{Input: "head tail", ExpectedOutput: []string{"head", "heal", "teal", "tell", "tall", "tail"}},
		{Input: "door lock", ExpectedOutput: []string{"door", "boor", "book", "look", "lock"}},
		//{Input: "bank loan", ExpectedOutput: []string{"bank", "bonk", "book", "look", "loon", "loan"}},
		//{Input: "wheat bread", ExpectedOutput: []string{"wheat", "cheat", "cheap", "creep", "creed", "breed", "bread"}},
	}
	for k, tc := range tests {
		t.Run("scenario "+strconv.Itoa(k+1)+" given:"+tc.Input, func(t *testing.T) {
			solver, err := CreateSolver(tc.Input)
			assert.Nil(t, err)
			returnedValue, returnedErr := solver.Run()
			assert.Equal(t, tc.ExpectedOutput, returnedValue)
			assert.Nil(t, returnedErr)
		})
	}
}
