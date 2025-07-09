# Anagram Algorithm Explained

## Problem Statement
Given two strings, determine if they are anagrams of each other (i.e., they contain the same characters with the same frequencies, in any order).

---

## Optimized Solution Overview
The optimized solution uses a single map to count character frequencies, incrementing for one string and decrementing for the other in a single pass. If all counts are zero at the end, the strings are anagrams.

---

## Algorithm Steps

1. **Length Check**
   - If the strings have different lengths, they cannot be anagrams.
   - `if len(s) != len(t) { return false }`

2. **Initialize Map**
   - Create a map to count occurrences of each character (as `rune`).
   - `count := make(map[rune]int)`

3. **Single Pass Counting**
   - Loop through both strings simultaneously:
     - Increment the count for the character in `s`.
     - Decrement the count for the character in `t`.
   - Example:
     ```go
     for i, char := range s {
         count[char]++
         count[rune(t[i])]--
     }
     ```

4. **Final Check**
   - If all values in the map are zero, the strings are anagrams.
   - Otherwise, return false.
   - Example:
     ```go
     for _, v := range count {
         if v != 0 {
             return false
         }
     }
     return true
     ```

---

## Space and Time Complexity Analysis

### Time Complexity
- **O(n)**, where n is the length of the strings.
  - We loop through each string once (single pass).
  - The final check loops through the map, which has at most O(n) entries (in the worst case, all unique characters).

### Space Complexity
- **O(1)** if the character set is fixed (e.g., ASCII), because the map size does not grow with input size.
- **O(n)** in the worst case for arbitrary Unicode strings, where all characters are unique.

---

## Why Use Runes Instead of Bytes?

- **Go strings are sequences of bytes**, but many characters (especially non-English) are represented by multiple bytes (UTF-8 encoding).
- **A `rune` in Go is an alias for `int32` and represents a Unicode code point** (a full character, regardless of how many bytes it takes).
- Using `range s` gives you runes, not bytes, so you correctly handle all Unicode characters.
- If you use bytes (`s[i]`), you may split multi-byte characters and get incorrect results for non-ASCII input.

### Example
```go
s := "café"
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i]) // Prints: c a f Ã ©
}
for _, r := range s {
    fmt.Printf("%c ", r) // Prints: c a f é
}
```
- The first loop prints incorrect characters for 'é' because it is two bytes in UTF-8.
- The second loop prints the correct characters.

---

## Summary
- Always use runes when processing characters in Go strings to handle Unicode correctly.
- The optimized anagram algorithm is efficient and suitable for interviews, with O(n) time and O(1)/O(n) space depending on the character set. 