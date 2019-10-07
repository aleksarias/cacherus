package main

import (
	"github.com/chzyer/readline"
	"io"
	"log"
	"strings"
)

//
// cache.go
//
// A basic cache written in Go.
// This is an exercise where we will implement a simple key value store written in Go.
// We will use a simple readline interface and two commands: PUT and GET.
//
// Requirements:
//
// 1. PUT key value     Set a value in the cache.
// 2. GET key           Get a value stored in the cache.
// 3. EXIT/QUIT         Exits the interactive prompt (can also be done with Ctrl-d thanks to the readline pkg).
// 4. Use only packages from the stdlib (except for the readline package already imported below).
//

func main() {
	prompt, err := readline.New("> ")
	if err != nil {
		log.Fatal(err)
	}
	defer prompt.Close()

	cache, _ := NewCache(3)

to:
	for {
		line, err := prompt.Readline()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lineItems := strings.Split(line, " ")
		command := lineItems[0]
		switch command {
		case "PUT":
			key, value := lineItems[1], lineItems[2]
			cache.Put(key, value)
			log.Printf("PUT %v: %v", key, value)
		case "GET":
			key := lineItems[1]
			value, err := cache.Get(key)
			if err != nil {
				log.Printf("GOT %v: %v", key, err.Error())
				continue
			}
			log.Printf("GOT %v: %v", key, value)
		case "EXIT", "QUIT":
			log.Print("Exiting...")
			break to
		default:
			continue
		}
	}
}
