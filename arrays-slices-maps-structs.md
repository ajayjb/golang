# Go Data Structures: Arrays, Slices, Maps, and Structs

## Table of Contents
1. [Arrays](#arrays)
2. [Slices](#slices)
3. [Maps](#maps)
4. [Structs](#structs)

---

## Arrays

### Definition
An array is a fixed-size collection of elements of the same type. Once created, the size of an array cannot be changed.

### Syntax
```go
var arrayName [size]Type
```

### Declaration and Initialization
```go
// Declaration with size
var numbers [5]int

// Declaration with initialization
var colors [3]string = [3]string{"red", "green", "blue"}

// Short declaration with type inference
numbers := [5]int{1, 2, 3, 4, 5}

// Partial initialization (remaining elements are zero values)
arr := [5]int{1, 2} // [1, 2, 0, 0, 0]

// Array size inferred from initialization
arr := [...]int{10, 20, 30} // Size is 3
```

### Accessing Elements
```go
arr := [3]string{"apple", "banana", "cherry"}

fmt.Println(arr[0])    // Output: apple
arr[1] = "blueberry"   // Modify element
fmt.Println(len(arr))  // Output: 3
```

### Iterating Over Arrays
```go
arr := [4]int{10, 20, 30, 40}

// Using for loop with index
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

// Using range
for index, value := range arr {
    fmt.Println(index, value)
}

// Using range (ignore index)
for _, value := range arr {
    fmt.Println(value)
}
```

### Key Characteristics
- **Fixed size**: Cannot grow or shrink
- **Zero-indexed**: First element is at index 0
- **Type-safe**: All elements must be the same type
- **Value type**: Arrays are passed by value to functions

---

## Slices

### Definition
A slice is a flexible, dynamic view into elements of an array. Unlike arrays, slices have variable length and can grow dynamically.

### Syntax
```go
var sliceName []Type
```

### Declaration and Initialization
```go
// Empty slice declaration
var slice []int

// Slice initialization
slice := []int{1, 2, 3, 4, 5}

// Slice from array
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4] // Elements at indices 1, 2, 3 → [20, 30, 40]

// Using make function
slice := make([]int, 5)           // Length 5, capacity 5
slice := make([]int, 3, 10)       // Length 3, capacity 10
slice := make([]int, 0, 5)        // Length 0, capacity 5
```

### Slice Operations

#### Length and Capacity
```go
slice := make([]int, 3, 10)
fmt.Println(len(slice))  // Output: 3 (number of elements)
fmt.Println(cap(slice))  // Output: 10 (capacity)
```

#### Append
```go
slice := []int{1, 2, 3}
slice = append(slice, 4)           // [1, 2, 3, 4]
slice = append(slice, 5, 6, 7)     // [1, 2, 3, 4, 5, 6, 7]

// Append another slice
slice2 := []int{8, 9}
slice = append(slice, slice2...)   // [1, 2, 3, 4, 5, 6, 7, 8, 9]
```

#### Slicing
```go
slice := []int{10, 20, 30, 40, 50}

slice[1:3]   // [20, 30]
slice[:3]    // [10, 20, 30]
slice[2:]    // [30, 40, 50]
slice[:]     // [10, 20, 30, 40, 50]
```

#### Copy
```go
source := []int{1, 2, 3, 4, 5}
dest := make([]int, 3)
copy(dest, source)  // dest = [1, 2, 3]
```

### Iterating Over Slices
```go
slice := []string{"Go", "is", "awesome"}

for i, value := range slice {
    fmt.Println(i, value)
}

for _, value := range slice {
    fmt.Println(value)
}
```

### Key Characteristics
- **Dynamic size**: Can grow and shrink
- **Reference type**: Changes in functions affect original slice
- **Backed by array**: Slices are views into arrays
- **Capacity**: Can allocate more space than current length

---

## Maps

### Definition
A map is an unordered collection of key-value pairs. Keys must be of comparable types, and all values must be of the same type.

### Syntax
```go
var mapName map[KeyType]ValueType
```

### Declaration and Initialization
```go
// Declaration (nil map, cannot add elements)
var person map[string]string

// Initialization using make
person := make(map[string]string)
person["name"] = "John"
person["age"] = "30"

// Initialization with values
person := map[string]string{
    "name": "Alice",
    "age":  "25",
    "city": "New York",
}

// Initialization with integers
scores := map[string]int{
    "Alice": 90,
    "Bob":   85,
    "Charlie": 92,
}
```

### Accessing and Modifying Elements
```go
person := map[string]string{"name": "John"}

// Access value
fmt.Println(person["name"])  // Output: John

// Add or update
person["name"] = "Jane"
person["email"] = "jane@example.com"

// Check if key exists
value, exists := person["email"]
if exists {
    fmt.Println(value)
}

// Delete key
delete(person, "email")
```

### Iterating Over Maps
```go
person := map[string]string{
    "name": "John",
    "age":  "30",
    "city": "NYC",
}

// Iterate over all key-value pairs
for key, value := range person {
    fmt.Println(key, value)
}

// Iterate over keys only
for key := range person {
    fmt.Println(key)
}

// Iterate over values only
for _, value := range person {
    fmt.Println(value)
}
```

### Common Map Operations
```go
// Get length (number of key-value pairs)
len(myMap)

// Clear all elements (Go 1.21+)
clear(myMap)

// Check if map is nil
if myMap == nil {
    // Map is not initialized
}

// Initialize nil map before use
if myMap == nil {
    myMap = make(map[string]int)
}
```

### Key Characteristics
- **Unordered**: No guaranteed order of iteration
- **Reference type**: Changes affect original map
- **Dynamic size**: Can add/remove keys at runtime
- **Safe lookup**: Returns zero value and false for missing keys
- **Keys must be comparable**: Cannot use slices or maps as keys

---

## Structs

### Definition
A struct is a composite data type that groups related variables (fields) together. It's useful for organizing data into meaningful structures.

### Syntax
```go
type StructName struct {
    Field1 Type1
    Field2 Type2
    // ...
}
```

### Declaration and Initialization
```go
// Define a struct
type Person struct {
    Name string
    Age  int
    City string
}

// Create instance with field names
p1 := Person{
    Name: "John",
    Age:  30,
    City: "New York",
}

// Create instance with positional values
p2 := Person{"Alice", 25, "London"}

// Create instance with some fields
p3 := Person{Name: "Bob", Age: 28}  // City is empty string

// Create pointer to struct
p4 := &Person{Name: "Charlie", Age: 35, City: "Paris"}
```

### Accessing Fields
```go
person := Person{Name: "John", Age: 30, City: "NYC"}

// Access fields
fmt.Println(person.Name)  // Output: John
fmt.Println(person.Age)   // Output: 30

// Modify fields
person.Age = 31

// Access through pointer
ptr := &person
fmt.Println(ptr.Name)     // Output: John (automatic dereferencing)
ptr.Age = 32
```

### Embedded Structs
```go
type Address struct {
    Street string
    City   string
    ZipCode string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // Embedding struct
}

// Access embedded fields
person := Person{
    Name: "John",
    Age:  30,
    Address: Address{
        Street:  "123 Main St",
        City:    "New York",
        ZipCode: "10001",
    },
}

fmt.Println(person.Address.City)  // Output: New York
```

### Methods on Structs
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can modify struct)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Using methods
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())  // Output: 50
rect.Scale(2)
fmt.Println(rect.Area())  // Output: 200
```

### Struct Tags
```go
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

// Tags are used by packages like json for marshaling/unmarshaling
data := []byte(`{"name":"John","age":30}`)
var person Person
json.Unmarshal(data, &person)
```

### Key Characteristics
- **Named types**: Group related fields together
- **Value type**: Structs are passed by value (copy) to functions
- **Zero values**: Uninitialized fields have zero values
- **Exported fields**: Fields starting with uppercase are exported
- **Anonymous fields**: Can use unnamed embedded structs

---

---

## Pass by Value vs Pass by Reference

### Understanding the Difference

In Go, function arguments are always passed **by value**. However, the behavior differs depending on the data type:
- **Pass by Value**: A copy of the variable is created and passed to the function
- **Pass by Reference**: You explicitly pass a pointer to the variable, allowing the function to modify the original

### Pass by Value

When you pass a variable by value, the function receives a copy. Changes made inside the function do not affect the original variable.

#### Example with Primitive Types
```go
func incrementByValue(num int) {
    num++  // This only modifies the copy
}

func main() {
    x := 5
    incrementByValue(x)
    fmt.Println(x)  // Output: 5 (unchanged)
}
```

#### Example with Arrays
```go
func modifyArray(arr [3]int) {
    arr[0] = 999  // This modifies the copy only
}

func main() {
    numbers := [3]int{1, 2, 3}
    modifyArray(numbers)
    fmt.Println(numbers)  // Output: [1 2 3] (unchanged)
}
```

#### Example with Structs
```go
type Person struct {
    Name string
    Age  int
}

func incrementAge(p Person) {
    p.Age++  // This modifies the copy only
}

func main() {
    person := Person{Name: "John", Age: 30}
    incrementAge(person)
    fmt.Println(person.Age)  // Output: 30 (unchanged)
}
```

### Pass by Reference (Using Pointers)

When you pass a pointer to a variable, you're passing the memory address. The function can dereference the pointer to access and modify the original variable.

#### Syntax
```go
// Declaring a pointer
var ptr *Type

// Getting the address of a variable
ptr := &variable

// Dereferencing a pointer
value := *ptr
```

#### Example with Primitive Types
```go
func incrementByReference(num *int) {
    *num++  // Dereference and modify the original
}

func main() {
    x := 5
    incrementByReference(&x)
    fmt.Println(x)  // Output: 6 (changed)
}
```

#### Example with Structs
```go
type Person struct {
    Name string
    Age  int
}

func incrementAge(p *Person) {
    p.Age++  // Can access fields directly on pointer
}

func main() {
    person := Person{Name: "John", Age: 30}
    incrementAge(&person)
    fmt.Println(person.Age)  // Output: 31 (changed)
}
```

#### Example with Arrays
```go
func modifyArray(arr *[3]int) {
    arr[0] = 999  // Automatic dereferencing for arrays
}

func main() {
    numbers := [3]int{1, 2, 3}
    modifyArray(&numbers)
    fmt.Println(numbers)  // Output: [999 2 3] (changed)
}
```

### Special Case: Slices, Maps, and Interfaces

Slices and maps are **reference types**, meaning they contain pointers internally. Even when passed by value, they reference the same underlying data.

#### Slices - Modifying Elements
```go
func modifySliceElements(slice []int) {
    slice[0] = 999  // This modifies the original slice
}

func main() {
    numbers := []int{1, 2, 3}
    modifySliceElements(numbers)
    fmt.Println(numbers)  // Output: [999 2 3] (changed)
}
```

#### Slices - Appending
```go
func appendToSlice(slice []int) []int {
    slice = append(slice, 4)  // Returns modified slice
    return slice
}

func main() {
    numbers := []int{1, 2, 3}
    numbers = appendToSlice(numbers)  // Must reassign
    fmt.Println(numbers)  // Output: [1 2 3 4]
}
```

#### Maps - Modifying Values
```go
func modifyMap(m map[string]int) {
    m["score"] = 100  // This modifies the original map
}

func main() {
    scores := make(map[string]int)
    scores["john"] = 50
    modifyMap(scores)
    fmt.Println(scores["score"])  // Output: 100 (changed)
}
```

### Methods: Value Receivers vs Pointer Receivers

#### Value Receiver (Makes a Copy)
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver - receives a copy
func (r Rectangle) ScaleByValue(factor float64) {
    r.Width *= factor    // Only modifies the copy
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    rect.ScaleByValue(2)
    fmt.Println(rect.Width)  // Output: 10 (unchanged)
}
```

#### Pointer Receiver (Modifies Original)
```go
// Pointer receiver - receives a pointer
func (r *Rectangle) ScaleByPointer(factor float64) {
    r.Width *= factor    // Modifies the original
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    rect.ScaleByPointer(2)
    fmt.Println(rect.Width)  // Output: 20 (changed)
    
    // Note: You can also call with explicit pointer
    (&rect).ScaleByPointer(3)
}
```

### Comprehensive Example

```go
package main

import "fmt"

type Account struct {
    Owner   string
    Balance int
}

// Value receiver - creates a copy
func (a Account) CheckBalance() int {
    return a.Balance
}

// Pointer receiver - operates on original
func (a *Account) Deposit(amount int) {
    a.Balance += amount
}

// Pointer receiver - operates on original
func (a *Account) Withdraw(amount int) bool {
    if a.Balance >= amount {
        a.Balance -= amount
        return true
    }
    return false
}

func main() {
    account := Account{Owner: "Alice", Balance: 1000}
    
    fmt.Println("Initial balance:", account.CheckBalance())  // 1000
    
    account.Deposit(500)
    fmt.Println("After deposit:", account.CheckBalance())   // 1500
    
    account.Withdraw(200)
    fmt.Println("After withdraw:", account.CheckBalance())  // 1300
}
```

### Decision Guide

| Scenario | Use Pass by Value | Use Pass by Reference |
|----------|-------------------|----------------------|
| Small data (int, bool, small string) | ✓ | - |
| Large structs (to avoid copying) | - | ✓ |
| Need to modify original variable | - | ✓ |
| Need to modify struct fields | - | ✓ |
| Want to prevent unintended changes | ✓ | - |
| Method needs to modify receiver | - | ✓ |
| Arrays (fixed-size) | - | ✓ |
| Slices (already references) | ✓ | - |
| Maps (already references) | ✓ | - |

### Common Pitfalls

#### Pitfall 1: Forgetting to Reassign After Append
```go
// WRONG - changes lost
func addToSlice(s []int) {
    s = append(s, 4)  // Local copy is extended
}
// Changes are not visible to caller

// RIGHT - return the modified slice
func addToSlice(s []int) []int {
    return append(s, 4)
}
numbers = addToSlice(numbers)  // Reassign
```

#### Pitfall 2: Modifying Slice Length vs Capacity
```go
func tryExtend(s []int) {
    s = append(s, 5)  // This might allocate a new array
}

numbers := make([]int, 2, 5)
tryExtend(numbers)
// If the function's append causes reallocation,
// the caller's slice still points to old array!
```

#### Pitfall 3: Pointer to Array vs Slice
```go
// Pointer to array - different behavior
func modifyArray(arr *[3]int) {
    arr[0] = 999
}

// Slice - different behavior
func modifySlice(s []int) {
    s[0] = 999
}
```

---

## Comparison Table

| Feature | Array | Slice | Map | Struct |
|---------|-------|-------|-----|--------|
| **Fixed Size** | Yes | No | No | N/A |
| **Dynamic** | No | Yes | Yes | N/A |
| **Type** | Value | Reference | Reference | Value |
| **Zero Value** | All zeros | nil | nil | Zero values for fields |
| **Ordered** | Yes | Yes | No | N/A |
| **Comparable** | Yes | No | No | Yes (if fields are) |
| **Key Types** | N/A | N/A | Comparable types | N/A |
| **Pass by Value** | Copies entire array | Shares underlying data | Shares underlying data | Copies struct |
| **Pass by Reference** | Use `*[Size]Type` | Rare, use pointers | Rare, use pointers | Use `*StructType` |

---

## Best Practices

### Arrays
- Use when you need a fixed-size collection
- Pass by pointer to functions if you need to modify

### Slices
- Preferred over arrays for most use cases
- Use `make()` to pre-allocate capacity when size is known
- Be careful with slice references when modifying

### Maps
- Use for key-value lookups and unordered collections
- Always check if key exists before accessing
- Initialize nil maps before use
- Don't rely on iteration order

### Structs
- Use to group related data into meaningful types
- Use pointer receivers when modifying structs in methods
- Export fields that need to be accessible outside the package
- Use struct tags for serialization (JSON, YAML, etc.)

---

 **Slices** offer flexibility and dynamic sizing
- **Maps** enable efficient key-value lookups
- **Structs** allow you to create custom, meaningful data types

Choose the right data structure based on your needs for type safety, performance, and code clarity.
