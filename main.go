package main

import "log"

func main() {
	store, err := newDBStore()

	if err != nil {
		log.Fatal(err)
	}

	server := newapiServer(":3000", store)
	server.Run()
}
