package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
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
	var file *os.File = nil
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
		file = os.Stdin
	} else {
		data, err = os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	}

	// if -c flag is set, get file size in bytes
	if *sizeInBytes {
		fmt.Println("File size in bytes: ", getFileSizeInBytes(file, data))
	}
	if *lineCount {

		fmt.Println("Line count: ", bytes.Count(data, []byte{'\n'}))
	}
	if *wordCount {

		fmt.Println("Word count: ", len(bytes.Fields(data)))
	}
}

func getFileSizeInBytes(file *os.File, data []byte) int64 {
	if file == os.Stdin {
		return int64(len(data))
	} else {
		file.Seek(0, 0)
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("Error: ", err)
			panic(err)
		}
		return fileInfo.Size()
	}
}
