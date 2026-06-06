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


func usingOs(os string) {
	switch os {
	case "Windows10":
		fmt.Println("You are using Windows 10.")
	case "Linux":
		fmt.Println("You are using Linux.")
	case "macOS":
		fmt.Println("You are using macOS.")
	case "FreeBSD":
		fmt.Println("You are using FreeBSD.")
	case "OpenBSD":
		fmt.Println("You are using OpenBSD.")
	case "NetBSD":
		fmt.Println("You are using NetBSD.")
	case "Windows11":
		fmt.Println("You are using Windows 11.")
	case "Ubuntu":
		fmt.Println("You are using Ubuntu.")
	case "Debian/GNU Linux":
		fmt.Println("You are using Debian or GNU Linux.")
	case "Fedora":
		fmt.Println("You are using Fedora.")
	case "Arch":
		fmt.Println("You are using Arch Linux.")
	case "CentOS":
		fmt.Println("You are using CentOS.")
	case "RedHat":
		fmt.Println("You are using Red Hat Enterprise Linux.")
	case "SUSE":
		fmt.Println("You are using SUSE Linux.")
	default:
		fmt.Println("Unknown operating system.")
	}
}

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	os := "SUSE"
	defer usingOs(os)
	defer fmt.Println("遅延実行を使用しています。")

	fmt.Printf("Operating system: %s\n", os)
	
}