/**
 * Here we define the models
 */
package main

import "time"

/*
 * Message model
 */
type Message struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

/*
 * Messages slice of Message
 */
type Messages []Message

/*
 * Messages slice of Message structure
 */
var (
	mockupMessages = Messages{
		Message{ID: 1, Content: "Hello World !"},
		Message{ID: 2, Content: "Goodbye World !"},
		Message{ID: 12345, Content: "my test message to store"},
	}
)
