package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//define options as flags
	sizeInBytes := flag.Bool("c", false, "Find file size in bytes")

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
}

func getFileSizeInBytes(file *os.File) int64 {

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	return fileInfo.Size()
}
