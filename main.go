package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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
	//if no flags set then do all operations
	if !*sizeInBytes && !*lineCount && !*wordCount {
		*sizeInBytes = true
		*lineCount = true
		*wordCount = true
	}
	// filename from non flag arguments
	fileName := flag.Arg(0)
	var data []byte
	var err error
	if fileName == "" {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("No file provided")
			os.Exit(1)
		}
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		// create a new txt file to write the data to
		file, err := os.Create("temp.txt")
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		// write the data to the file
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		// close the file
		file.Close()
	} else {
		data, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	}

	// if -c flag is set, get file size in bytes
	if *sizeInBytes {
		fmt.Println("File size in bytes: ", len(data))
	}
	if *lineCount {

		fmt.Println("Line count: ", bytes.Count(data, []byte{'\n'}))
	}
	if *wordCount {

		fmt.Println("Word count: ", len(bytes.Fields(data)))
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
	if file == os.Stdin {
		bytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		return int64(len(bytes))
	} else {
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("Error: ", err)
			panic(err)
		}
		return fileInfo.Size()
	}
}
