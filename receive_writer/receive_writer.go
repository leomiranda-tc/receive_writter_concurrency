package receive_writer

type ReceiverWriter interface {
	Receive(filename string, dataReceived []byte)
	Close()
}

func New() ReceiverWriter {
	return &implReceiveWriter{
		channelsMap: make(map[filename]chan []byte),
	}
}
