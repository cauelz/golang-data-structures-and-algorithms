package groupanagram

import (
	"reflect"
	"sort"
	"testing"
)

// Helper to sort groups and their contents for comparison
func sortGroups(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 || len(groups[j]) == 0 {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		input []string
		expect [][]string
	}{
		{
			name:   "Example 1",
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expect: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:   "Example 2",
			input:  []string{""},
			expect: [][]string{{""}},
		},
		{
			name:   "Example 3",
			input:  []string{"a"},
			expect: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.input)
			sortGroups(got)
			sortGroups(tt.expect)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("got %v, want %v", got, tt.expect)
			}
		})
	}
} 