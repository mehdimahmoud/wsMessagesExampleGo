/**
 * Example to launch a basic webserver
 * with handling the routes (endpoints)
 * We use here the "mux router"" from the Gorilla Web Toolkit
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
 * Function IndexHandler for the endpoint "" after domain
 * http.Request is a data structure that represents the client HTTP request
 * http.ResponseWriter is the response at any request
 */
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "\nHello World !")
}

/*
 * Function IndexMessagesHandler for the endpoint "/test" after "domain"
 * returns all mockup messages
 * http.Request is a data structure that represents the client HTTP request
 * http.ResponseWriter is the response at any request
 */
func IndexMessagesHandler(w http.ResponseWriter, r *http.Request) {
	// specify to the client that the request should be respect the JSON format
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// Encode the messages to send back some JSON to the client
	if err := json.NewEncoder(w).Encode(mockupMessages); err != nil {
		panic(err)
	}
}

/*
 * Function MessagesHandler for the endpoint "/messages" after "domain"
 * This function lists all messages
 * http.Request is a data structure that represents the client HTTP request
 * http.ResponseWriter is the response at any request
 */
func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "no message to show")
	// specify to the client that the request should be respect the JSON format
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// Encode the messages to send back some JSON to the client
	if err := json.NewEncoder(w).Encode(listMessages); err != nil {
		panic(err)
	}
}

/*
 * Function GetMessageHandler for the endpoint "/messages/{msgId}"
 * http.Request is a data structure that represents the client HTTP request
 * http.ResponseWriter is the response at any request
 */
func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	// retrieves the parameters from the request
	vars := mux.Vars(r)
	msgID := vars["msgID"]
	//fmt.Fprintln(w, "message ID: ", msgID)

	// From fake DB, storage the given message
	iMsgID, err := strconv.Atoi(msgID)
	if err != nil {
		panic(err)
	}
	msg := GetMessageDB(iMsgID)
	// Encode the msg to JSON in the response header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

/*
 * Function CreateMessageHandler for the endpoint "/messages/{msgId}"
 *  allows to store a message
 * @param r http.Request, a data structure that represents the client HTTP request
 * http.ResponseWriter is the response at any request
 */
func StoreMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message Message

	// This allows to limit the texte size to store
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// get from the body the plain texte
	content := string(body)
	//fmt.Fprintln(w, "message in the body: \n", content)
	message.Content = content

	// We can use this snippet code if the body is JSON format
	/*
		if err := json.Unmarshal(body, &message); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // respond with the status 422 and send back the err msg in JSON
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
	*/

	// From fake DB, storage the given message
	msg := StoreMessageDB(message)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}
