package main

import (
	"./statsite"
	"log"
)

func main() {
	client, err := statsite.NewClient("0.0.0.0:8125")
	if err != nil {
		log.Println("Unable to connect to statsite: ", err)
	}
	msg := &statsite.CountMsg{"rewards", "1"}
	for i := 0; i < 20; i++ {
		client.Emit(msg)
	}
}
