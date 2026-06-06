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


	var a []int = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30}
	fmt.Printf("Length of array: %d\n", len(a))
	fmt.Printf("Capacity of array: %d\n", cap(a))
	fmt.Printf("Array contents: %v\n", a)

	fmt.Println(a[1:10])
	fmt.Println(a[10:20])
	fmt.Println(a[:15])

	/*
	* Array
	* 2次元配列
	* 固定の長さの配列
	*/
	var b [][]int = [][]int{{1,2,3}, {4,5,6}, {7,8,9}, {10,11,12}, {13,14,15}, {16,17,18}, {19,20,21}, {22,23,24}, {25,26,27}, {28,29,30}, {31,32,33}, {34,35,36}, {37,38,39}, {40,41,42}, {43,44,45}, {46,47,48}, {49,50,51}, {52,53,54}, {55,56,57}, {58,59,60}, {61,62,63}, {64,65,66}, {67,68,69}, {70,71,72}, {73,74,75}, {76,77,78}, {79,80,81}, {82,83,84}, {85,86,87}, {88,89,90}, {91,92,93}, {94,95,96}, {97,98,99}}
	fmt.Printf("2D array contents: %v\n", b)
	fmt.Printf("Element at [1][2]: %d\n", b[1][2])

	fmt.Println(b[:10][:])

	/**
	* Slice
	* *可変の長さの配列
	* *make関数を使用することで、スライスを作成可能。
	* @param a スライスの型
	* @param length スライスの初期長さ
	* @param capacity スライスの初期容量
	*/
	d := make([]int, 1,10)

	fmt.Printf("Initial slice contents: %v\n", d)
	fmt.Printf("Length of slice: %d\n", len(d))
	fmt.Printf("Capacity of slice: %d\n", cap(d))
	fmt.Println()
	
	/*
	* append関数を使用して、スライスに要素を追加可能。
	* スライスの容量が不足した場合、自動的に新しい配列が割り当てられ、要素がコピーされる。
	*/
	e := append(d, 1,2,3,4,5,6,7,8,9)
	fmt.Printf("Slice contents: %v\n", e)
	fmt.Printf("Length of slice: %d\n", len(e))
	fmt.Printf("Capacity of slice: %d\n", cap(e))

	f := make([]int, 0, 10000)
	g := make([]int, 1,1000000)

	fmt.Println()
	fmt.Printf("Slice f contents: %v\n", f)
	fmt.Printf("Slice g contents: %v\n", g)
	fmt.Printf("Length of slice f: %d\n", len(f))
	fmt.Printf("Length of slice g: %d\n", len(g))
	fmt.Printf("Capacity of slice f: %d\n", cap(f))
	fmt.Printf("Capacity of slice g: %d\n", cap(g))
}