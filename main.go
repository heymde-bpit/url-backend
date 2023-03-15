package main

import "log"

func main() {
	store, err := newDBStore()

	if err != nil {
		log.Fatal(err)
	}

	server := newapiServer(":8080", store)
	server.Run()
}
