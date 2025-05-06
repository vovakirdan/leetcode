package guessnumberhigherorlower

import (
	"testing"
)

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		n    int
		pick int
	}{
		{n: 10, pick: 6},
		{n: 1, pick: 1},
		{n: 2, pick: 1},
		{n: 2, pick: 2},
		{n: 100, pick: 73},
		{n: 1000000, pick: 999999},
		{n: 2147483647, pick: 2147483647},
		{n: 2147483647, pick: 1},
	}

	for _, tt := range tests {
		picked = tt.pick // установить значение глобально
		result := guessNumber(tt.n)
		if result != tt.pick {
			t.Errorf("guessNumber(%d), picked = %d — expected %d, got %d", tt.n, tt.pick, tt.pick, result)
		}
	}
}
