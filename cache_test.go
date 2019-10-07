package main

import (
	"log"
	"testing"
)

var tests = []Entry{
	{
		key:   "cat",
		value: "meow",
	},
	{
		key:   "dog",
		value: "woof",
	},
	{
		key:   "bird",
		value: "hoot",
	},
	{
		key:   "sheep",
		value: "baa",
	},
	{
		key:   "dog",
		value: "",
	},
	{
		key:   "cat",
		value: "",
	},
}

func TestSequential(t *testing.T) {
	cache, _ := NewCache(3)
	for test := range tests {
		entry := tests[test]
		if entry.value != "" {
			cache.Put(entry.key, entry.value)
			log.Printf("PUT %v: %v", entry.key, entry.value)
		} else {
			value, err := cache.Get(entry.key)
			if err != nil {
				log.Printf("GOT %v: %v", entry.key, err.Error())
				continue
			}
			log.Printf("GOT %v: %v", entry.key, value)
		}
	}
}
