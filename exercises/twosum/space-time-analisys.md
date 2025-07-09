# Two Sum - Space-Time Complexity Analysis

## Problem Statement
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

## Solution Approaches

### 1. Brute Force Approach (Nested Loops)

**Algorithm:**
- Use two nested loops to check every possible pair of numbers
- For each element, check if there's another element that sums to target

**Time Complexity:** O(n²)
- Outer loop: O(n)
- Inner loop: O(n) for each outer iteration
- Total: O(n × n) = O(n²)

**Space Complexity:** O(1)
- Only using a constant amount of extra space for variables

**Advantages:**
- Simple to understand and implement
- No additional data structures needed
- Works with any input (no assumptions about data)

**Disadvantages:**
- Very inefficient for large arrays
- Time complexity grows quadratically
- Not suitable for production use with large datasets

**When to Use:**
- Small arrays (n < 100)
- Educational purposes
- When memory is extremely constrained

---

### 2. Hash Map Approach (Your Current Implementation)

**Algorithm:**
- Use a hash map to store each number and its index
- For each number, calculate complement (target - current_number)
- Check if complement exists in hash map

**Time Complexity:** O(n)
- Single pass through the array: O(n)
- Hash map operations (insert/lookup): O(1) average case
- Total: O(n)

**Space Complexity:** O(n)
- Hash map stores up to n elements in worst case
- Each element takes O(1) space

**Advantages:**
- Optimal time complexity
- Single pass through array
- Very efficient for large datasets
- Easy to implement

**Disadvantages:**
- Uses additional space proportional to input size
- Hash map overhead
- May not be optimal when memory is very limited

**When to Use:**
- Most practical scenarios
- When time efficiency is priority
- Large datasets
- Production environments

---

### 3. Two Pointers Approach (Sorted Array)

**Algorithm:**
- Sort the array first
- Use two pointers (left and right)
- Move pointers based on sum comparison with target

**Time Complexity:** O(n log n)
- Sorting: O(n log n)
- Two pointers traversal: O(n)
- Total: O(n log n + n) = O(n log n)

**Space Complexity:** O(1) or O(n)
- If we can modify original array: O(1)
- If we need to preserve original array: O(n) for copy

**Advantages:**
- Constant space if array can be modified
- Good for memory-constrained environments
- Can be extended to find all pairs

**Disadvantages:**
- Requires sorting (destroys original order)
- Time complexity dominated by sorting
- Need to handle duplicate indices if original order matters

**When to Use:**
- Memory is very limited
- Array can be modified
- Need to find all pairs, not just one

---

### 4. Binary Search Approach

**Algorithm:**
- Sort the array
- For each element, use binary search to find complement

**Time Complexity:** O(n log n)
- Sorting: O(n log n)
- For each element: binary search O(log n)
- Total: O(n log n + n log n) = O(n log n)

**Space Complexity:** O(1) or O(n)
- Same considerations as two pointers approach

**Advantages:**
- Conceptually simple
- Good for educational purposes
- Can be extended to find closest sum

**Disadvantages:**
- Slower than hash map approach
- Requires sorting
- More complex than two pointers for this specific problem

**When to Use:**
- Educational purposes
- When you need to find closest match
- Memory constraints with sorted data

---

## Comparison Summary

| Approach | Time Complexity | Space Complexity | Best For |
|----------|----------------|------------------|----------|
| Brute Force | O(n²) | O(1) | Small arrays, education |
| Hash Map | O(n) | O(n) | Most practical cases |
| Two Pointers | O(n log n) | O(1) | Memory-constrained |
| Binary Search | O(n log n) | O(1) | Educational, sorted data |

## Your Current Implementation Analysis

Your hash map implementation is **optimal** for most practical scenarios:

**Strengths:**
- ✅ O(n) time complexity - best possible
- ✅ Single pass through array
- ✅ Clean, readable code
- ✅ Handles edge cases properly

**Potential Improvements:**
- Consider early termination if target is impossible
- Add input validation
- Consider memory usage for very large arrays

## Recommendations

1. **For Interviews:** Hash map approach (your current solution)
2. **For Production:** Hash map approach with proper error handling
3. **For Memory-Constrained Systems:** Two pointers approach
4. **For Education:** Implement all approaches to understand trade-offs

## Follow-up Questions

1. How would you modify the solution to return all pairs that sum to target?
2. What if the array contains duplicates?
3. How would you handle the case where no solution exists?
4. What if you need to find the pair with minimum difference from target?
