package dish

import (
	"log"
	"time"
)

var (
	Dispatcher *Stream
)

// https://github.com/gin-gonic/examples/blob/master/server-sent-event/main.go
func NewDispatcher() (stream *Stream) {
	stream = &Stream{
		Message:       make(chan Message),
		NewClients:    make(chan chan Message),
		ClosedClients: make(chan chan Message),
		TotalClients:  make(map[chan Message]bool),
	}

	go stream.listen()
	go stream.heartbeat()

	return stream
}

func (stream *Stream) NewEvent(msg Message) {
	stream.Message <- msg
}

// https://github.com/gin-gonic/examples/blob/master/server-sent-event/main.go
func (stream *Stream) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
				log.Println("send message")
			}
		}
	}
}

func (stream *Stream) heartbeat() {
	var composingMessage bool = false

	for {
		if !composingMessage && time.Now().Unix()%25 == 0 {
			stream.Message <- Message{
				Content:    "heartbeat",
				SocketList: []string{"lmaoooo"},
				Timestamp:  time.Now().Unix(),
			}
			log.Println("message composed")

			time.Sleep(time.Second * 1)
			composingMessage = false
		}
	}
}
