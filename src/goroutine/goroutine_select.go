package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

)

func hogeHoge() {
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

func networkRequest1(ch chan<- string, requestName string) {
	// Simulate a network request with a delay
	for{
		time.Sleep(1900 * time.Millisecond)
		ch <- fmt.Sprintf("packet from %s completed", requestName)
	}
}
func networkRequest2(ch chan<- string, requestName string) {
	// Simulate a network request with a delay
	for{
		time.Sleep(750 * time.Millisecond)
		ch <- fmt.Sprintf("packet from %s completed", requestName)
	}
}

/**
* ?GoroutineでゴルーチンのSelect文を使用して、複数のチャネルからのメッセージを同時に処理する例
*/
func main() {
	startTime := time.Now()
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go networkRequest1(ch1, "Request 1")
	go networkRequest2(ch2, "Request 2")

	for{
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from channel 2:", msg2)
		default:
			fmt.Println("No messages received, waiting...")
			time.Sleep(500 * time.Millisecond)
	}
		break
	}
	
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
