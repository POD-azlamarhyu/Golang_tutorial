package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"bufio"
	"strings"
	"strconv"
)

func hogeHoge(){
		// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Printf("Current working directory: %s\n", cwd)

	// List files in current directory
	files, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	fmt.Println("Files in current directory:")
	for _, file := range files {
		fmt.Println(" -", file.Name())
	}

	// Create a new file and write to it
	newFilePath := filepath.Join(cwd, "example.txt")
	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString("This is an example file created by Go.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("File '%s' created and written to successfully.\n", newFilePath)

	// Read the contents of the file
	readFile, err := os.Open(newFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	fmt.Printf("Contents of '%s':\n", newFilePath)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Example of string manipulation
	sampleString := "Go is a great programming language!"
	words := strings.Fields(sampleString)
	fmt.Printf("Words in sample string: %v\n", words)

	// Example of converting string to integer
	numStr := "42"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
	fmt.Printf("Converted number: %d\n", num)
}

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println(runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()


	var a int = 111
	var b float64 = 5.3114
	var c float32 = 8.422
	var d string ="1.31114"

	fmt.Printf("a: %d, b: %f, c: %f\n", a, b, c)
	fmt.Printf("a int >> float64: %f (%T)\n", float64(a), float64(a))
	fmt.Printf("c float32 >> float64: %f (%T)\n", float64(c), float64(c))
	fmt.Printf("b float64 >> float32: %f (%T) \n", float32(b), float32(b))

	fmt.Printf("b float64 >>> int: %d (%T)\n", int(b), int(b))

	dd, _ := strconv.ParseFloat(d, 64)
	fmt.Printf("d string >> float64: %f (%T)\n", dd, dd)

	ddInt, _ := strconv.Atoi(d)
	fmt.Printf("d string >> int: %d (%T)\n", ddInt, ddInt)
}