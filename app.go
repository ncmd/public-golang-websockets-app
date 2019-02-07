// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var homefile_stock = ""
var homefile_chat = ""

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path == "/" {
		http.ServeFile(w, r, homefile_stock)
		return
	}
	if r.URL.Path == "/chat" {
		http.ServeFile(w, r, homefile_chat)
		return
	}
	// if r.URL.Path == "/main.js" {
	// 	http.ServeFile(w, r, "main.js")
	// 	return
	// }
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func main() {

	// This determines if the environment is local or production
	// need to create an environment variable:
	// export APP_ENV=local
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		if variable[0] == "APP_ENV" {
			fmt.Println(variable[1])
			if variable[1] == "local" {
				fmt.Println("Using Local Environment")
				homefile_stock = "home_local_stock.html"
				homefile_chat = "home_local_chat.html"
				// Use Production
			} else {
				homefile_stock = "home_prod_stock.html"
				homefile_chat = "home_prod_chat.html"
			}
		}
	}

	http.HandleFunc("/", serveHome)

	flag.Parse()
	hubChat := newHubChat()
	go hubChat.run()
	http.HandleFunc("/wschat", func(w http.ResponseWriter, r *http.Request) {
		serveWsChat(hubChat, w, r)
	})

	hubStock := newHubStock()
	go hubStock.run()
	http.HandleFunc("/wsstock", func(w http.ResponseWriter, r *http.Request) {
		serveWsStock(hubStock, w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Setting a Default port to 8000 to be used locally
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
