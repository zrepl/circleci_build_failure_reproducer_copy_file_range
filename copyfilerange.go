package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("argv: ", os.Args)
	check := func(err error, msg string) {
		if err != nil {
			panic(fmt.Sprintf("%s: %s", msg, err))
		}
	}
	src, err := os.Open(os.Args[1])
	check(err, "open src")
	defer func() {
		err := src.Close()
		check(err, "close src")
	}()
	dst, err := os.Create(os.Args[2])
	check(err, "open dst")
	defer func() {
		err := dst.Close()
		check(err, "close dst")
	}()

	_, err = io.Copy(dst, src) // this command line runs 
	check(err, "io.Copy")
}
