// Copyright (c) 2016, CleverTap
// All rights reserved.
// Licensed under the New 3-Clause BSD License
//
// A mock server for the HTTP2 Apple push gateway

package main

import (
	"strconv"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	serverPort := flag.Int("port", 8443, "The port for the mock server to listen to")
	serverCert := flag.String("cert", "cert.pem", "The path to the certificate")
	serverKey := flag.String("key", "key.pem", "The path to the certificate key")

	flag.Parse()

	var srv http.Server
	http2.VerboseLogs = false
	srv.Addr = ":" + strconv.Itoa(*serverPort)
	http2.ConfigureServer(&srv, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := rand.Int31n(10);

		statusCode := 200

		var body string

		if n == 7 {
			statusCode = 400
			body = "{\"reason\": \"BadDeviceToken\"}"
		} else if n > 7 {
			statusCode = 410
			body = "{\"reason\": \"Unregistered\"}"
		}

		w.WriteHeader(statusCode)

		if statusCode != 200 {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, body)
		}
	})

	log.Println("Using certificate", *serverCert, "with key", *serverKey)
	log.Println("Starting server on port", *serverPort)
	log.Println("Hit ctrl + c to stop...")
	log.Fatal(srv.ListenAndServeTLS(*serverCert, *serverKey))
}