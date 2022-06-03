package receive_writer

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"sync"
	"time"
)

var (
	maxBufferChannel = 1000
)

type filename = string

type implReceiveWriter struct {
	channelsMap map[filename]chan []byte
}

func (r *implReceiveWriter) Close() {
	for _, channel := range r.channelsMap {
		close(channel)
	}
}

func (r *implReceiveWriter) Receive(filename string, dataReceived []byte) {

	// if handle writer for this file not initialize
	if _, ok := r.channelsMap[filename]; !ok {
		var m sync.Mutex

		m.Lock()
		r.channelsMap[filename] = make(chan []byte, maxBufferChannel)
		m.Unlock()

		go handleWriter(filename, r.channelsMap[filename])
	}

	// format and send to write sync
	r.channelsMap[filename] <- formatData(dataReceived)
}

func handleWriter(filename string, dataCh <-chan []byte) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// receive and write sync
	for {
		data, openChannel := <-dataCh
		if !openChannel {
			break
		}

		if _, err := f.Write(data); err != nil {
			fmt.Println("handleWriter:", err.Error())
			continue
		}

		fmt.Println("handleWriter:", "wrote complete")
	}

	fmt.Println("handleWriter: channel ", filename, " was closed")
}

func formatData(data []byte) []byte {
	return []byte(fmt.Sprintf("{\"received\":\"%v\",\"data\":%s}\n", time.Now().String(), data))
}
