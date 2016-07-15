/**
 * Web logger
 * In standard library, there's not available web logging package
 * so we have to create it
 */
package main

import (
	"log"
	"net/http"
	"time"
)

/*
 * Logger function wrap the http.HandlerFunc in order to catch some informations
 * to print
 * @param inner http.handler, handler concerned
 * @param routeName string, name of the handled route
 * @return http.Handler
 */
func Logger(inner http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			routeName,
			time.Since(start),
		)
	})
}
