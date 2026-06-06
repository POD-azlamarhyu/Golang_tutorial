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
	fmt.Println("OS: ",runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()

	/**
	* * Mapの宣言と初期化
	* * GoのMapは、キーと値のペアを格納するデータ構造で、ハッシュテーブルを使用して実装されています。	
	* * Mapは、キーを使用して値にアクセスすることができ、キーは一意でなければなりません。
	* * 初期値はnilで、使用する前にmake関数を使用して初期化する必要があります。
	* ! make関数で作成されたMapは、実体の配列を持っていません。
	*/
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Printf("Map contents: %v\n", m)
	fmt.Printf("Value for key 'one': %d\n", m["one"])
	fmt.Printf("Value for key 'two': %d\n", m["two"])
	fmt.Printf("Map length: %d\n\n", len(m))

	/**
	* * Mapの宣言と初期化
	* * GoのMapは、キーと値のペアを格納するデータ構造で、ハッシュテーブルを使用して実装されています。	
	* * Mapは、キーを使用して値にアクセスすることができ、キーは一意でなければなりません。
	* * 初期値はnilで、使用する前にmake関数を使用して初期化する必要があります。
	* ! 組み込みのcap関数は使用できないため、Mapのサイズを指定することはできません。
	*/
	n := map[string]string{
		"first": "Golang",
		"second": "Java",
		"third": "Python",
		"fourth": "JavaScript",
		"fifth": "C++",
		"sixth": "Ruby",
		"seventh": "Swift",
		"eighth": "Kotlin",
		"ninth": "PHP",
		"tenth": "Rust",
	}
	fmt.Printf("Map contents: %v\n", n)
	fmt.Printf("Value for key 'first': %s\n", n["first"])
	fmt.Printf("Value for key 'second': %s\n", n["second"])
	fmt.Printf("Value for key 'third': %s\n", n["third"])
	fmt.Printf("Value for key 'fourth': %s\n", n["fourth"])
	fmt.Printf("Value for key 'fifth': %s\n", n["fifth"])

	v,k := n["ninth"]
	fmt.Printf("Value: %s, Key exists: %t\n\n", v, k)

	/**
	* * Mapの更新
	* * Mapの値は、キーを使用して直接更新することができます。存在しないキーに対して値を割り当てると、新しいキーと値のペアがMapに追加されます。
	*/
	n["eleventh"] = "TypeScript"
	n["twelfth"] = "Scala"
	n["thirteenth"] = "Perl"
	n["fourteenth"] = "Haskell"
	n["fifteenth"] = "Lua"
	n["sixteenth"] = "C#"
	n["seventeenth"] = "Dart"
	n["eighteenth"] = "Elixir"
	n["nineteenth"] = "Clojure"
	n["twentieth"] = "F#"
	n["twenty-first"] = "Erlang"
	fmt.Printf("Updated Map contents: %v\n", n)
}