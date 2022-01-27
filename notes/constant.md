### Constants
[Program Link](../src/constants/main.go)

- We make use of `const` keyword to declare constant.
```
    const PORT int = 3300
```
- Constants are immutablem, but it can be shadowed
- Replaced by compiler at the compile time

**Naming convention**
Unlike convention in other language wherein we use uppercase or capitalize word to declare constant. We cannot use those convention in go because it can be exported to other package(in case of we don't want to expose it other package). We name it as other variable.

**Typed constants**
When we specify the type of the constant while it means we are declaring typed constant.
    ```
        const pi float = 3.1416
        //error if we perform operation with different type
        var i int = 10
        var sum = i + pi // error
    ```

**Untyped constants**
 - Declaring a constant without specifying type.
 Benefit: It replaces the identifier while compiling; so it there are any operation with different it will support it if it is valid.

    ```
        const pi = 10
        var i int64 = 10
        sum := pi + i
        fmt.Printf(" %v , %T\n", sum, sum) // 20 , int64
    ```

**Enumerated constants**
    We have a special symble iota which allows to assign value `0` to a constant
    ```
        const (
            a = iota //0
            b        //1
            c        //2
        )
    ```

- Enumeration expression
    We can use iota with bitwise, arithmetic and bit-shifting operation as a expression

    ```
        const (
            a = iota + 10  //10
            b              //11
            c              //12
        )
    ```