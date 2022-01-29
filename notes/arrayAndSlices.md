[Program Link](../src/arraySlices/main.go)
### Array
- Array is basis for slices.
- Array consist of same type of element.
- While declaring a array, its size has to be defined.
- To find the length of the array we use built-in len function; len(students)
- when we assign a array to another variable, it make `hard-copy` of the initial array
- To make `soft-copy` assign pointer of initial array

Syntax to declare array:
    ```
        variableName := [<size>]<type>{<initialValues>}
        var students [3]string

        2D- Array
        var matrix [2][2]int
        matrix[0] = [2]int{0,1}
        matrix[1] = [3]int{1,0}

        // Hardcopy
        b := matrix

        //Softcopy
        b := &matrix
    ```

### Slices
- While declaring slices we don't have to mention size to the array.
- Array and slices have same behavior i.e. All the operation of array can be use for slices ( with some exception).
- Slices make softcopy of the initial array/slices (it points to reference of the initial array or slices)

    ```
        names := []string{"Rajan", "Raj", "Das"}
        namesCopy = names
        namesFirstTwo = names[:2] // `:` defines the range of index to include in the array
    ``` 

**Declaring slices using make() function**
- Make is a built-in function, that takes 2 or 3 arguments.

Syntax
    make(<type> , <length>, <capacity>)

    example : 
        make([]int , 2 , 100)

* We can define capacity of the slices, if we know number of element we will deal with in order to avoid the auto expansion of the array and avoid copy operation internally.
