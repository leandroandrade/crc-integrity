package main

import (
	"hash/crc32"
	"fmt"
	"hash/crc64"
)

func main() {
	value32 := "this is a test"

	checksum32 := crc32.ChecksumIEEE([]byte(value32))
	fmt.Printf("checksum32 hex: 0x%x\n", checksum32)
	fmt.Printf("checksum32 decimal: %v\n", checksum32)

	value32 += " changed"
	checksum32 = crc32.ChecksumIEEE([]byte(value32))
	fmt.Printf("changed checksum32 hex: 0x%x\n", checksum32)
	fmt.Printf("changed checksum32 decimal: %v\n", checksum32)

	fmt.Println()

	value64 := "this is a new test"

	crcTable := crc64.MakeTable(crc64.ECMA)
	checksum64 := crc64.Checksum([]byte(value64), crcTable)
	fmt.Printf("checksum64 hex: 0x%x\n", checksum64)
	fmt.Printf("checksum64 decimal: %v\n", checksum64)

	value64 += " changed"
	checksum64 = crc64.Checksum([]byte(value64), crcTable)
	fmt.Printf("changed checksum64 hex: 0x%x\n", checksum64)
	fmt.Printf("changed checksum64 decimal: %v\n", checksum64)

}
