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

func alternate(x *int){
	*x = *x + 1
}


func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	/**
	* *メモリ確保とポインタについて
	* new関数と make関数の違い
	* new関数は、指定された型のゼロ値を持つ新しい変数を作成し、そのアドレスを返します。例えば、new(int)は0を持つ新しい整数変数を作成し、そのアドレスを返します。
	* make関数は、スライス、マップ、チャネルなどの組み込みデータ構造を初期化するために使用されます。例えば、make([]int, 5)は長さ5の整数スライスを作成します。
	*
	* ポインタは、変数のメモリアドレスを格納する変数です。ポインタを使用すると、関数に変数のアドレスを渡すことができ、関数内でその変数の値を変更できます。
	* 例えば、alternate関数は整数へのポインタを受け取り、その整数の値を1増やします。main関数でxを定義し、そのアドレスをalternate関数に渡すことで、xの値を変更できます
	*/

	var x *int = new(int)
	var y *int
	fmt.Printf("Initial value of x: %d\n", *x)
	fmt.Printf("Initial value of y: %v\n", y)
	fmt.Println()

	s := make([]int, 0)
	m := make(map[string]int)
	c := make(chan int)

	var p *int = new(int)
	var st = new(struct{ a int; b string })

	fmt.Printf("Initial value of s: %v (%T) address: %p\n", s, s, &s)
	fmt.Printf("Initial value of m: %v (%T) address: %p\n", m, m, &m)
	fmt.Printf("Initial value of c: %v (%T) address: %p\n", c, c, &c)
	fmt.Println()

	fmt.Printf("Initial value of p: %d (%T)\n", *p, p)
	fmt.Printf("Initial value of st: %+v (%T)\n", *st, st)
}