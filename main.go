package main

import (
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	compress := flag.Bool("c", false, "Compress the input")
	extract := flag.Bool("x", false, "Extract the input")
	input := flag.String("i", "", "Input path")
	output := flag.String("o", "", "Output path")
	flag.Parse()

	if *input == "" {
		fmt.Println("ERROR: Input path empty.")
		os.Exit(1)
	}
	if *output == "" {
		fmt.Println("ERROR: Output path empty.")
		os.Exit(1)
	}
	if *compress && *extract || !*compress && !*extract {
		fmt.Println("ERROR: Either the compress or extract flag must be set. Both cannot be set.")
		os.Exit(1)
	}

	in, err := os.Open(*input)
	if err != nil {
		fmt.Printf("ERROR: Failed to open input file. %v\n", err)
		os.Exit(1)
	}
	defer in.Close()

	out, err := os.Create(*output)
	if err != nil {
		fmt.Printf("ERROR: Failed to create output file. %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	if *compress {
		zout := zlib.NewWriter(out)
		if _, err = io.Copy(zout, in); err != nil {
			fmt.Printf("ERROR: Failed to write to output file. %v\n", err)
			os.Exit(1)
		}
	} else if *extract {
		zin, err := zlib.NewReader(in)
		if err != nil {
			fmt.Printf("ERROR: Failed to create a zlib reader. %v\n", err)
			os.Exit(1)
		}
		if _, err = io.Copy(out, zin); err != nil {
			fmt.Printf("ERROR: Failed to write to output file. %v\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
