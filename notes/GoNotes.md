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
  - [loops, flow, and conditional statements](#loops-flow-and-conditional-statements)
    - [for loop syntax:](#for-loop-syntax)
    - [range iteration](#range-iteration)
    - [if statements:](#if-statements)
    - [switch statements:](#switch-statements)
    - [defer](#defer)
  - [Readers and I/O](#readers-and-io)
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
- the `io` package includes an `io.Reader` interface which represents the read end of a stream of data
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

- apparently a "Reader type" can mean output not consuming data 
- "emit" means to fill a given buffer with bytes and return the number of bytes filled
- "Reader" sometimes means "Rewriter" or "Writer" (gay)

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