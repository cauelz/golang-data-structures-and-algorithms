package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"anagrams", "anagram", "nagaram", true},
		{"not anagrams", "rat", "car", false},
		{"empty strings", "", "", true},
		{"different lengths", "a", "ab", false},
		{"unicode anagrams", "àççéñt", "tçéñàç", true},
		{"unicode not anagrams", "àççéñt", "accent", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v; want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
} 