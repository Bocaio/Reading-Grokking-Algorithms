package ch05hashtables

import "testing"

func TestGetFromEmptyTable(t *testing.T) {
	table := CreateHashTable()

	got, success := GetHashTable("apple", table)

	if success {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestSingleInsert(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("apple", "fruit", table)

	got, _ := GetHashTable("apple", table)

	if got != "fruit" {
		t.Errorf("expected fruit, got %q", got)
	}
}

func TestMultipleInserts(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("apple", "fruit", table)
	AddHashTable("carrot", "vegetable", table)
	AddHashTable("salmon", "fish", table)

	tests := []struct {
		key      string
		expected string
	}{
		{"apple", "fruit"},
		{"carrot", "vegetable"},
		{"salmon", "fish"},
	}

	for _, tt := range tests {
		got, _ := GetHashTable(tt.key, table)

		if got != tt.expected {
			t.Errorf("key=%s expected=%s got=%s",
				tt.key,
				tt.expected,
				got,
			)
		}
	}
}

func TestCollisionHandling(t *testing.T) {
	table := CreateHashTable()

	// hash("ab") = (97 + 98) % 10 = 5
	// hash("ba") = (98 + 97) % 10 = 5

	AddHashTable("ab", "first", table)
	AddHashTable("ba", "second", table)

	if got, _ := GetHashTable("ab", table); got != "first" {
		t.Errorf("expected first, got %q", got)
	}

	if got, _ := GetHashTable("ba", table); got != "second" {
		t.Errorf("expected second, got %q", got)
	}
}

func TestMissingKey(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("apple", "fruit", table)

	got, success := GetHashTable("banana", table)

	if success {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestLongCollisionChain(t *testing.T) {
	table := CreateHashTable()

	keys := []string{
		"ab",
		"ba",
		"aab",
		"aba",
		"baa",
	}

	for i, key := range keys {
		AddHashTable(key, string(rune('A'+i)), table)
	}

	for i, key := range keys {
		expected := string(rune('A' + i))

		got, _ := GetHashTable(key, table)

		if got != expected {
			t.Errorf(
				"key=%s expected=%s got=%s",
				key,
				expected,
				got,
			)
		}
	}
}

func TestDeleteHeadOfChain(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("ab", "first", table)
	AddHashTable("ba", "second", table)

	DeleteHashTable("ab", table)

	if _, exist := GetHashTable("ab", table); exist {
		t.Errorf("ab should have been deleted")
	}

	if value, _ := GetHashTable("ba", table); value != "second" {
		t.Errorf("ba should still exist")
	}
}

func TestDeleteTailOfChain(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("ab", "first", table)
	AddHashTable("ba", "second", table)

	DeleteHashTable("ba", table)

	if _, exist := GetHashTable("ba", table); exist {
		t.Errorf("ba should have been deleted")
	}

	if value, _ := GetHashTable("ab", table); value != "first" {
		t.Errorf("ab should still exist")
	}
}

func TestDeleteOnlyNode(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("apple", "fruit", table)

	DeleteHashTable("apple", table)

	got, _ := GetHashTable("apple", table)

	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestDeleteFromEmptyTable(t *testing.T) {
	table := CreateHashTable()

	DeleteHashTable("apple", table)

	got, exist := GetHashTable("apple", table)

	if exist {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestDeleteMissingKey(t *testing.T) {
	table := CreateHashTable()

	AddHashTable("apple", "fruit", table)

	success := DeleteHashTable("banana", table)
	if success {
		t.Errorf("Deleting should be failed")
	}
}
