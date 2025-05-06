package stringcompression

import (
	"reflect"
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		expectedLen  int
		expectedData []byte
	}{
		{
			name:         "Example 1 - multiple groups",
			input:        []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expectedLen:  6,
			expectedData: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			name:         "Example 2 - single character",
			input:        []byte{'a'},
			expectedLen:  1,
			expectedData: []byte{'a'},
		},
		{
			name:         "Example 3 - group over 9",
			input:        []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expectedLen:  4,
			expectedData: []byte{'a', 'b', '1', '2'},
		},
		{
			name:         "All unique",
			input:        []byte{'a', 'b', 'c', 'd'},
			expectedLen:  4,
			expectedData: []byte{'a', 'b', 'c', 'd'},
		},
		{
			name:         "Single run of 10",
			input:        []byte{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'},
			expectedLen:  3,
			expectedData: []byte{'x', '1', '0'},
		},
		{
			name:         "Two runs of 1",
			input:        []byte{'x', 'y'},
			expectedLen:  2,
			expectedData: []byte{'x', 'y'},
		},
		{
			name:         "Single run of 100",
			input:        bytesRepeat('z', 100),
			expectedLen:  4,
			expectedData: []byte{'z', '1', '0', '0'},
		},
		{
			name:         "Mixed short and long runs",
			input:        append([]byte{'a', 'a'}, bytesRepeat('b', 12)...),
			expectedLen:  5,
			expectedData: []byte{'a', '2', 'b', '1', '2'},
		},
		{
			name:         "Ends with single character",
			input:        []byte{'a', 'a', 'b'},
			expectedLen:  3,
			expectedData: []byte{'a', '2', 'b'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]byte, len(tt.input))
			copy(inputCopy, tt.input)
			actualLen := compress(inputCopy)
			if actualLen != tt.expectedLen {
				t.Errorf("Expected length %d, got %d", tt.expectedLen, actualLen)
			}
			if !reflect.DeepEqual(inputCopy[:actualLen], tt.expectedData) {
				t.Errorf("Expected data %v, got %v", tt.expectedData, inputCopy[:actualLen])
			}
		})
	}
}

// Вспомогательная функция для повторения символа
func bytesRepeat(ch byte, count int) []byte {
	res := make([]byte, count)
	for i := range res {
		res[i] = ch
	}
	return res
}
