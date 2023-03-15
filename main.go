package main

import "log"

func main() {
	store, err := newDBStore()

	if err != nil {
		log.Fatal(err)
	}
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	server := newapiServer(":" + port, store)
	server.Run()
}
