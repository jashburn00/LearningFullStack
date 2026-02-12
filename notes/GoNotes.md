# Go Notes 

## Contents
- [Go Notes](#go-notes)
  - [Contents](#contents)
  - [Project Structure](#project-structure)
  - [Basic Syntax and Data Types](#basic-syntax-and-data-types)
    - [instantializing variables](#instantializing-variables)
    - [functions](#functions)
    - [arrays and slices](#arrays-and-slices)
    - [maps](#maps)
    - [loops](#loops)
    - [conditional statements](#conditional-statements)
    - [error handling](#error-handling)
    - [print statements](#print-statements)
  - [User Defined Types](#user-defined-types)
    - [struct](#struct)
  - [Testing](#testing)
    - [file organization](#file-organization)
    - [writing tests](#writing-tests)

## Project Structure
- use the following command to initialize a Go project:
> go mod init myproject
- then you can start creating files:
> touch main.go
- each file needs to start with a package name: `package main`
- and an accompanying function: `func main() {...`
- you can run a file with the run command:
> go run main.go

## Basic Syntax and Data Types
### instantializing variables
- `a := "hello world"` or `var a string`
- assigning variables just uses `=`
- `const language = "EN"`
- block declaration can be used in some situations (const declarations, import statements):

            const (
                language = "EN"
                date = "today"
            )

### functions
- name with PascalCase for public, and camelCase for private 
- declared with `func`:

            func SayHello(name string) {

            }

- you can optionally name and type the return variable in the signature:
  - `func SayHello(name string) (result string) { ...`

### arrays and slices
- arrays 
  - fixed length and type 
  - `fibonacci := [6]int{0,1,1,2,3,5}
- slices
  - dynamic sizing
  - `var fibonacci []int = primes[1:4]`
### maps
- maps
  - key-value pairs following the convention `map[keyType]valueType`
  - `m := make(map[string]int)`
  - `m["firstkey"] = 1`
### loops
- for loop syntax: 
            for i := 0; i < 5; i++ {
                // do something
            }
- Go has no while loop, but you can use for loops to imitate it:
            for sum < 5 {
                sum += sum
            }
- you can use the "range" syntax to loop through slices conveniently:
            s := []int{1,2,3,4,5}
            for i, v := range s {
                fmt.Println(i,v)
            }
- you can create an infinite loop with `for { `
### conditional statements
- if statements: 
            if num1 > num2 {
                //do stuff
            } else {
                //do other stuff
            }
- switch statements:
            switch x {
                case "uno":
                    //do something
                case "dos":
                    //do something
                default:
                    //do something
            }
### error handling
- many (most?) library functions return a value and an error 
            f, err := os.Open("filename.txt")
            if err != nil {
                log.Fatal(err)
            }
### print statements
- usually use the `fmt` library, which is similar to C's print functionality but has more features

            fmt.Println("hello world")
            fmt.Printf("%s times %d", "hello world", 100) //hello world times 100
            s := fmt.Sprintf("I will be stored in a variable") //the 'S' prefix returns a string instead of printing
- you can use `%q` to print something in quotation marks


## User Defined Types
### struct
- has keys and values, which can be functions
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


## Testing
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