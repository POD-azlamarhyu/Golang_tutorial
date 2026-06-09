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

/**
* !このコードは実行されずに処理が終了する。
*/
func goroutineExample() {
	fmt.Println("Goroutine example started")
	for i := 0; i < 5; i++ {
		fmt.Printf("goroutine : %d\n", i)
		// time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Goroutine example finished")
	fmt.Println()
}

/**
* * このコードは実行される。
*/
func blockingExample() {
	fmt.Println("Blocking example started")
	for i := 0; i < 5; i++ {
		fmt.Printf("Blocking : %d\n", i)
		// time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Blocking example finished")
	fmt.Println()
}

func main() {
	startTime := time.Now()
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	/**
	* ! ゴルーチンは処理が終わった場合、main関数が終了してしまうため、goroutineExample()の処理が完了する前にmain関数が終了してしまう。
	*/
	go goroutineExample()
	blockingExample()

	fmt.Printf("main finished : %s\n", time.Since(startTime))
}