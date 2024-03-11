package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//define options as flags
	sizeInBytes := flag.Bool("c", false, "Find file size in bytes")
	lineCount := flag.Bool("l", false, "Find line count")
	wordCount := flag.Bool("w", false, "Find word count")

	//parse the flags
	flag.Parse()
	// filename from non flag arguments
	fileName := flag.Arg(0)
	if fileName == "" {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}
	fmt.Println("File name: ", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()
	// if -c flag is set, get file size in bytes
	if *sizeInBytes {
		fileSize := getFileSizeInBytes(file)
		fmt.Println("File size in bytes: ", fileSize)
	}
	if *lineCount {
		lineCount := getLineCount(file)
		fmt.Println("Line count: ", lineCount)
	}
	if *wordCount {
		wordCount := getWordCount(file)
		fmt.Println("Word count: ", wordCount)
	}
}
func getWordCount(file *os.File) int64 {
	scanner := bufio.NewScanner(file)
	var wordCount int64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += int64(len(words))
	}
	file.Seek(0, 0)
	return wordCount
}
func getLineCount(file *os.File) int64 {
	scanner := bufio.NewScanner(file)
	var lineCount int64 = 0
	for scanner.Scan() {
		lineCount++
	}
	file.Seek(0, 0)
	return lineCount
}
func getFileSizeInBytes(file *os.File) int64 {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	return fileInfo.Size()
}
