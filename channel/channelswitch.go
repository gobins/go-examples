package main

import (
	"fmt"
)

func main() {
	msgChan := make(chan Message, 1)
	errChan := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"gobin@gobin.com"},
		From:    "golang@golang.org",
		Content: "This is a golang tutorial.",
	}
	failedMessage := FailedMessage{
		ErrorMessage:    "Message sending failed.",
		OriginalMessage: Message{},
	}

	msgChan <- msg
	errChan <- failedMessage

	select {
	case receivedMsg := <-msgChan:
		fmt.Println(receivedMsg)
	case receivedErr := <-errChan:
		fmt.Println(receivedErr)
	}
}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
