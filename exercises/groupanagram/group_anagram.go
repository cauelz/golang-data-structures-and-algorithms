package groupanagram

import "fmt"


func groupAnagrams(strs []string) [][]string {

	if len(strs) == 0 || len(strs) == 1 {
        return [][]string{strs}
    }

    mapper := make(map[string][]string)

    for _, s := range strs {
        count := make([]int, 26)

        for _, char := range s {
            count[char - 'a']++
        }

        key := ""
        for _, c := range count {
            key += fmt.Sprintf("%d#", c)
        }

        mapper[key] = append(mapper[key], s)
    }

    result := make([][]string, 0, len(mapper))
    for _, group := range mapper {
        result = append(result, group)
    }

    return result
}