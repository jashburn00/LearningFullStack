package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	title := "Quality Learing Center"
	fmt.Printf("Welcome to the %s\n", title)
	fmt.Printf("\n")

	//slice example
	mySlice := []string{"apple", "banana", "rice"}
	for i, v := range mySlice {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	fmt.Printf("\n")

	//map example
	myMap := map[string]int{
		"apple":  -5,
		"banana": 100,
		"rice":   100,
	}
	for k, v := range myMap {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Printf("\n")

	//switch example
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Running on macOS")
	case "linux":
		fmt.Println("Running on Linux")
	case "windows":
		fmt.Println("Running on Windows")
	default:
		fmt.Printf("Running on %s\n", os)
	}
	fmt.Printf("\n")

	//error handling example
	f, err := os.Open("nonexistent_file.txt")
	if err != nil {
		f.Close()
		log.Fatal(err) //this will exit the program immediately
		//log.Println(err)
	}
	//more code here...
	f.Close()
	fmt.Printf("\n")

}
