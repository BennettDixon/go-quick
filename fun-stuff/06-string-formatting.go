package main

import "fmt"

// source: https://go.dev/blog/strings

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	sampleBytes := []byte(sample)

	fmt.Println("Println:")
	fmt.Println(sampleBytes)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sampleBytes); i++ {
		fmt.Printf("%x ", sampleBytes[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sampleBytes)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sampleBytes)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sampleBytes)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sampleBytes)

	fmt.Println("Byte loop Printf with %+q:")
	for i := 0; i < len(sampleBytes); i++ {
		fmt.Printf("%q ", sampleBytes[i])
	}
	fmt.Printf("\n")
}
