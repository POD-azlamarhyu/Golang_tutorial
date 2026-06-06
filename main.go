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
	"log"
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

	log.Println("This is a log message.")
}

type Region struct {
	Nation string
	Area string
	Address string
	Name string
	Population int
}

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	region := Region{
		Nation: "日本",
		Area: "関東",
		Address: "東京都 千代田区",
		Name: "東京",
		Population: 13960000,
	}

	fmt.Printf("Region: %+v\n", region)
	fmt.Printf("Nation: %s\n", region.Nation)
	fmt.Printf("Area: %s\n", region.Area)
	fmt.Printf("Address: %s\n", region.Address)
	fmt.Printf("Name: %s\n", region.Name)
	fmt.Printf("Population: %d\n", region.Population)

	p := &Region{
		Nation: "日本",
		Area: "北海道",
		Address: "北海道 札幌市",
		Name: "大通り",
		Population: 1952000,
	}

	fmt.Printf("Region: %+v\n", p)
	fmt.Printf("Nation: %s\n", p.Nation)
	fmt.Printf("Area: %s\n", p.Area)
	fmt.Printf("Address: %s\n", p.Address)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Population: %d\n", p.Population)

	p.Address = "北海道 札幌市 中央区"
	fmt.Printf("Updated Address: %s\n", p.Address)
}