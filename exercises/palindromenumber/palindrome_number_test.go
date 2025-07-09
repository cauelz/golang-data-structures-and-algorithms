package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"positive palindrome", 121, true},
		{"negative number", -121, false},
		{"not a palindrome", 10, false},
		{"negative palindrome-like", -101, false},
		{"single digit", 7, true},
		{"zero", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.want {
				t.Errorf("isPalindrome(%d) = %v; want %v", tt.x, got, tt.want)
			}
		})
	}
} 