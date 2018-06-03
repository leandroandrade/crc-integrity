package main

import (
	"os"
	"log"
	"bufio"
	"hash/crc32"
	"fmt"
	"bytes"
)

func main() {
	fileName := "/tmp/file.txt"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	file.Sync()

	// write file
	writter := bufio.NewWriter(file)
	_, err = writter.Write([]byte("This is text on file\n"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	writter.Flush()
	file.Close()

	// read file to checksum
	checksum32File(fileName)

	// open file with append option
	fileOpen, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = fileOpen.WriteString("some new text on file\n")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fileOpen.Close()

	// read file to checksum changed
	checksum32File(fileName)

}

func checksum32File(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	var buf bytes.Buffer

	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	checksum32 := crc32.ChecksumIEEE([]byte(buf.Bytes()))
	fmt.Printf("checksum32 hex: 0x%x\n", checksum32)
	fmt.Printf("checksum32 decimal: %v\n", checksum32)
}
