package main

import(
	"log"
)

func handleErr(customMsg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v\n", customMsg, err)
	}
}
