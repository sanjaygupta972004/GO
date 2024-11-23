package main

import (
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Working with file in golang")

	// how to create directory
	err := os.Mkdir("Code", 0755)
	if err != nil {
		log.Println("Error while creating directory", err)
	}
	// remove directory is dir. is empty
	err = os.Remove("Code")
	if err != nil {
		log.Println("Error while removing directory", err)
	}

	// how to read directory
	files, err := os.ReadDir("data")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			log.Println("=>", file.Name())
			continue
		}
	}

	// how to create file in golang
	var exFile *os.File
	defer exFile.Close()
	fileInfo, err := os.Stat("sunny.txt")
	log.Println("Description of file", fileInfo)
	if err == nil {
		exFile, err = os.OpenFile("sunny.txt", os.O_RDWR, 0644)
		if err != nil {
			log.Println("Error while opening file")
			return
		}
		log.Println("File already exists")
		// reading entire file if file already exit
		content, err := io.ReadAll(exFile)
		if err != nil {
			log.Println("Error while reading file")
			return
		}
		log.Println("File content:", string(content))
	} else {
		exFile, err = os.Create("sunny.txt")
		if err != nil {
			log.Println("Error while creating file")
			return
		}
	}

	// how to write file
	content := "this is a example of file i would like to write which file has been created"

	bytesWritten, err := io.WriteString(exFile, content+"\n")
	if err != nil {
		log.Println("Error while writing content on file")
	}

	log.Println("number of bytes of file ", bytesWritten)

	// how to read file as small chunk
	file, err := os.Open("sunny.txt")
	if err != nil {
		log.Println("Error while opening file")
	}

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error while reading in chunk file", err)
		}

		// process for reading file
		log.Println("content of file", string(buffer[:n]))

	}

	// how to remove file from a directory
	err = os.Remove("sunny.txt")
	if err != nil {
		log.Println("Error while removing file from main directory", err)
	}

}
