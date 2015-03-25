package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// open input file
	fi, err := os.Open("cipher")
	if err != nil {
		panic(err)
	}

	// close fi on exit with a defer
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(fi)

	// make a buffer to keep the chunks that are read
	buf := make([]byte, 1024)
	var s string
	read := 0

	for {
		// read chunk
		n, err := r.Read(buf)
		fmt.Println(n)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			s = string(buf[:read])
			break
		}
		read = n
	}

	fmt.Println(s)
}
