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
	"sync"
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

/**
* * Goroutineで値を渡すチャネルのバージョン
*/
func publisher(a int, c chan int) {
	c <- a
}

/**
* * Goroutineで値を受け取るチャネルのバージョン
*/
func receiver(c chan int, wg *sync.WaitGroup) {
	for i := 0; i < 30; i++ {
		fmt.Printf("Received: %d\n", <-c)
		wg.Done()
	}
}
/**
* ?Goroutineでゴルーチンの関数からゴルーチンの関数へ値を渡すチャネルのバージョン
*/
func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	ch := make(chan int, 2)
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go publisher(a[i], ch)
	}
	go receiver(ch,&wg)
	wg.Wait()
	close(ch)
	
	fmt.Printf("main finished : %s\n", time.Since(startTime))
}
