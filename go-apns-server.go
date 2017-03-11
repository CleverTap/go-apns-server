// Copyright (c) 2016, CleverTap
// All rights reserved.
// Licensed under the New 3-Clause BSD License
//
// A mock server for the HTTP2 Apple push gateway

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	serverPort := flag.Int("port", 8443, "The port for the mock server to listen to")
	serverCert := flag.String("cert", "cert.pem", "The path to the certificate")
	serverKey := flag.String("key", "key.pem", "The path to the certificate key")

	flag.Parse()

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("Error:", err)
		}
		num := rand.Int31n(10)

		statusCode := 200

		var body string

		if num == 7 {
			statusCode = 400
			body = "{\"reason\": \"BadDeviceToken\"}"
		} else if num > 7 {
			statusCode = 410
			body = "{\"reason\": \"Unregistered\"}"
		}

		resp.WriteHeader(statusCode)

		if statusCode != 200 {
			resp.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(resp, body)
		}
	})

	log.Printf("Using certificate %#v with key %#v", *serverCert, *serverKey)
	log.Println("Starting server on port", *serverPort)
	log.Println("Press Ctrl+C to stop...")

	serverAddr := ":" + strconv.Itoa(*serverPort)
	log.Fatal(http.ListenAndServeTLS(serverAddr, *serverCert, *serverKey, nil))
}
