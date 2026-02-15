# Go Notes 

## Contents
- [Go Notes](#go-notes)
  - [Contents](#contents)
  - [Project Structure](#project-structure)
  - [Basic Syntax and Data Types](#basic-syntax-and-data-types)
    - [instantializing variables](#instantializing-variables)
    - [basic data types](#basic-data-types)
    - [functions](#functions)
    - [print statements](#print-statements)
  - [Advanced data types](#advanced-data-types)
    - [arrays and slices](#arrays-and-slices)
    - [maps](#maps)
    - [pointers](#pointers)
    - [User Defined Types](#user-defined-types)
    - [type parameters](#type-parameters)
    - [generic types](#generic-types)
  - [loops, flow, and conditional statements](#loops-flow-and-conditional-statements)
    - [for loop syntax:](#for-loop-syntax)
    - [range iteration](#range-iteration)
    - [if statements:](#if-statements)
    - [switch statements:](#switch-statements)
    - [defer](#defer)
  - [Readers and I/O](#readers-and-io)
  - [Concurrency](#concurrency)
    - [goroutines](#goroutines)
    - [channels](#channels)
    - [Locking and Mutex](#locking-and-mutex)
  - [Testing and error handling](#testing-and-error-handling)
    - [file organization](#file-organization)
    - [writing tests](#writing-tests)

## Project Structure
- use the following command to initialize a Go project:
> go mod init myproject
  - then you can start creating files
- each file needs to start with a package name: `package main`
- you can run a file with the run command:
> go run main.go
- good practice(?) to keep all packages in their own folder inside the module 

## Basic Syntax and Data Types
### instantializing variables
- `a := "hello world"` or `var a string`
- `var i, j int = 1, 2`
- assigning variables just uses `=`
- const assignment cannot use `:=`
- `const language = "EN"`
- block declaration can be used in some situations (const declarations, import statements):

            const (
                language = "EN"
                date = "today"
            )

### basic data types
- basic data types:
  - bool
  - string
  - int, int8, int16, int32, int64, uint, uint8, ... uint64, uintptr
  - byte (alias for uint8)
  - rune (alias for int32)
  - float32, float64
  - complex64, complex128
- types can be converted using desiredtype(existingvalue)
  - `f := float64(myint)`

### functions
- name with PascalCase for public, and camelCase for private 
- declared with `func`:

            func SayHello(name string) {

            }

- you can optionally name and type the return variable in the signature:
  - `func SayHello(name string) (result string) { ...`
- high order functions need to declare the types of each parameter in parameter functions as well as the return type:

            func HighOrd (toBeExecuted(int, int) int) string {
                if toBeExecuted(1,3) > 5 {
                    return "yay"
                } else {
                    return "nay"
                }
            }

- __closures__
  - function values that are bound to variables from outside the function body (because they are referenced inside the function body)

            func fibonacci() func() int {
                x := 0
                y := 1
                return func () int {
                    result := x+y
                    x, y = y, x+y
                    return result
                }
            }

- __methods__
  - functions belonging to a data type 
  - need a special "receiver" argument before the signature
    - `func (p Person) jump() {` where (p Person) is the receiver
    - you can use pointer receivers as well: `(p *Person)` - helpful for mutating the receiver (non-static context)
  - you can only create methods for data types that are defined *inside the current module*
  - while functions with pointer parameters must receive a pointer, *methods* with pointer *receivers* will simply insert pointers for non-pointer receivers passed in.
    - e.g. `func (p *Person) jump() {` will treat `bob.jump()` as if it were `(*bob).jump()`
    - and vice-versa: functions with value receivers will accept pointer receivers and just treat them as value receivers
  - typically, pointer receivers are preferred to avoid copying large structs and allow mutation
  - best practice is for each data type to either have all pointer receivers or all value receivers (no race mixing)

### print statements
- usually use the `fmt` library, which is similar to C's print functionality but has more features

            fmt.Println("hello world")
            fmt.Printf("%s times %d", "hello world", 100) //hello world times 100
            s := fmt.Sprintf("I will be stored in a variable") //the 'S' prefix returns a string instead of printing

- you can use `%q` to print something in quotation marks

## Advanced data types 
### arrays and slices
- arrays 
  - fixed length and type 
  - `fibonacci := [6]int{0,1,1,2,3,5}`
- slices
  - slices do not store data, but are references to sections of arrays
  - changing the elements in a slice mutates the referenced array (as well as other slices that share the same array) 
  - declared slice literals create an array and then reference it 

            var nums = [4]int{1,2,3,4}
            var mySlice = nums[0,3] //contains 1,2

            var a [10]int
            // the following are all equivalent:
            s1 := a[0:10]
            s2 := a[:10]
            s3 := a[0:]
            s4 := a[:]

  - slices have a length and a capacity - the capacity of the slice is the size of the referenced array
    - `len(s1)` and `cap(s1)`
  - zero (default) value of a slice is `nil` and has no referenced array
  - you can use `make` for dynamically-sized arrays
    - `make` allocates a zeroed out array and returns a slice that refers to the array

            a := make([]int, 5) //len(a) = 5
            b := make([]int, 5, 10) //optionally specify capacity

  - you can append elements to slices like so:
    - `s = append(s, 1, 2, 3)` where the appended elements are the same type as the rest of the slice
    - `append` can be used to exceed the capacity of a slice, and will automatically allocate and populate a new array to accommodate the operation 

### maps
- maps
  - key-value pairs following the convention `map[keyType]valueType`
  - setting values: `m["key"] = 1`
  - accessing values: `v := m["key"]`
  - delete a value: `delete(m, key)`
  - check if a key exists: 

            elt, ok := m[key]
    
    - if the key is not found, elt will = the zero value of the map's element type 

  - default (zero) map value is `nil`, and a `nil` map cannot be given new keys
  - use the `make` function to initialize a map of a given type pair
    - `m := make(map[string]int)`
  - or you can declare a literal:

            var m := map[string]int {
                "kilogram": 2,
                "car": 3000,
                "yo mama": 93485092,
            }

  - you can omit type names when declaring literals inside of a map literal

            var m := map[string]Vertex {
                "origin": {2,9}
            }

### pointers
- syntax: `var p *int`
- default value is `nil`
- the `&` operator generates a pointer to its operand

            i := 42
            p := &i

- the `*` operator denotes the pointer's underlying value (dereferencing) 

### User Defined Types
- **structs** are collections of fields
  - have keys and values, which can be functions

            type Vertex struct {
                X, Y int    
            }
            
            type Person struct {
                Name    String  `json:"name"`
                Age     int     `json:"age"`
                Alive   bool    `json:"alive"`
            }
            func (p Person) Jump(){
                fmt.Println("I jumped")
            }
            func (p Person) BuyBeer() (bool, err){
                // use p.Age
            }

- pointers to structs are allowed to skip explicit dereferencing
  - `p.name` as opposed to `(*p).name` 
- you can declare struct literals a few different ways:
  
            var v1 = Vertex{1,2}
            var v2 = Vertex{Y: 1} //X: 0 is implicit
            var v3 = Vertex{} // X: 0 and Y: 0
            var p = &Vertex{0,1} //has type *Vertex

- **interface** types are defined as a set of signatures

            type Animal interface {
                MakeNoise()
            }

- interface types can hold any literal value that implements those methods
- there is never an explicit statement of a literal intending to implement an interface

            type Dog struct {
                Name string
            }

            func (d *Dog) MakeNoise(){
                fmt.Println("*gets ripped in half*")
            }

            var Bubby Animal = Dog("Bubby")

- now Bubby implements Animal implicitly
- under the hood, interfaces are a tuple of (value, type)
  - when you call the interface value e.g. MakeNoise() on Bubby, the method attached to the underlying type (in this case Animal)
- sometimes an interface instance may have a `nil` value, and method calls will be sent with a `nil` receiver
  - it is good practice to gracefully handle `nil` receivers. Interface tuples with a `nil` value are not themselves `nil`
- type assertions can be done by declaring a variable of a certain type and assuming that the interface instance has that underlying concrete type

            val, ok := Bubby.(Dog) //this would succeed
            val, ok := Bubby.(int) //this would cause a panic

### type parameters
- functions can be written to work on multiple types using type parameters 
- type parameters are written in square brackets before arguments

        func Index [T comparable](s []T, x T) int {
            //this function's arguments can now be any type that
            // supports comparison (this ensures you can use != 
            // and ==)
        } 

### generic types 
- using `[T any]`, you create a type parameter. Sort of like a placeholder for a type.
- used for when you want a data structure to work for multiple types
- When creating an instance of a type-parameterized object, you must pass the type in with the same square brackets 

        type List[T any] struct {
            Next *List[T]
            value T
        }

        mh := List[string]{value: "more heroin"} 
        h := List[string]{&mh, "heroin"}

- instances of `List[string]` and `List[int]` are different concrete types
- more examples:
  - `type Stack[T any] struct { ...` allows for stacks of different types of data to be created 
  - `func Map[T any] (s []T) []T { ...` this function will infer the type based on the input, but otherwise it simply accepts a slice of [T] and returns a slice of [T] (e.g. doubling all the numbers in a slice or something)
    - here is an example to make that function accept only number types:
  
            // this interface "Number" describes both ints and floats
            type Number interface{
                ~int | ~float64
            }

            // this function requires a generic type described by "Number"
            func Map[T Number] (s []T) []T {
                result := make([]T, len(s))

                for i, v := range s {
                    result[i] = v*2
                }

                return result
            }

  - `func (l *List[T]) Add(v T) { ...` is a method of the type *List[T] that takes in a value of [T] 
    - e.g. it adds the value of `v` to the receiver list 

## loops, flow, and conditional statements
### for loop syntax: 

            for i := 0; i < 5; i++ {
                // do something
            }

- Go has no while loop, but you can use for loops to imitate it:

            for sum < 5 {
                sum += sum
            }

- you can create an infinite loop with `for { `

### range iteration
- you can use the "range" syntax to loop through slices conveniently:

            s := []int{1,2,3,4,5}

            for ind, elt := range s {
                fmt.Println(ind,elt)
            }

- you can use `_` to skip (aka announce that you don't care about) the index or value
- you can optionally use only one variable, and only have the index available in the loops

### if statements: 
  - if statements can use short statements before condition 
    - variables dclared in this way are only scoped inside the if block and attached else blocks

            if num1 > num2 {
                //do stuff
            } else {
                //do other stuff
            }
            if x := 5; x > 3 {
                //do stuff
            }

### switch statements:
- switch statements in Go short-circuit when a case succeeds
- switch statements do not need an argument

            switch x {
                case "uno":
                    //do something
                case "dos":
                    //do something
                default:
                    //do something
            }

- **type switch** statements allow type assertions in quick succession

            switch v := i.(type) {
                case int:
                    //do something
                case string:
                    //do something
                case Dog:
                    //rip it in half 
            }

### defer
- `defer` defers the execution of the following statement until the surrounding function returns (aka scope is exited?)
  - arguments are evaluated immediately but function calls are delayed (so it uses a closure? sometimes?) 
- defer calls are pushed onto a stack (LIFO)
- `panic` ???

## Readers and I/O
- the `io` package includes an `io.Reader` interface which represents the "read end of a stream of data"
  - ex:

        func main() {
            r := strings.NewReader("Hello, Reader!")

            b := make([]byte, 8) //create slice of bytes
            for {
                n, err := r.Read(b) //iteratively read a byte
                fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
                fmt.Printf("b[:n] = %q\n", b[:n])
                if err == io.EOF {
                    break
                }
            }
        }

- "emit" means to fill a given buffer with bytes and return the number of bytes filled
- "Reader" sometimes acts more like a "Rewriter" or "Writer"  

## Concurrency
- performance tricks fuck yeah
- in the below example, trees are "walked" concurrently 

        package main

        import "golang.org/x/tour/tree"
        import "fmt"
        /*
        type Tree struct {
            Left  *Tree
            Value int
            Right *Tree
        }
        */

        // Walk walks the tree t sending all values
        // from the tree to the channel ch.
        func Walk(t *tree.Tree, ch chan int) {
            if (t.Left != nil) {
                Walk(t.Left, ch)
            }
            ch <- t.Value
            if(t.Right != nil){
                Walk(t.Right, ch)	
            }
        }

        // Same determines whether the trees
        // t1 and t2 contain the same values.
        func Same(t1, t2 *tree.Tree) bool {
            ch1 := make(chan int)
            ch2 := make(chan int)
            
            go func() {
                Walk(t1, ch1)
                close(ch1)
            }()
            go func() {
                Walk(t2, ch2)
                close(ch2)
            }()
            
            
            var result1, result2 string 
            
            for v1 := range ch1 {
                result1 = result1 + fmt.Sprintf("%d",v1)
            }
            
            for v2 := range ch2 {
                result2 = result2 + fmt.Sprintf("%d",v2)
            }
            
            
            
            return result1 == result2
        }

        func main() {
            if Same(tree.New(1), tree.New(1)) {
                fmt.Println("Success!")
            } else {
                fmt.Println("Failure.")
            }
            
            if Same(tree.New(1), tree.New(3)) {
                fmt.Println("Failure.")
            } else {
                fmt.Println("Success!")
            }	
            
        }

### goroutines
- goroutines are lightweight threads managed by the Go runtime
- start a goroutine with `go`
  - `go foo(x)`
  - the evaluation of `foo(x)` happens in the new goroutine 
- goroutines run in the same address space, so access to shared memory must be synchronized. 
  - the `sync` package provides primitives, but you won't always need those

### channels 
- are FIFO 
- must be declared (with type) before use: `ch := make(chan int)`
- functions can restrict the direction of channel parameters
  - `func Consume(ch <-chan int) { ...` this parameter channel can only be used to have variables receive values 
- by default, send and receive operations are blocking until both sides are ready 
  - send: `ch <- 42`
  - receive: `num := <- ch`
- buffered channels can help prevent this
  - `ch := make(chan int, 3)` this channel can accept 3 sends before sends become blocking
  - receiving from the channel is only blocking if the channel is empty
- should be closed by senders, not receivers (best practice), *but* channels do not always need to be closed
  - closing a channel is only necessary when the receivers need to be notified that there will be no more values coming

            func sum( s []int, c chan int){
                sum := 0
                for _, v := range s {
                    sum += v
                }
                c <- sum //send sum to channel instead of returning
            }
            
            go sum(sliceA, channelToUse)
            go sum(sliceB, channelToUse)
            x, y := <-c, <-c

            fmt.Println(x, y, x+y)
            
- __deadlocks__ occur when there are leftover channels with no further receiving operations, and they are a fatal error
- channels can be closed with `close(ch)`
  - you can check for closure when receiving: `v, ok := <- ch`
    - `ok == false` indicates channel closed 
  - the `range` feature can automatically detect closure

        func Producer(ch chan int){
            for i := 1; i <= 5; i++ {
                ch <- i
            }
            close(ch)
        }

        func main(){
            ch := make(chan int)

            go Producer(ch)

            //this for loop will stop once ch is closed
            for v := range ch {  
                fmt.Println(v) 
            } 
        }

- **select** lets a goroutine wait on several channel operations and proceed with the first one that becomes ready
  - "wait until one of these channel send or receive ops can proceed, then proceed with it"
- `select` blocks until one of the cases is ready
  - if several are ready, one is chosen at random

            select {
                case v := <- ch1:
                    //ready to receive something from ch1
                case ch2 <-x:
                    //ready to send something into ch2
                default:
                    //runs immediately if no channels available 
                    //(optional)
                    // can be useful for situations like:
                    // "not ready yet"
            }

- use the `time.After(Time)` function to implement timeouts 
  - `time.After` returns a channel that sends once after the specified duration

        select {
            case x <- ch:
                //value received
            case <- time.After(time.Second * 2):
                //fmt.Println("timed out")
        }

- you can also use `select` to implement an infinite event loop pattern:

        for {
            select {
                case msg <- ch1:
                    //received a message 
                case cmd <- ch2:
                    //received a command 
            }
        }

### Locking and Mutex
- TODO

## Testing and error handling
- many (most?) library functions return a value and an error 

            f, err := os.Open("filename.txt")
            if err != nil {
                log.Fatal(err)
            }

- standard library "testing"

### file organization
- for testing a file **main.go** the corresponding test file would be named **main_test.go** (they are both in the same package).
- once complete, tests can be run using the command:
> go test
- verbose is preferred 
> go text -v

### writing tests
- inside the test file, you create test functions with the same name as they are in the testee file, prepended with 'Test' 
- test functions should have the parameter `t *testing.T`
- test functions are able to call the testee function (inherently?)

            func TestHello(t *testing.T) {
                t.Run("saying hello world", func (t *testing.T) {
                    got:= Hello()
                    want:= "Hello World"

                    if got != want {
                        t.Errorf("got %q want %q", got, want)
                        //or use helper created below:
                        //performAssertion(t, got, want)
                    }
                })
            }

- you can create helper functions for use in testing:

            func performAssertion(t testing.TB, got, want string) {
                t.Helper() //to let compiler know
                if got != want {
                    t.Errorf("got %q want %q", got, want)
                }
            }

- by calling `t.Helper()` in the helper function, the compiler knows that the function is a helper
  - line numbers reported in failed test cases will be the line number calling the helper function instead of the helper function code block
- **Example tests** are used for documentation purposes
  - must start with `Example` just like test functions start with `Test`
  - get automatically added to your package's documentation
  - get executed after Test functions __IF__ they contain a `// Output: [value]` line

            func ExampleAdd() {
                sum := Add(1,5)
                fmt.Println(sum)
                // Output: 6
            }