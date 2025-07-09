# Contains Duplicate – Solution Discussion

## Solution Review

```go
func ContainsDuplicate(nums []int) bool {
    mapper := make(map[int]int)
    for _, value := range nums {
        if _, ok := mapper[value]; ok {
            return true
        }
        mapper[value]++
    }
    return false
}
```

## Analysis

### 1. Time Complexity
- Each element is checked and inserted into the map in O(1) average time.
- **Total time complexity:** O(n), where n is the length of `nums`.

### 2. Space Complexity
- In the worst case (no duplicates), every element is stored in the map.
- **Total space complexity:** O(n).

### 3. Correctness
- The function returns `true` as soon as a duplicate is found.
- If no duplicates are found, it returns `false`.
- The logic is correct and efficient.

### 4. Optimization
- This is the optimal solution for this problem in terms of time and space for unsorted data.
- Minor Go-specific improvement: use `map[int]struct{}` instead of `map[int]int` since only existence matters, not counts. This saves a small amount of memory.

#### Example:
```go
mapper := make(map[int]struct{})
for _, value := range nums {
    if _, ok := mapper[value]; ok {
        return true
    }
    mapper[value] = struct{}{}
}
```

### 5. Alternative Approaches
- **Sorting:** Sort the array and check for consecutive duplicates. Time: O(n log n), Space: O(1) or O(n) depending on the sort implementation. Not as efficient as the hash map approach for large n.
- **Set:** In languages with a built-in set type, use a set for existence checks (Go uses map for this).

### 6. FAANG Interview Readiness
- This solution is the standard, optimal approach expected in FAANG interviews for this problem.
- It demonstrates knowledge of hash maps and efficient duplicate detection.
- You can mention the alternative sorting approach and explain why you chose the hash map for O(n) time.

---

## Summary
- **Is this solution optimal for FAANG?**  
  **Yes, it is optimal and interview-ready.**
- **Minor improvement:**  
  Use `map[int]struct{}` for slightly less memory usage (not critical).

Your current solution is perfectly acceptable and will impress interviewers. If you want to discuss further optimizations or edge cases, feel free to ask! 