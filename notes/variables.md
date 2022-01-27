[language specification](https://go.dev/ref/spec) 

### Variable Declaration
[Program Link](../src/variables/main.go)

- In go, we use "var" declare variable; followed by variable name and type of the variable. 
    example: 

        ``` 
        var programmingName string = "golang"
         ```

        or (dynamically assign a type)

        ``` 
        programmingName := 1
        ```
- Variable in go always has to be used.
- We can't redeclare a variable, but we can shadow them in different scope.

#### Type of Primitive Variable

- Boolean Type

    ``` 
    	var isCorrect bool = true; 
    ```

- Numeric Type
    ``` 
        var PORT int = 3001
        var expenses float = 10.98 //float64
    ```

- Text types

#### Type conversion
    Explicit type conversion; **no implicit conversion**

    ```
        var i int = 24
        vat j float64 
        j = float64(i)
    ```

    In order to convert a number to a string we make use of `strconv` package.
    example : 

    ```
        var j string
        j = strconv.Itoa(i)
    ```


**How to increase code readability for variable declaration?**
In order to group a set of variable so that it represent a set, we can write the variable declaration in the following format.

    ```
    var (
		firstName   string = "Rajan"
		address     string = "ABC"
		company     string = "XYZ"
		phoneNumber int    = 123
	)
    ```
