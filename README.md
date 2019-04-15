# GoLang101


Go (often referred to as Golang) is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

Here is an hands-on introduction to Go using example programs.
Run them here: https://play.golang.org/

##### hello world.go

Our first program will print the classic “hello world” message.

##### var-values-const.go

  - Go has various value types including strings, integers, floats, booleans, etc. 
  - Variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
  - The := syntax is shorthand for declaring and initializing a variable
  - Go supports constants of character, string, boolean, and numeric values.
##### for.go

for is Go’s only looping construct.

#####  if-else-switch.go

  - Branching with if and else in Go is straight-forward.
  - Note that you don’t need parentheses around conditions in Go, but that the braces are required.
  - There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
  - Switch statements express conditionals across many branches.

##### arrays.go

In Go, an array is a numbered sequence of elements of a specific length. The builtin len returns the length of an array.
Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

##### slices.go

  - You’ll see slices much more often than arrays in typical Go. Slices are a key data type, giving a more powerful interface to sequences than array.
  - Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make.
  - In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. Note that we need to accept a return value from append as we may get a new slice value.

##### maps.go

Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).

##### range.go

range iterates over elements in a variety of data structures.

##### functions.go

Functions are central in Go. Go requires explicit returns. Call a function just as you’d expect, with name(args).
Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.
