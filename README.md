## [golang feature](https://go.dev/)

- Strong and statically types
    The data types are static ( declare a variable with the type you want to assign ) and cannot be changed during runtime implicitly. 
- Excellent community
    Faster response in the community.
    Nicely documented

- Key features
    - Simplicity
        The syntax of the language is easy to understand
    - Fast compile times
        One of the basic reason for faster compiling is the it maintain the acyclic dependency graph and any unused import throws an error while compiling
    - Garbage collected
        It has built-in GC, so the developer don't have to care about the memory-management
    - Built-in concurrency
    - Compile to standalone binaries
        It compile everything in golang package in a single binary including libraries and other go dependency 


Other Resources: 
[Why Go compiles so fast ](https://devrajcoder.medium.com/why-go-compiles-so-fast-772435b6bd86)

[Go: A Tale of Concurrency ( A Beginners Guide )](https://medium.com/swlh/go-a-tale-of-concurrency-a-beginners-guide-b8976b26feb)

[Go library](https://pkg.go.dev/std)



Points
1. Execution Flow
    Program execution start with main package and main function
2. Running a go program
    ``` go run <program path> ```
3. Building a executables
    ``` go build <program path> ```
    Output : executable with package name





Topics 

1. [Variables](./notes/variables.md)