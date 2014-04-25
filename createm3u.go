package main

import "io/ioutil"
import "log"

func main() {

	entries, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		log.Println(entry.Name())
	}

	log.Println("No more files found!")
}
