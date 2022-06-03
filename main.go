package main

import (
	"strconv"
	"time"

	"github.com/leomirandadev/problem-1/receive_writer"
)

// main simulating the receive data
func main() {

	filename1 := "data1.txt"
	filename2 := "data2.txt"

	rw := receive_writer.New()

	for i := 0; i < 1000; i++ {

		go func(i int) {
			var filename string = filename1
			// change name of file to simulate to filenames
			if i%2 == 0 {
				filename = filename2
			}

			rw.Receive(filename, []byte("Writing message "+strconv.Itoa(i)))
		}(i)
	}

	time.Sleep(10 * time.Second)
	rw.Close()
	time.Sleep(10 * time.Second)

}
