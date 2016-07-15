/**
 * Fake DB to simulte the CRUD operation
 */
package main

var currentID int = 1
var listMessages Messages

// Give us some seed data
func init() {
	StoreMessageDB(Message{Content: "Message XXXX"})
	StoreMessageDB(Message{Content: "Message YYYY"})
}

/*
 * GetMessageDB function to simulate the message to retrieve from its ID
 */
func GetMessageDB(id int) Message {
	for _, msg := range listMessages {
		if msg.ID == id {
			return msg
		}
	}
	// return empty Todo if not found
	return Message{Content: "there's no message stored"}
}

/*
 * StoreMessageDB function to simulate the creation of the message in the DB
 */
func StoreMessageDB(msg Message) Message {
	currentID++
	msg.ID = currentID
	listMessages = append(listMessages, msg)
	return msg
}
