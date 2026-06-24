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

/** 
* * Generics 少し実践的な例
* ? 
*/

type Post struct {
	PostId int
	Title   string
	Content string
	Category string
	UserId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	PostId  int
	Content string
	UserId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GenericsInterface interface {
	GetUserId() int
	GetCreatedAt() time.Time
}

func (p *Post) GetUserId() int {
	return p.UserId
}

func (c *Comment) GetUserId() int {
	return c.UserId
}

func (p *Post) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (c *Comment) GetCreatedAt() time.Time {
	return c.CreatedAt
}

type SearchService[T PostAndComment] interface {
	SearchPostsByUserId(userId int) T
	SearchCommentsByUserId(userId int) T
	SearchByCreatedAt(startDate, endDate time.Time) T
}

type PostAndComment interface {
	*Post | *Comment
	GenericsInterface
}

func FilterByUserId[T PostAndComment](items []T, userId int) []T {
	var filtered []T
	for _, item := range items {
		if (item).GetUserId() == userId {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func SearchByCreatedAt[T PostAndComment](items []T, startDate, endDate time.Time) []T {
	var filtered []T
	for _, item := range items {
		createdAt := (item).GetCreatedAt()
		if createdAt.After(startDate) && createdAt.Before(endDate) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}


// Generics 例
func main() {
	startTime := time.Now()
	fmt.Println("Hello, World!")

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Println()
	
	posts := []*Post{
		{PostId: 1, Title: "Post 1", Content: "Content 1", Category: "Category 1", UserId: 1, CreatedAt: time.Now().Add(-24 * time.Hour), UpdatedAt: time.Now()},
		{PostId: 2, Title: "Post 2", Content: "Content 2", Category: "Category 2", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 3, Title: "Post 3", Content: "Content 3", Category: "Category 3", UserId: 1, CreatedAt: time.Now().Add(-10 * time.Hour), UpdatedAt: time.Now()},
		{PostId: 4, Title: "Post 4", Content: "Content 4", Category: "Category 4", 	UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 5, Title: "Post 5", Content: "Content 5", Category: "Category 5", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 6, Title: "Post 6", Content: "Content 6", Category: "Category 6", UserId: 1, CreatedAt: time.Now().Add(-48 * time.Hour), UpdatedAt: time.Now()},
		{PostId: 7, Title: "Post 7", Content: "Content 7", Category: "Category 7", UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 8, Title: "Post 8", Content: "Content 8", Category: "Category 8", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 9, Title: "Post 9", Content: "Content 9", Category: "Category 9", UserId: 1, CreatedAt: time.Now().Add(-72 * time.Hour), UpdatedAt: time.Now()},
		{PostId: 10, Title: "Post 10", Content: "Content 10", Category: "Category 10", UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 11, Title: "Post 11", Content: "Content 11", Category: "Category 11", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 12, Title: "Post 12", Content: "Content 12", Category: "Category 12", UserId: 1, CreatedAt: time.Now().Add(-96 * time.Hour), UpdatedAt: time.Now()},
		{PostId: 13, Title: "Post 13", Content: "Content 13", Category: "Category 13", UserId: 3, CreatedAt: time.Now().Add(-12 * time.Hour), UpdatedAt: time.Now()},
	}

	comments := []*Comment{
		{PostId: 1, Content: "Comment 1", UserId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 2, Content: "Comment 2", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 3, Content: "Comment 3", UserId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 4, Content: "Comment 4", UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 5, Content: "Comment 5", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 6, Content: "Comment 6", UserId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 7, Content: "Comment 7", UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 8, Content: "Comment 8", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 9, Content: "Comment 9", UserId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 10, Content: "Comment 10", UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 11, Content: "Comment 11", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 12, Content: "Comment 12", UserId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 13, Content: "Comment 13", UserId: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 14, Content: "Comment 14", UserId: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{PostId: 15, Content: "Comment 15", UserId: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	userId := 1
	filteredPosts := FilterByUserId(posts, userId)
	filteredComments := FilterByUserId(comments, userId)

	fmt.Printf("Filtered Posts for User ID %d:\n", userId)
	for _, post := range filteredPosts {
		fmt.Printf("Post ID: %d, Title: %s, Content: %s, Category: %s, User ID: %d\n", post.PostId, post.Title, post.Content, post.Category, post.UserId)
	}

	fmt.Println()

	fmt.Printf("Filtered Comments for User ID %d:\n", userId)
	for _, comment := range filteredComments {
		fmt.Printf("Post ID: %d, Content: %s, User ID: %d\n", comment.PostId, comment.Content, comment.UserId)
	}

	filteredPostsByDate := SearchByCreatedAt(posts, time.Now().Add(-24*time.Hour), time.Now())
	filteredCommentsByDate := SearchByCreatedAt(comments, time.Now().Add(-24*time.Hour), time.Now())
	
	fmt.Println()
	fmt.Printf("Filtered Posts created in the last 24 hours:\n")
	for _, post := range filteredPostsByDate {
		fmt.Printf("Post ID: %d, Title: %s, Content: %s, Category: %s, User ID: %d\n", post.PostId, post.Title, post.Content, post.Category, post.UserId)
	}
	
	fmt.Println()
	fmt.Printf("Filtered Comments created in the last 24 hours:\n")
	for _, comment := range filteredCommentsByDate {
		fmt.Printf("Post ID: %d, Content: %s, User ID: %d\n", comment.PostId, comment.Content, comment.UserId)
	}

	fmt.Println("####################################################################################")
	elapsedTime := time.Since(startTime)
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
