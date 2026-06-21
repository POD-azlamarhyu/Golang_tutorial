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


func main() {
	startTime := time.Now()
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()
	t := time.Now()
	fmt.Printf("Now time: %s\n", t.Format(time.RFC1123Z)) //RFC1123ZはRFC1123のタイムゾーンを数字で表す形式
	fmt.Printf("RFC1123: %s\n", t.Format(time.RFC1123)) //RFC1123はRFC1123のタイムゾーンを文字列で表す形式
	fmt.Printf("RFC3339: %s\n", t.Format(time.RFC3339)) //RFC3339はRFC3339のタイムゾーンを数字で表す形式
	fmt.Printf("RFC3339Nano: %s\n", t.Format(time.RFC3339Nano)) //RFC3339NanoはRFC3339のタイムゾーンを数字で表す形式で、ナノ秒まで表示する形式
	fmt.Printf("Kitchen: %s\n", t.Format(time.Kitchen))//Kitchenは12時間制の時間を表示する形式
	fmt.Printf("Stamp: %s\n", t.Format(time.Stamp)) //Stampは月、日、時間を表示する形式
	
	fmt.Printf("Year: %d, Month: %d, Day: %d, Hour: %d, Minute: %d, Second: %d, Nanosecond: %d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()) //Local()はローカルタイムを返すメソッド、UTC()はUTCタイムを返すメソッド
	fmt.Printf("Local Year: %d, Local Month: %d, Local Day: %d, Local Hour: %d, Local Minute: %d, Local Second: %d, Local Nanosecond: %d\n", t.Local().Year(), t.Local().Month(), t.Local().Day(), t.Local().Hour(), t.Local().Minute(), t.Local().Second(), t.Local().Nanosecond())
	fmt.Printf("UTC Year: %d, UTC Month: %d, UTC Day: %d, UTC Hour: %d, UTC Minute: %d, UTC Second: %d, UTC Nanosecond: %d\n", t.UTC().Year(), t.UTC().Month(), t.UTC().Day(), t.UTC().Hour(), t.UTC().Minute(), t.UTC().Second(), t.UTC().Nanosecond())
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
