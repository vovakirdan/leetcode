package oddevenlinkedlist

import (
	"reflect"
	"testing"
)

// Хелпер для создания списка из слайса
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// Хелпер для преобразования списка в слайс
func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5, 2, 4},
		},
		{
			name:     "Example 2",
			input:    []int{2, 1, 3, 5, 6, 4, 7},
			expected: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Two elements",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Three elements",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "All even values but odd indices",
			input:    []int{2, 4, 6, 8, 10},
			expected: []int{2, 6, 10, 4, 8},
		},
		{
			name:     "Negative values",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-1, -3, -2, -4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			result := oddEvenList(head)
			resultSlice := listToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, resultSlice)
			}
		})
	}
}
