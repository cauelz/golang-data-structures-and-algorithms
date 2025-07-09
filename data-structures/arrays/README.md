# Arrays in Go

## Overview

An array is a fixed-size collection of elements of the same type. In Go, arrays have a fixed length that must be specified at declaration time and cannot be changed afterward.

## Key Characteristics

- **Fixed Size**: Once declared, the size cannot be changed
- **Zero-based Indexing**: Elements are accessed using indices starting from 0
- **Type Safety**: All elements must be of the same type
- **Value Semantics**: Arrays are copied when assigned or passed to functions

## Declaration and Initialization

### Basic Declaration

```go
// Declare an array of 5 integers (all initialized to zero)
var numbers [5]int

// Declare and initialize
var scores [3]int = [3]int{85, 92, 78}

// Short declaration with initialization
grades := [4]int{90, 85, 88, 92}

// Let compiler count the elements
names := [...]string{"Alice", "Bob", "Charlie"}
```

### Array Literals

```go
// Initialize specific elements
numbers := [5]int{1, 2, 3} // [1, 2, 3, 0, 0]

// Initialize with specific indices
scores := [5]int{0: 100, 2: 85, 4: 92} // [100, 0, 85, 0, 92]
```

## Accessing Elements

```go
numbers := [5]int{10, 20, 30, 40, 50}

// Access by index
first := numbers[0]  // 10
last := numbers[4]   // 50

// Modify elements
numbers[1] = 25
```

## Common Operations

### Iterating Over Arrays

```go
numbers := [5]int{1, 2, 3, 4, 5}

// Using for loop with index
for i := 0; i < len(numbers); i++ {
    fmt.Printf("Index %d: %d\n", i, numbers[i])
}

// Using range (index and value)
for index, value := range numbers {
    fmt.Printf("Index %d: %d\n", index, value)
}

// Using range (index only)
for index := range numbers {
    fmt.Printf("Index %d\n", index)
}

// Using range (value only)
for _, value := range numbers {
    fmt.Printf("Value: %d\n", value)
}
```

### Finding Length

```go
numbers := [5]int{1, 2, 3, 4, 5}
length := len(numbers) // 5
```

### Copying Arrays

```go
original := [3]int{1, 2, 3}
copy := original // Creates a copy of the entire array
```

## Multi-dimensional Arrays

### 2D Arrays

```go
// Declare a 2x3 array
var matrix [2][3]int

// Initialize 2D array
grid := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

// Access elements
value := grid[0][1] // 2
```

### 3D Arrays

```go
// Declare a 2x3x4 array
var cube [2][3][4]int

// Initialize 3D array
cube = [2][3][4]int{
    {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    },
    {
        {13, 14, 15, 16},
        {17, 18, 19, 20},
        {21, 22, 23, 24},
    },
}
```

## Array Functions

### Built-in Functions

```go
numbers := [5]int{1, 2, 3, 4, 5}

// Length
length := len(numbers)

// Capacity (same as length for arrays)
capacity := cap(numbers)
```

### Custom Array Functions

```go
// Sum all elements
func sumArray(arr [5]int) int {
    sum := 0
    for _, value := range arr {
        sum += value
    }
    return sum
}

// Find maximum value
func maxArray(arr [5]int) int {
    max := arr[0]
    for _, value := range arr {
        if value > max {
            max = value
        }
    }
    return max
}

// Reverse array
func reverseArray(arr [5]int) [5]int {
    result := arr
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return result
}
```

## Arrays vs Slices

| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed | Dynamic |
| Declaration | `var arr [5]int` | `var slice []int` |
| Length | `len(arr)` | `len(slice)` |
| Capacity | `cap(arr)` (same as length) | `cap(slice)` (can be larger) |
| Passing | Copied by value | Passed by reference |

## Common Use Cases

1. **Fixed-size collections**: When you know the exact number of elements
2. **Performance-critical code**: Arrays can be more efficient than slices
3. **Multi-dimensional data**: Matrices, grids, etc.
4. **Small, known datasets**: Configuration data, lookup tables

## Best Practices

1. **Use arrays when size is known and fixed**
2. **Use slices when size is dynamic or unknown**
3. **Be careful with array bounds** - Go will panic on out-of-bounds access
4. **Consider memory usage** - arrays are copied when passed to functions
5. **Use meaningful variable names** for array indices

## Example: Student Grades System

```go
package main

import "fmt"

type Student struct {
    Name   string
    Grades [5]int
}

func (s Student) Average() float64 {
    sum := 0
    for _, grade := range s.Grades {
        sum += grade
    }
    return float64(sum) / float64(len(s.Grades))
}

func main() {
    students := [3]Student{
        {"Alice", [5]int{85, 90, 88, 92, 87}},
        {"Bob", [5]int{78, 85, 80, 88, 82}},
        {"Charlie", [5]int{92, 95, 89, 91, 94}},
    }

    for _, student := range students {
        fmt.Printf("%s: Average = %.2f\n", student.Name, student.Average())
    }
}
```

## Performance Considerations

- **Memory**: Arrays are stored contiguously in memory
- **Access**: O(1) time complexity for random access
- **Copying**: O(n) time complexity when copying
- **Size**: Fixed size can be both an advantage and limitation

## Summary

Arrays in Go are powerful, fixed-size collections that provide:
- Type safety and compile-time bounds checking
- Efficient memory layout and access
- Simple syntax for multi-dimensional data
- Predictable performance characteristics

While slices are more commonly used in Go due to their flexibility, arrays are essential for scenarios where fixed size is required or beneficial.
