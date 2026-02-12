# Go Notes 

## Contents
- [Go Notes](#go-notes)
  - [Contents](#contents)
  - [Project Structure](#project-structure)
  - [Basic Syntax and Data Types](#basic-syntax-and-data-types)
    - [instantializing variables](#instantializing-variables)
    - [arrays and slices](#arrays-and-slices)
    - [maps](#maps)
    - [loops](#loops)
    - [conditional statements](#conditional-statements)
    - [error handling](#error-handling)
  - [User Defined Types](#user-defined-types)
    - [struct](#struct)

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
### arrays and slices
- arrays 
  - fixed length and type 
  - `fibonacci := [6]int{0,1,1,2,3,5}
- slices
  - dynamic sizing
  - `var fibonacci []int = primes[1:4]`
### maps
- maps
  - key-value pairs following the convention `map[keyType]valueType
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