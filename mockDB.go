/**
 * Fake DB to simulte the CRUD operation
 */
package main

var currentID = 0
var listMessages Messages

// Some datas to store
func init() {
	StoreMessageDB(Message{Content: "Message XXXX"})
	StoreMessageDB(Message{Content: "Message YYYY"})
	StoreMessageDB(Message{Content: "my test message to store"})
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
	return Message{Content: "NULL"}
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
