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

type City struct {
	Name       string
	Population int
}

type Country struct {
	Name   string
	Cities []City
}

func CountryFactory(cities []City) Country {
	return Country{
		Name:   "日本",
		Cities: cities,
	}
}


func (c City) toString() string {
	return fmt.Sprintf("City: %s, Population: %d", c.Name, c.Population)
}

/**
* *Go　Structのエンベディング
*/
func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	cities := []City{
		{Name: "東京", Population: 13929286},
		{Name: "大阪", Population: 2691186},
		{Name: "名古屋", Population: 2295638},
		{Name: "札幌", Population: 1952356},
		{Name: "福岡", Population: 1531916},
	}

	country := CountryFactory(cities)
	fmt.Printf("Country: %s\n", country.Name)
	for _, city := range country.Cities {
		fmt.Printf("City: %s, Population: %d\n", city.Name, city.Population)
	}
	fmt.Println()
	city := country.Cities[0].toString()
	fmt.Println(city)
}