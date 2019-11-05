package main

import "testing"

func TestInsertStringSlice(t *testing.T) {
	var insertSlice = InsertStringSlice([]int{1}, []int{5, 5, 5, 5, 5}, 2)
	var validSlice = []int{5, 5, 1, 5, 5, 5}
	if len(insertSlice) != len(validSlice) {
		t.Error(`InsertStringSlice error`)
	}
	noEquals := true
	for k, v := range insertSlice {
		if v != validSlice[k] {
			noEquals = false
		}
	}
	if !noEquals {
		t.Error(`InsertStringSlice error`)
	}

}

func TestRemoveStringSlice(t *testing.T) {
	var removeSlice = RemoveStringSlice([]string{"a", "b", "c", "d", "e"}, 1, 2)
	var validSlice = []string{"a", "c", "d", "e"}
	if len(removeSlice) != len(validSlice) {
		t.Error(`removeSlice error`)
	}
	noEquals := true
	for k, v := range removeSlice {
		if v != validSlice[k] {
			noEquals = false
		}
	}
	if !noEquals {
		t.Error(`removeSlice error`)
	}
}
