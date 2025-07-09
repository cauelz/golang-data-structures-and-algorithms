# Group Anagrams Solution Analysis - FAANG Interview Optimization

## Problem Overview
Given an array of strings `strs`, group the anagrams together. An anagram is a word formed by rearranging the letters of another word.

**Example:**
- Input: `["eat","tea","tan","ate","nat","bat"]`
- Output: `[["bat"],["nat","tan"],["ate","eat","tea"]]`

## Why This Solution is FAANG-Optimized

### 1. **Optimal Time Complexity: O(n * k)**
- Where `n` = number of strings, `k` = average length of strings
- This is the best possible time complexity for this problem
- FAANG interviewers expect you to recognize and achieve optimal complexity

### 2. **Optimal Space Complexity: O(n * k)**
- Space needed to store all strings in the result
- This is the minimum space required for any solution

### 3. **Elegant Algorithm Design**
- Uses a clever key generation technique
- Avoids sorting (which would add O(k log k) per string)
- Demonstrates understanding of hash maps and character counting

## Line-by-Line Breakdown

### Edge Case Handling (Lines 8-10)
```go
if len(strs) == 0 || len(strs) == 1 {
    return [][]string{strs}
}
```
**What it does:** Handles edge cases efficiently
**Why it works:** 
- Empty array: Returns `[[]]` (empty group)
- Single element: Returns `[[element]]` (single group)
- **FAANG Bonus:** Shows you think about edge cases first

### Data Structure Setup (Line 12)
```go
mapper := make(map[string][]string)
```
**What it does:** Creates a hash map to group anagrams
**Why it works:**
- Key: Character frequency signature
- Value: List of strings with that signature
- **FAANG Bonus:** Demonstrates understanding of hash map operations (O(1) average case)

### Main Algorithm Loop (Lines 14-30)
```go
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
```

#### Character Counting (Lines 15-19)
```go
count := make([]int, 26)
for _, char := range s {
    count[char - 'a']++
}
```
**What it does:** Counts frequency of each lowercase letter
**Why it works:**
- `char - 'a'` converts 'a'→0, 'b'→1, ..., 'z'→25
- `count[i]` stores frequency of character at position i
- **FAANG Bonus:** Shows ASCII manipulation knowledge

#### Key Generation (Lines 21-24)
```go
key := ""
for _, c := range count {
    key += fmt.Sprintf("%d#", c)
}
```
**What it does:** Creates a unique signature for each anagram group
**Why it works:**
- Example: "eat" → count[4]=1, count[0]=1, count[19]=1 → "0#0#0#0#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#1#0#0#0#0#0#0#"
- All anagrams of "eat" will produce the same key
- **FAANG Bonus:** Creative solution that avoids sorting

#### Grouping (Line 26)
```go
mapper[key] = append(mapper[key], s)
```
**What it does:** Adds current string to its anagram group
**Why it works:**
- Strings with same key are anagrams
- **FAANG Bonus:** Efficient hash map usage

### Result Construction (Lines 32-35)
```go
result := make([][]string, 0, len(mapper))
for _, group := range mapper {
    result = append(result, group)
}
return result
```
**What it does:** Converts map values to required output format
**Why it works:**
- Pre-allocates slice capacity for efficiency
- **FAANG Bonus:** Shows attention to memory optimization

## Time & Space Complexity Analysis

### Time Complexity: O(n * k)
- **n**: Number of strings in input array
- **k**: Average length of strings
- **Breakdown:**
  - Outer loop: O(n) - iterate through each string
  - Character counting: O(k) - count characters in each string
  - Key generation: O(26) = O(1) - fixed size alphabet
  - Hash map operations: O(1) average case
  - Result construction: O(n) - iterate through map values
- **Total:** O(n * k) + O(n) = O(n * k)

### Space Complexity: O(n * k)
- **Input storage:** O(n * k) - storing all input strings
- **Character count arrays:** O(n * 26) = O(n) - one array per string
- **Hash map:** O(n * k) - storing all strings grouped by key
- **Output:** O(n * k) - storing all strings in result
- **Total:** O(n * k)

## Why This Beats Alternative Approaches

### 1. **vs. Sorting Approach**
- **Sorting:** O(n * k log k) time complexity
- **This solution:** O(n * k) time complexity
- **FAANG Advantage:** Demonstrates you can optimize beyond obvious solutions

### 2. **vs. Prime Number Multiplication**
- **Prime approach:** Risk of integer overflow with long strings
- **This solution:** No overflow risk, handles any string length
- **FAANG Advantage:** Shows consideration of practical limitations

### 3. **vs. Character-by-Character Comparison**
- **Comparison:** O(n² * k) time complexity
- **This solution:** O(n * k) time complexity
- **FAANG Advantage:** Demonstrates algorithmic thinking

## FAANG Interview Tips Demonstrated

1. **Edge Case Handling:** Lines 8-10 show you think about corner cases
2. **Optimal Complexity:** Achieves best possible time/space complexity
3. **Clean Code:** Readable, well-structured solution
4. **Efficient Data Structures:** Smart use of hash maps
5. **Memory Optimization:** Pre-allocates slice capacity
6. **Creative Problem Solving:** Unique key generation approach

## Common Follow-up Questions

1. **"What if strings contain uppercase letters?"**
   - Modify `char - 'a'` to handle both cases
   
2. **"What if strings contain non-alphabetic characters?"**
   - Use a larger count array or different key generation

3. **"Can you optimize the key generation?"**
   - Use a more compact representation or different encoding

This solution demonstrates the level of algorithmic thinking and optimization skills that FAANG companies expect from their candidates. 